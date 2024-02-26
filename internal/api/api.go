// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/pkg/vars/varflag"
)

const defaultOrigin = "https://github.com/cardano-community/koios-cli@v4"

var (
	clientSharedFlags = []varflag.FlagCreateFunc{
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
		varflag.BoolFunc("enable-req-stats", false, "Enable request stats"),
		varflag.BoolFunc("no-format", false, "prints response as machine readable json string"),
	}

	clientPagingFlags = []varflag.FlagCreateFunc{
		varflag.UintFunc("page", 1, "Set page number for paginated response"),
		varflag.UintFunc("page-size", koios.PageSize, "Set page size for paginated response"),
	}

	queryFlag = varflag.StringFunc("query", "", "Custom query paramaetrs for the request. e.g. key1=value1&key2=value2")
	epochFlag = varflag.UintFunc("epoch", 320, "Set epoch number")
)

type client struct {
	kc       *koios.Client
	noFormat bool
}

func Command() *happy.Command {
	cmd := happy.NewCommand("api",
		happy.Option("description", "Interact with Koios API REST endpoints"),
		happy.Option("before.shared", true),
		happy.Option("category", "API"),
	).WithFalgs(
		clientSharedFlags...,
	)

	api := &client{}
	cmd.Before(api.configure)

	// Network:
	// https://api.koios.rest/#tag--Network
	cmd.DescribeCategory("network", "Query information about the network")
	cmd.AddSubCommand(api.cmdNetworkTip())
	cmd.AddSubCommand(api.cmdNetworkGenesis())
	cmd.AddSubCommand(api.cmdNetworkTotals())
	cmd.AddSubCommand(api.cmdNetworkParamUpdates())
	cmd.AddSubCommand(api.cmdNetworkReserveWithdrawals())
	cmd.AddSubCommand(api.cmdNetworkTreasuryWithdrawals())

	// Epoch:
	// https://api.koios.rest/#tag--Epoch
	cmd.DescribeCategory("epoch", "Query epoch-specific details")
	cmd.AddSubCommand(notimplCmd("epoch", "epoch_info"))
	cmd.AddSubCommand(notimplCmd("epoch", "epoch_params"))
	cmd.AddSubCommand(notimplCmd("epoch", "epoch_block_protocols"))

	// Block:
	// https://api.koios.rest/#tag--Block
	cmd.DescribeCategory("block", "Query information about particular block on chain")
	cmd.AddSubCommand(notimplCmd("block", "blocks"))
	cmd.AddSubCommand(notimplCmd("block", "block_info"))
	cmd.AddSubCommand(notimplCmd("block", "block_txs"))

	// Transactions:
	// https://api.koios.rest/#tag--Transactions
	cmd.DescribeCategory("transactions", "Query blockchain transaction details")
	cmd.AddSubCommand(notimplCmd("transactions", "utxo_info"))
	cmd.AddSubCommand(notimplCmd("transactions", "tx_info"))
	cmd.AddSubCommand(notimplCmd("transactions", "tx_metadata"))
	cmd.AddSubCommand(notimplCmd("transactions", "tx_metalabels"))
	cmd.AddSubCommand(notimplCmd("transactions", "submittx"))
	cmd.AddSubCommand(notimplCmd("transactions", "tx_status"))
	cmd.AddSubCommand(notimplCmd("transactions", "tx_utxos"))

	// Stake Account:
	// https://api.koios.rest/#tag--Stake-Account
	cmd.DescribeCategory("stake account", "Query details about specific stake account addresses")
	cmd.AddSubCommand(notimplCmd("stake account", "account_list"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_info"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_info_cached"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_utxos"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_txs"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_rewards"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_updates"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_addresses"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_assets"))
	cmd.AddSubCommand(notimplCmd("stake account", "account_history"))

	// Address:
	// https://api.koios.rest/#tag--Address
	cmd.DescribeCategory("address", "Query information about specific address(es)")
	cmd.AddSubCommand(notimplCmd("address", "address_info"))
	cmd.AddSubCommand(notimplCmd("address", "address_utxos"))
	cmd.AddSubCommand(notimplCmd("address", "credential_utxos"))
	cmd.AddSubCommand(notimplCmd("address", "address_txs"))
	cmd.AddSubCommand(notimplCmd("address", "credential_txs"))
	cmd.AddSubCommand(notimplCmd("address", "address_assets"))

	// Asset:
	// https://api.koios.rest/#tag--Asset
	cmd.DescribeCategory("asset", "Query Asset related informations")
	cmd.AddSubCommand(notimplCmd("asset", "asset_list"))
	cmd.AddSubCommand(notimplCmd("asset", "policy_asset_list"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_token_registry"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_info"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_utxos"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_history"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_addresses"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_nft_address"))
	cmd.AddSubCommand(notimplCmd("asset", "policy_asset_addresses"))
	cmd.AddSubCommand(notimplCmd("asset", "policy_asset_info"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_summary"))
	cmd.AddSubCommand(notimplCmd("asset", "asset_txs"))

	// Pool:
	// https://api.koios.rest/#tag--Pool
	cmd.DescribeCategory("pool", "Query information about specific pools")
	cmd.AddSubCommand(notimplCmd("pool", "pool_list"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_info"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_stake_snapshot"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_delegators"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_delegators_history"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_blocks"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_history"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_updates"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_registrations"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_retirements"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_relays"))
	cmd.AddSubCommand(notimplCmd("pool", "pool_metadata"))

	// Script:
	// https://api.koios.rest/#tag--Script
	cmd.DescribeCategory("script", "Query information about specific scripts (Smart Contracts)")
	cmd.AddSubCommand(notimplCmd("script", "script_info"))
	cmd.AddSubCommand(notimplCmd("script", "native_script_list"))
	cmd.AddSubCommand(notimplCmd("script", "plutus_script_list"))
	cmd.AddSubCommand(notimplCmd("script", "script_redeemers"))
	cmd.AddSubCommand(notimplCmd("script", "script_utxos"))
	cmd.AddSubCommand(notimplCmd("script", "datum_info"))

	// Ogmios:
	// https://api.koios.rest/#tag--Ogmios
	cmd.DescribeCategory("ogmios", "Various stateless queries against Ogmios v6 instance")
	cmd.AddSubCommand(notimplCmd("ogmios", "ogmios"))
	return cmd
}

func (c *client) configure(sess *happy.Session, args happy.Args) (err error) {
	sess.Log().Debug("configure koios api client")

	apiVersion := args.Flag("api-version").String()
	enableReqStats, err := args.Flag("enable-req-stats").Var().Value().Bool()
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

	sess.Log().Debug(
		"configutation",
		slog.String("api-version", apiVersion),
		slog.Bool("enable-req-stats", enableReqStats),
		slog.Bool("no-format", c.noFormat),
		slog.Int("rate-limit", ratelimit),
		slog.String("sheme", sheme),
		slog.String("host", host),
		slog.Uint64("port", uint64(port)),
		slog.String("origin", origin),
	)
	c.kc, err = koios.New(
		koios.APIVersion(apiVersion),
		koios.EnableRequestsStats(enableReqStats),
		koios.Scheme(sheme),
		koios.Host(host),
		koios.Origin(origin),
		koios.Port(uint16(port)),
		koios.RateLimit(ratelimit),
	)
	return
}

func (c *client) newRequestOpts(sess *happy.Session, args happy.Args) (*koios.RequestOptions, error) {
	opts := c.kc.NewRequestOptions()
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

	if noFormat {
		out, err := json.Marshal(data)
		if handleErr(noFormat, err) {
			return
		}
		fmt.Println(string(out))
		return
	}

	out, merr := json.MarshalIndent(data, "", " ")
	if handleErr(noFormat, merr) {
		return
	}
	fmt.Println(string(out))
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
