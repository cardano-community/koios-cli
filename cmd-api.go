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
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/cardano-community/koios-go-client"
)

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
			if c.Bool("testnet") {
				hostopt = koios.Host(koios.TestnetHost)
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
			opts.Page(c.Uint("page"))
			opts.PageSize(c.Uint("page-size"))

			return err
		},
	}

	attachAPIAccountCommmands(apicmd)
	attachAPIAddressCommmands(apicmd)
	attachAPIAssetsCommmands(apicmd)
	attachAPIBlocksCommmands(apicmd)
	attachAPIEpochCommmands(apicmd)
	attachAPIGeneralCommmands(apicmd)
	attachAPINetworkCommmands(apicmd)
	attachAPIScriptCommmands(apicmd)
	attachAPITransactionsCommmands(apicmd)
	attachAPIPoolCommmands(apicmd)

	app.Commands = append(app.Commands, apicmd)
}

func apiCommonFlags() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:  "port",
			Usage: "Set port",
			Value: uint(koios.DefaultPort),
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
			Name:  "testnet",
			Usage: "use default testnet as host.",
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

func attachAPIAccountCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "account-list",
			Category: "ACCOUNT",
			Usage:    "Get a list of all accounts returns array of stake addresses.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountList(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-info",
			Category: "ACCOUNT",
			Usage:    "Get the account info of any (payment or staking) address.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountInfo(callctx, koios.Address(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-rewards",
			Category: "ACCOUNT",
			Usage:    "Get the full rewards history (including MIR) for a stake address, or certain epoch if specified.",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:    "epoch",
					Aliases: []string{"e"},
					Usage:   "Filter for earned rewards Epoch Number.",
					Value:   uint64(0),
				},
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}
				res, err := api.GetAccountRewards(callctx, koios.StakeAddress(ctx.String("address")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-updates",
			Category: "ACCOUNT",
			Usage:    "Get the account updates (registration, deregistration, delegation and withdrawals).",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountUpdates(callctx, koios.StakeAddress(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-addresses",
			Category: "ACCOUNT",
			Usage:    "Get all addresses associated with an account payment or staking address",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountAddresses(callctx, koios.StakeAddress(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-assets",
			Category: "ACCOUNT",
			Usage:    "Get the native asset balance of an account.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountAssets(callctx, koios.StakeAddress(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-history",
			Category: "ACCOUNT",
			Usage:    "Get the staking history of an account.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountHistory(callctx, koios.StakeAddress(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIAddressCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "address-info",
			Category: "ADDRESS",
			Usage:    "Get address info - balance, associated stake address (if any) and UTxO set.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAddressInfo(callctx, koios.Address(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "address-txs",
			Category: "ADDRESS",
			Usage: "Get the transaction hash list of input address array, optionally " +
				"filtering after specified block height (inclusive).",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:  "after-block-height",
					Usage: "Get transactions after specified block height.",
					Value: uint64(0),
				},
				&cli.StringSliceFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format, can be used multiple times for list of addresses",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var addresses []koios.Address
				for _, a := range ctx.StringSlice("address") {
					addresses = append(addresses, koios.Address(a))
				}

				res, err := api.GetAddressTxs(callctx, addresses, ctx.Uint64("after-block-height"), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "address-assets",
			Category: "ADDRESS",
			Usage:    "Get the list of all the assets (policy, name and quantity) for a given address.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAddressAssets(callctx, koios.Address(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "credential-txs",
			Category: "ADDRESS",
			Usage: "Get the transaction hash list of input payment credentials, " +
				"optionally filtering after specified block height (inclusive).",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:  "after-block-height",
					Usage: "Get transactions after specified block height.",
					Value: uint64(0),
				},
				&cli.StringSliceFlag{
					Name:     "payment-credential",
					Aliases:  []string{"c"},
					Usage:    "Cardano payment credential, can be used multiple times for list of addresses",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var credentials []koios.PaymentCredential
				for _, c := range ctx.StringSlice("payment-credential") {
					credentials = append(credentials, koios.PaymentCredential(c))
				}

				res, err := api.GetCredentialTxs(callctx, credentials, ctx.Uint64("after-block-height"), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIAssetsCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "asset-list",
			Category: "ASSET",
			Usage:    "Get the list of all native assets (paginated).",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetList(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "asset-address-list",
			Category: "ASSET",
			Usage:    "Get the list of all addresses holding a given asset.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "policy",
					Aliases:  []string{"p"},
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetAddressList(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "asset-info",
			Category: "ASSET",
			Usage:    "Get the information of an asset including first minting & token registry metadata.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "policy",
					Aliases:  []string{"p"},
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetInfo(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "asset-summary",
			Category: "ASSET",
			Usage: "Get the summary of an asset (total transactions exclude " +
				"minting/total wallets include only wallets with asset balance).",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "policy",
					Aliases:  []string{"p"},
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetSummary(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "asset-history",
			Category: "ASSET",
			Usage:    "Get the mint/burn history of an asset.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "policy",
					Aliases:  []string{"p"},
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetHistory(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "asset-txs",
			Category: "ASSET",
			Usage:    "Get the list of all asset transaction hashes (newest first).",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "policy",
					Aliases:  []string{"p"},
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Aliases:  []string{"n"},
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetTxs(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "asset-policy-info",
			Category: "ASSET",
			Usage:    "Get the information for all assets under the same policy.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "policy",
					Aliases:  []string{"p"},
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetPolicyInfo(callctx, koios.PolicyID(ctx.String("policy")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIBlocksCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "blocks",
			Category: "BLOCK",
			Usage:    "Get summarised details about all blocks (paginated - latest first).",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetBlocks(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "blocks-info",
			Category: "BLOCK",
			Usage:    "Get detailed information about a blocks.",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "block-hash",
					Aliases:  []string{"b"},
					Usage:    "Block Hashes in hex format, can be used multiple times for list of blocks",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var bhses []koios.BlockHash
				for _, h := range ctx.StringSlice("block-hash") {
					bhses = append(bhses, koios.BlockHash(h))
				}
				res, err := api.GetBlocksInfo(
					callctx,
					bhses,
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "block-info",
			Category: "BLOCK",
			Usage:    "Get detailed information about a specific block.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "block-hash",
					Aliases:  []string{"b"},
					Usage:    "Block Hash in hex format to fetch details for",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetBlockInfo(
					callctx,
					koios.BlockHash(ctx.String("block-hash")),
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "block-txs",
			Category: "BLOCK",
			Usage:    "Get a list of all transactions included in a provided block.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "block-hash",
					Aliases:  []string{"b"},
					Usage:    "Block Hash in hex format to fetch details for",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetBlockTxHashes(
					callctx,
					koios.BlockHash(ctx.String("block-hash")),
					opts,
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIEpochCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "epoch-info",
			Category: "EPOCH",
			Usage:    "Get the epoch information, all epochs if no epoch specified.",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:    "epoch",
					Aliases: []string{"e"},
					Usage:   "Epoch Number to fetch details for",
					Value:   uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetEpochInfo(callctx, epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "epoch-params",
			Category: "EPOCH",
			Usage: "Get the protocol parameters for specific epoch, " +
				"returns information about all epochs if no epoch specified.",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:    "epoch",
					Aliases: []string{"e"},
					Usage:   "Epoch Number to fetch details for",
					Value:   uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetEpochParams(callctx, epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIGeneralCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:      "get",
			Usage:     "send GET request to the specified API endpoint",
			Category:  "UTILS",
			ArgsUsage: "[endpoint]",
			Action: func(ctx *cli.Context) error {
				uri := ctx.Args().Get(0)
				if len(uri) == 0 {
					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument e.g. /tip")
				}

				u, err := url.ParseRequestURI(uri)
				handleErr(err)

				opts.QueryApply(u.Query())

				res, err := api.GET(callctx, u.Path, opts)
				handleErr(err)
				defer res.Body.Close()
				body, err := io.ReadAll(res.Body)
				handleErr(err)

				printJSON(ctx, body)
				return nil
			},
		},
		{
			Name:      "post",
			Usage:     "send POST request to the specified API endpoint",
			Category:  "UTILS",
			ArgsUsage: "[endpoint] [payload]",
			Action: func(ctx *cli.Context) error {
				uri := ctx.Args().Get(0)
				pl := ctx.Args().Get(1)
				if len(uri) == 0 {
					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument 1 e.g. /tip")
				}
				if len(pl) == 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "provide payload as argument 2")
				}

				u, err := url.ParseRequestURI(uri)
				handleErr(err)

				opts.QueryApply(u.Query())

				res, err := api.POST(callctx, u.Path, strings.NewReader(pl), opts)
				handleErr(err)
				defer res.Body.Close()
				body, err := io.ReadAll(res.Body)
				handleErr(err)

				printJSON(ctx, body)
				return nil
			},
		},
		{
			Name:      "head",
			Usage:     "head issues a HEAD request to the specified API endpoint",
			Category:  "UTILS",
			ArgsUsage: "[endpoint]",
			Action: func(ctx *cli.Context) error {
				uri := ctx.Args().Get(0)
				if ctx.NArg() == 0 || len(uri) == 0 {
					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument e.g. /tip")
				}
				u, err := url.ParseRequestURI(uri)
				handleErr(err)

				opts.QueryApply(u.Query())

				res, err := api.HEAD(callctx, u.Path, opts)
				handleErr(err)
				if res.Body != nil {
					res.Body.Close()
				}
				fmt.Println(res.Request.URL.String())
				fmt.Println(res.Status)
				return nil
			},
		},
	}...)
}

func attachAPINetworkCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "tip",
			Category: "NETWORK",
			Usage:    "Get the tip info about the latest block seen by chain.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTip(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "genesis",
			Category: "NETWORK",
			Usage:    "Get the Genesis parameters used to start specific era on chain.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetGenesis(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "totals",
			Category: "NETWORK",
			Usage: "Get the circulating utxo, treasury rewards, " +
				"supply and reserves in lovelace for specified epoch, all epochs if empty.",
			Flags: []cli.Flag{
				&cli.UintFlag{
					Name:    "epoch",
					Aliases: []string{"e"},
					Usage:   "Epoch Number to fetch details for",
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint("epoch"))
					epoch = &v
				}

				res, err := api.GetTotals(callctx, epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIScriptCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "native-script-list",
			Category: "SCRIPT",
			Usage:    "List of all existing native script hashes along with their creation transaction hashes.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetNativeScriptList(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "plutus-script-list",
			Category: "SCRIPT",
			Usage:    "List of all existing Plutus script hashes along with their creation transaction hashes.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPlutusScriptList(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "script-redeemers",
			Category: "SCRIPT",
			Usage:    "List of all redeemers for a given script hash.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "script-hash",
					Aliases:  []string{"s"},
					Usage:    "Script hash in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetScriptRedeemers(callctx, koios.ScriptHash(ctx.String("script-hash")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPITransactionsCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "txs-info",
			Category: "TRANSACTIONS",
			Usage:    "Get detailed information about transaction(s).",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "tx-hash",
					Aliases:  []string{"t"},
					Usage:    "Transaction Hashes, can be used multiple times for list of transactions",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var txs []koios.TxHash
				for _, tx := range ctx.StringSlice("tx-hash") {
					txs = append(txs, koios.TxHash(tx))
				}
				res, err := api.GetTxsInfo(callctx, txs, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "tx-info",
			Category: "TRANSACTIONS",
			Usage:    "Get detailed information about single transaction.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "tx-hash",
					Aliases:  []string{"t"},
					Usage:    "Transaction Hash to fetch details for",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTxInfo(callctx, koios.TxHash(ctx.String("tx-hash")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "tx-utxos",
			Category: "TRANSACTIONS",
			Usage:    "Get UTxO set (inputs/apiOutputs) of transactions.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "tx-hash",
					Aliases:  []string{"t"},
					Usage:    "Transaction Hash to fetch details for",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTxUTxOs(callctx, koios.TxHash(ctx.String("tx-hash")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "txs-metadata",
			Category: "TRANSACTIONS",
			Usage:    "Get metadata information (if any) for given transaction(s).",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "tx-hash",
					Aliases:  []string{"t"},
					Usage:    "Transaction Hash in hex format, can be used multiple times for list of transactions",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var txs []koios.TxHash
				for _, tx := range ctx.StringSlice("tx-hash") {
					txs = append(txs, koios.TxHash(tx))
				}
				res, err := api.GetTxsMetadata(callctx, txs, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "tx-metadata",
			Category: "TRANSACTIONS",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "tx-hash",
					Aliases:  []string{"t"},
					Usage:    "Transaction Hash to fetch details for",
					Required: true,
				},
			},
			Usage: "Get metadata information (if any) for given transaction.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTxMetadata(callctx, koios.TxHash(ctx.String("tx-hash")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "tx-metalabels",
			Category: "TRANSACTIONS",
			Usage:    "Get a list of all transaction metalabels.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTxMetaLabels(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "tx-submit",
			Category: "TRANSACTIONS",
			Usage:    "Submit signed transaction to the network.",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "submittx requires single arg (path to file tx.signed)")
				}

				stx := koios.TxBodyJSON{}

				txfile, err := os.ReadFile(ctx.Args().Get(0))
				if err != nil {
					return err
				}

				if err = json.Unmarshal(txfile, &stx); err != nil {
					return err
				}
				res, err := api.SubmitSignedTx(callctx, stx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "txs-statuses",
			Category: "TRANSACTIONS",
			Usage:    "Get the number of block confirmations for a given transaction hash list",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "tx-hash",
					Aliases:  []string{"t"},
					Usage:    "Transaction Hash in hex format, can be used multiple times for list of transactions",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var txs []koios.TxHash
				for _, tx := range ctx.StringSlice("tx-hash") {
					txs = append(txs, koios.TxHash(tx))
				}
				res, err := api.GetTxsStatuses(callctx, txs, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "tx-status",
			Category: "TRANSACTIONS",
			Usage:    "Get the number of block confirmations for a given transaction hash",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "tx-hash",
					Aliases:  []string{"t"},
					Usage:    "Transaction Hash to fetch details for",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTxStatus(callctx, koios.TxHash(ctx.String("tx-hash")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIPoolCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "pool-list",
			Category: "POOL",
			Usage:    "A list of all currently registered/retiring (not retired) pools.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolList(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-infos",
			Category: "POOL",
			Usage:    "Current pool statuses and details for a specified list of pool ids.",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format, can be used multiple times for list of transactions",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var pids []koios.PoolID
				for _, pid := range ctx.StringSlice("pool-id") {
					pids = append(pids, koios.PoolID(pid))
				}
				res, err := api.GetPoolInfos(callctx, pids, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-info",
			Category: "POOL",
			Usage:    "Current pool status and details for a specified pool by pool id.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolInfo(callctx, koios.PoolID(ctx.String("pool-id")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-delegators",
			Category: "POOL",
			Usage:    "Return information about delegators by a given pool and optional epoch (current if omitted).",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:    "epoch",
					Aliases: []string{"e"},
					Usage:   "Epoch Number to fetch details for",
					Value:   uint64(0),
				},
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetPoolDelegators(callctx, koios.PoolID(ctx.String("pool-id")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-history",
			Category: "POOL",
			Usage:    "Return information about delegators by a given pool and optional epoch (current if omitted).",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:    "epoch",
					Aliases: []string{"e"},
					Usage:   "Epoch Number to fetch details for",
					Value:   uint64(0),
				},
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetPoolHistory(callctx, koios.PoolID(ctx.String("pool-id")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-blocks",
			Category: "POOL",
			Usage:    "Return information about blocks minted by a given pool in current epoch (or _epoch_no if provided).",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:    "epoch",
					Aliases: []string{"e"},
					Usage:   "Epoch Number to fetch details for",
					Value:   uint64(0),
				},
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetPoolBlocks(callctx, koios.PoolID(ctx.String("pool-id")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-updates",
			Category: "POOL",
			Usage:    "Return all pool updates for all pools or only updates for specific pool if specified.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				pool := koios.PoolID(ctx.String("pool-id"))

				res, err := api.GetPoolUpdates(callctx, &pool, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-relays",
			Category: "POOL",
			Usage:    "A list of registered relays for all currently registered/retiring (not retired) pools.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolRelays(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-metadata",
			Category: "POOL",
			Usage:    "Metadata(on & off-chain) for all currently registered/retiring (not retired) pools.",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format, can be used multiple times for list of transactions",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var pids []koios.PoolID
				for _, pid := range ctx.StringSlice("pool-id") {
					pids = append(pids, koios.PoolID(pid))
				}

				res, err := api.GetPoolMetadata(callctx, pids, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}
