// Copyright 2022 The Cardano Community Authors
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//   or LICENSE file in repository root.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/urfave/cli/v2"

	"github.com/cardano-community/koios-go-client/v3"
)

var epochFlag = &cli.Uint64Flag{
	Name:    "epoch",
	Aliases: []string{"e"},
	Usage:   "Epoch Number to fetch details for",
	Value:   uint64(0),
}

func attachAPICommmand(app *cli.App) {
	apicmd := &cli.Command{
		Name:     "api",
		Category: "KOIOS REST API",
		Usage:    "Interact with Koios API REST endpoints",
		Flags:    apiCommonFlags(),
		Before: func(c *cli.Context) error {
			var (
				hostopt koios.Option
				err     error
			)
			if c.Bool("preview") {
				hostopt = koios.Host(koios.PreviewHost)
			} else if c.Bool("preprod") {
				hostopt = koios.Host(koios.PreProdHost)
			} else if c.Bool("guildnet") {
				hostopt = koios.Host(koios.GuildHost)
			} else {
				hostopt = koios.Host(c.String("host"))
			}

			api, err = koios.New(
				hostopt,
				koios.APIVersion(c.String("api-version")),
				koios.Port(uint16(c.Uint("port"))),
				koios.Scheme(c.String("scheme")),
				koios.RateLimit(c.Int("rate-limit")),
				koios.Origin(c.String("origin")),
				koios.CollectRequestsStats(c.Bool("enable-req-stats")),
			)

			opts = api.NewRequestOptions()
			opts.SetCurrentPage(c.Uint("page"))
			opts.SetPageSize(c.Uint("page-size"))

			return err
		},
	}

	attachAPINetworkCommmands(apicmd)
	attachAPIEpochCommmands(apicmd)
	attachAPIBlockCommmands(apicmd)
	attachAPITransactionsCommmands(apicmd)
	attachAPIAddressCommmands(apicmd)
	attachAPIAccountCommmands(apicmd)
	attachAPIAssetsCommmands(apicmd)
	attachAPIPoolCommmands(apicmd)
	attachAPIScriptCommmands(apicmd)

	// attachAPIGeneralCommmands(apicmd)

	app.Commands = append(app.Commands, apicmd)
}

func apiCommonFlags() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:  "port",
			Usage: "Set port",
		},
		&cli.StringFlag{
			Name:  "host",
			Usage: "Set host",
			Value: koios.MainnetHost,
		},
		&cli.StringFlag{
			Name:  "api-version",
			Usage: "Set API version",
			Value: koios.DefaultAPIVersion,
		},
		&cli.StringFlag{
			Name:  "scheme",
			Usage: "Set URL scheme",
			Value: koios.DefaultScheme,
		},
		&cli.StringFlag{
			Name:  "origin",
			Usage: "Set Origin header for requests.",
			Value: koios.DefaultOrigin,
		},
		&cli.UintFlag{
			Name:  "rate-limit",
			Usage: "Set API Client rate limit for outgoing requests",
			Value: uint(koios.DefaultRateLimit),
		},
		&cli.BoolFlag{
			Name:  "no-format",
			Usage: "prints response json strings directly without calling json pretty.",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "enable-req-stats",
			Usage: "Enable request stats.",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "preview",
			Usage: "use preview host.",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "preprod",
			Usage: "use preprod host.",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "guildnet",
			Usage: "use guildnet host.",
			Value: false,
		},
		&cli.UintFlag{
			Name:  "page",
			Usage: "Set current page for request",
			Value: 1,
		},
		&cli.UintFlag{
			Name:  "page-size",
			Usage: "Set page size for request",
			Value: 1000,
		},
	}
}

// func attachAPIGeneralCommmands(apicmd *cli.Command) {
// 	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
// 		{
// 			Name:      "get",
// 			Usage:     "send GET request to the specified API endpoint",
// 			Category:  "UTILS",
// 			ArgsUsage: "[endpoint]",
// 			Action: func(ctx *cli.Context) error {
// 				uri := ctx.Args().Get(0)
// 				if len(uri) == 0 {
// 					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument e.g. /tip")
// 				}

// 				u, err := url.ParseRequestURI(uri)
// 				handleErr(err)

// 				opts.QueryApply(u.Query())

// 				res, err := api.GET(callctx, u.Path, opts)
// 				handleErr(err)
// 				defer res.Body.Close()
// 				body, err := io.ReadAll(res.Body)
// 				handleErr(err)

// 				printJSON(ctx, body)
// 				return nil
// 			},
// 		},
// 		{
// 			Name:      "post",
// 			Usage:     "send POST request to the specified API endpoint",
// 			Category:  "UTILS",
// 			ArgsUsage: "[endpoint] [payload]",
// 			Action: func(ctx *cli.Context) error {
// 				uri := ctx.Args().Get(0)
// 				pl := ctx.Args().Get(1)
// 				if len(uri) == 0 {
// 					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument 1 e.g. /tip")
// 				}
// 				if len(pl) == 1 {
// 					return fmt.Errorf("%w: %s", ErrCommand, "provide payload as argument 2")
// 				}

// 				u, err := url.ParseRequestURI(uri)
// 				handleErr(err)

// 				opts.QueryApply(u.Query())

// 				res, err := api.POST(callctx, u.Path, strings.NewReader(pl), opts)
// 				handleErr(err)
// 				defer res.Body.Close()
// 				body, err := io.ReadAll(res.Body)
// 				handleErr(err)

// 				printJSON(ctx, body)
// 				return nil
// 			},
// 		},
// 		{
// 			Name:      "head",
// 			Usage:     "head issues a HEAD request to the specified API endpoint",
// 			Category:  "UTILS",
// 			ArgsUsage: "[endpoint]",
// 			Action: func(ctx *cli.Context) error {
// 				uri := ctx.Args().Get(0)
// 				if ctx.NArg() == 0 || len(uri) == 0 {
// 					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument e.g. /tip")
// 				}
// 				u, err := url.ParseRequestURI(uri)
// 				handleErr(err)

// 				opts.QueryApply(u.Query())

// 				res, err := api.HEAD(callctx, u.Path, opts)
// 				handleErr(err)
// 				if res.Body != nil {
// 					res.Body.Close()
// 				}
// 				fmt.Println(res.Request.URL.String())
// 				fmt.Println(res.Status)
// 				return nil
// 			},
// 		},
// 	}...)
// }
