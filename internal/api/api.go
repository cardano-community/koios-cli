// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"sync"
	"time"

	"github.com/cardano-community/koios-cli/v2/internal/auth"
	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/pkg/vars/varflag"
)

const defaultOrigin = "https://github.com/cardano-community/koios-cli/v2"

var (
	pagingFlags = []varflag.FlagCreateFunc{
		varflag.UintFunc("page", 1, "Set page number for paginated response"),
		varflag.UintFunc("page-size", koios.PageSize, "Set page size for paginated response"),
	}

	queryFlag = varflag.StringFunc("query", "", "Custom query for the request. e.g. key1=value1&key2=value2")

	// koios api params
	epochNoFlag = varflag.UintFunc("epoch", 320, "Set epoch number")

	afterBlockHeightFlag = varflag.UintFunc("after-block-height", 0, "Block height for specifying time delta")
)

type client struct {
	mu           sync.Mutex
	kc           *koios.Client
	noFormat     bool
	subscription *auth.Subscription
}

func Command() *happy.Command {
	cmd := happy.NewCommand("api",
		happy.Option("description", "Interact with Koios API REST endpoints"),
		happy.Option("before.shared", true),
		// happy.Option("category", "API"), // enable when more subcommands are implemented
	).WithFlags(
		varflag.UintFunc("port", uint(koios.DefaultPort), "Set port number for the API server"),
		varflag.StringFunc("scheme", koios.DefaultScheme, "Set scheme for the API server"),
		varflag.StringFunc("api-version", koios.DefaultAPIVersion, "Set API version"),
		varflag.IntFunc("rate-limit", koios.DefaultRateLimit, "Set rate limit for the API server"),
		varflag.StringFunc("origin", defaultOrigin, "Set origin for the API server"),
		varflag.StringFunc("host", koios.MainnetHost, "Set host for the API server"),
		varflag.BoolFunc("host-eu", false, "Use eu mainet network host"),
		varflag.BoolFunc("host-preview", false, "Use preview network host"),
		varflag.BoolFunc("host-preprod", false, "Use preprod network host"),
		varflag.BoolFunc("host-guildnet", false, "Use guildnet network host"),
		varflag.BoolFunc("stats", false, "Enable request stats"),
		varflag.BoolFunc("no-format", false, "prints response as machine readable json string"),
		varflag.DurationFunc("timeout", time.Duration(time.Minute), "Set timeout for the API server"),
		varflag.StringFunc("auth", "", "JWT Bearer Auth token generated via https://koios.rest Profile page."),
	)

	api := &client{}
	cmd.Before(api.configure)

	// Add allcategorized subcommands
	network(cmd, api)
	epoch(cmd, api)
	block(cmd, api)
	transactions(cmd, api)
	stakeAccount(cmd, api)
	address(cmd, api)
	asset(cmd, api)
	pool(cmd, api)
	script(cmd, api)
	ogmios(cmd, api)

	return cmd
}

func (c *client) configure(sess *happy.Session, args happy.Args) (err error) {
	sess.Log().Debug("configure koios api client")

	apiVersion := args.Flag("api-version").String()
	enableReqStats, err := args.Flag("stats").Var().Value().Bool()
	if err != nil {
		return err
	}

	c.noFormat = args.Flag("no-format").Var().Bool()
	sheme := args.Flag("scheme").String()
	host, err := getHost(args)
	if err != nil {
		return err
	}
	port := args.Flag("port").Var().Uint()
	origin := defaultOrigin
	if args.Flag("origin").Present() {
		origin = args.Flag("origin").String()
	}
	ratelimit := args.Flag("rate-limit").Var().Int()
	duration := args.Flag("timeout").Var().Duration()

	sess.Log().Debug(
		"configutation",
		slog.String("api-version", apiVersion),
		slog.Bool("stats", enableReqStats),
		slog.Bool("no-format", c.noFormat),
		slog.Int("rate-limit", ratelimit),
		slog.String("sheme", sheme),
		slog.String("host", host),
		slog.Uint64("port", uint64(port)),
		slog.String("origin", origin),
		slog.Duration("timeout", duration),
	)
	c.kc, err = koios.New(
		koios.APIVersion(apiVersion),
		koios.EnableRequestsStats(enableReqStats),
		koios.Scheme(sheme),
		koios.Host(host),
		koios.Origin(origin),
		koios.Port(uint16(port)),
		koios.RateLimit(ratelimit),
		koios.Timeout(duration),
	)

	if args.Flag("profile").Present() && args.Flag("auth").Present() {
		return fmt.Errorf("profile and auth flags cannot be used together")
	}

	if args.Flag("auth").Present() {
		if err := c.kc.SetAuth(args.Flag("auth").String()); err != nil {
			return fmt.Errorf("failed to set auth token: %w", err)
		}
	} else if args.Flag("profile").Present() {
		subscription, err := auth.LoadSubscription(sess)
		if err != nil {
			return err
		}
		c.subscription = subscription
		return c.kc.SetAuth(subscription.JWT)
	}

	return
}

func (c *client) koios() *koios.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.kc
}

func (c *client) newRequestOpts(sess *happy.Session, args happy.Args) (*koios.RequestOptions, error) {
	opts := c.koios().NewRequestOptions()
	if args == nil {
		return opts, nil
	}
	if args.Flag("page").Present() {
		opts.SetCurrentPage(args.Flag("page").Var().Uint())
	}
	if args.Flag("page-size").Present() {
		opts.SetPageSize(args.Flag("page-size").Var().Uint())
	}
	if args.Flag("query").Present() {
		qraw := args.Flag("query").String()
		q, err := url.ParseQuery(qraw)
		if err != nil {
			return nil, fmt.Errorf("failed to parse query parameters: %w", err)
		}
		opts.QueryApply(q)
	}

	if c.subscription != nil {
		c.subscription.RequestsToday++
		opts.SetRequestsToday(c.subscription.RequestsToday)
		if err := c.subscription.Save(); err != nil {
			return nil, fmt.Errorf("failed to save subscription: %w", err)
		}
	}

	// opts.SetPageSize(args.Flag("page-size").Var().Uint())
	return opts, nil
}

func getHost(args happy.Args) (string, error) {
	host := args.Flag("host").String()

	hostflags := 0
	if args.Flag("host").Present() {
		host = args.Flag("host").String()
		hostflags++
	}
	if args.Flag("host-eu").Present() && args.Flag("host-eu").Var().Bool() {
		host = koios.MainnetHostEU
		hostflags++
	}
	if args.Flag("host-preview").Present() && args.Flag("host-preview").Var().Bool() {
		host = koios.PreviewHost
		hostflags++
	}
	if args.Flag("host-preprod").Present() && args.Flag("host-preprod").Var().Bool() {
		host = koios.PreProdHost
		hostflags++
	}
	if args.Flag("host-guildnet").Present() && args.Flag("host-guildnet").Var().Bool() {
		host = koios.GuildHost
		hostflags++
	}
	if hostflags > 1 {
		return "", fmt.Errorf("only one host flag can be used")
	}
	return host, nil
}

// output koios api client responses.
func apiOutput(noFormat bool, data any, err error) {
	if err != nil {
		handleErr(noFormat, err)
		return
	}

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	if !noFormat {
		encoder.SetIndent("", "  ")
	}

	if err := encoder.Encode(data); err != nil {
		handleErr(noFormat, err)
		return
	}
	fmt.Println(buffer.String())
}

type apiError struct {
	Error string `json:"error"`
}

func handleErr(noFormat bool, err error) bool {
	if err == nil {
		return false
	}
	apiOutput(noFormat, apiError{err.Error()}, nil)
	return true
}

func flagSlice(flags ...varflag.FlagCreateFunc) []varflag.FlagCreateFunc {
	return flags
}

func notimplCmd(category, name string) *happy.Command {
	cmd := happy.NewCommand(name,
		happy.Option("description", "NOT IMPLEMENTED"),
		happy.Option("category", category),
	)
	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		sess.Log().NotImplemented(fmt.Sprintf("command %q not implemented", name))
		return nil
	})
	return cmd
}
