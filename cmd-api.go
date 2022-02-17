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
	"io/ioutil"

	"github.com/urfave/cli/v2"

	"github.com/cardano-community/koios-go-client"
)

func attachAPICommmand(app *cli.App) {
	api, err := koios.New()
	handleErr(err)

	apicmd := &cli.Command{
		Name:     "api",
		Category: "KOIOS REST API",
		Usage:    "Interact with Koios API REST endpoints",
		Flags:    apiCommonFlags(),
		Before: func(c *cli.Context) error {
			if c.Bool("testnet") {
				handleErr(koios.Host(koios.TestnetHost)(api))
			} else {
				handleErr(koios.Host(c.String("host"))(api))
			}
			handleErr(koios.APIVersion(c.String("api-version"))(api))
			handleErr(koios.Port(uint16(c.Uint("port")))(api))
			handleErr(koios.Schema(c.String("schema"))(api))
			handleErr(koios.RateLimit(uint8(c.Uint("rate-limit")))(api))
			handleErr(koios.Origin(c.String("origin"))(api))
			handleErr(koios.CollectRequestsStats(c.Bool("enable-req-stats"))(api))
			return nil
		},
	}

	attachAPIAccountCommmands(apicmd, api)
	attachAPIAddressCommmands(apicmd, api)
	attachAPIAssetsCommmands(apicmd, api)
	attachAPIBlocksCommmands(apicmd, api)
	attachAPIEpochCommmands(apicmd, api)
	attachAPIGeneralCommmands(apicmd, api)
	attachAPINetworkCommmands(apicmd, api)
	attachAPIScriptCommmands(apicmd, api)
	attachAPITransactionsCommmands(apicmd, api)
	attachAPIPoolCommmands(apicmd, api)

	app.Commands = append(app.Commands, apicmd)
}

func apiCommonFlags() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "Set port",
			Value:   uint(koios.DefaultPort),
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
			Name:  "schema",
			Usage: "Set URL schema",
			Value: koios.DefaultSchema,
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
	}
}

func attachAPIAccountCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "account-list",
			Category: "ACCOUNT",
			Usage:    "Get a list of all accounts returns array of stake addresses.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountList(callctx)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "account-info",
			Category:  "ACCOUNT",
			Usage:     "Get the account info of any (payment or staking) address.",
			ArgsUsage: "[account]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "account-info requires single address")
				}
				res, err := api.GetAccountInfo(callctx, koios.Address(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "account-rewards",
			Category:  "ACCOUNT",
			Usage:     "Get the full rewards history (including MIR) for a stake address, or certain epoch if specified.",
			ArgsUsage: "[stake-address]",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:  "epoch",
					Usage: "Filter for earned rewards Epoch Number.",
					Value: uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "account-rewards requires single stake address")
				}
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}
				res, err := api.GetAccountRewards(callctx, koios.StakeAddress(ctx.Args().Get(0)), epoch)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "account-updates",
			Category:  "ACCOUNT",
			Usage:     "Get the account updates (registration, deregistration, delegation and withdrawals).",
			ArgsUsage: "[stake-address]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "account-updates requires single stake address")
				}
				res, err := api.GetAccountUpdates(callctx, koios.StakeAddress(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "account-addresses",
			Category:  "ACCOUNT",
			Usage:     "Get all addresses associated with an account payment or staking address",
			ArgsUsage: "[account]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "account-updates requires single stake or payment address")
				}
				res, err := api.GetAccountAddresses(callctx, koios.StakeAddress(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "account-assets",
			Category:  "ACCOUNT",
			Usage:     "Get the native asset balance of an account.",
			ArgsUsage: "[account]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "account-updates requires single stake or payment address")
				}
				res, err := api.GetAccountAssets(callctx, koios.StakeAddress(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-history",
			Category: "ACCOUNT",
			Usage:    "Get the staking history of an account.",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "account-history requires single stake or payment address")
				}
				res, err := api.GetAccountHistory(callctx, koios.StakeAddress(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIAddressCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:      "address-info",
			Category:  "ADDRESS",
			Usage:     "Get address info - balance, associated stake address (if any) and UTxO set.",
			ArgsUsage: "[address]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "address-info requires single address")
				}
				res, err := api.GetAddressInfo(callctx, koios.Address(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "address-txs",
			Category: "ADDRESS",
			Usage: "Get the transaction hash list of input address array, optionally " +
				"filtering after specified block height (inclusive).",
			ArgsUsage: "[address...]",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:  "after-block-height",
					Usage: "Get transactions after specified block height.",
					Value: uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				var addresses []koios.Address
				for _, a := range ctx.Args().Slice() {
					addresses = append(addresses, koios.Address(a))
				}

				res, err := api.GetAddressTxs(callctx, addresses, ctx.Uint64("after-block-height"))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "address-assets",
			Category:  "ADDRESS",
			Usage:     "Get the list of all the assets (policy, name and quantity) for a given address.",
			ArgsUsage: "[address]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "address-info requires single address")
				}
				res, err := api.GetAddressAssets(callctx, koios.Address(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "credential-txs",
			Category: "ADDRESS",
			Usage: "Get the transaction hash list of input payment credential array, " +
				"optionally filtering after specified block height (inclusive).",
			ArgsUsage: "[address...]",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:  "after-block-height",
					Usage: "Get transactions after specified block height.",
					Value: uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				var credentials []koios.PaymentCredential
				for _, c := range ctx.Args().Slice() {
					credentials = append(credentials, koios.PaymentCredential(c))
				}

				res, err := api.GetCredentialTxs(callctx, credentials, ctx.Uint64("after-block-height"))
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIAssetsCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "asset-list",
			Category: "ASSET",
			Usage:    "Get the list of all native assets (paginated).",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetList(callctx)
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
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetAddressList(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
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
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetInfo(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
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
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetSummary(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
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
					Usage:    "Asset Policy ID in hexadecimal format (hex)",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "Asset Name in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetTxs(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIBlocksCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "blocks",
			Category: "BLOCK",
			Usage:    "Get summarised details about all blocks (paginated - latest first).",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetBlocks(callctx)
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
					Usage:    "Block Hash in hex format to fetch details for",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetBlockInfo(
					callctx,
					koios.BlockHash(ctx.String("block-hash")),
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
					Usage:    "Block Hash in hex format to fetch details for",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetBlockTxHashes(
					callctx,
					koios.BlockHash(ctx.String("block-hash")),
				)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIEpochCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "epoch-info",
			Category: "EPOCH",
			Usage:    "Get the epoch information, all epochs if no epoch specified.",
			Flags: []cli.Flag{
				&cli.Uint64Flag{
					Name:  "epoch",
					Usage: "Epoch Number to fetch details for",
					Value: uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetEpochInfo(callctx, epoch)
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
					Name:  "epoch",
					Usage: "Epoch Number to fetch details for",
					Value: uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetEpochParams(callctx, epoch)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIGeneralCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "get",
			Usage:    "get issues a GET request to the specified API endpoint",
			Category: "UTILS",
			Action: func(ctx *cli.Context) error {
				endpoint := ctx.Args().Get(0)
				if len(endpoint) == 0 {
					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument e.g. /tip")
				}
				res, err := api.GET(callctx, endpoint, nil, nil)
				handleErr(err)
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				handleErr(err)
				printResponseBody(ctx, body)
				return nil
			},
		},
		{
			Name:     "head",
			Usage:    "head issues a HEAD request to the specified API endpoint",
			Category: "UTILS",
			Action: func(ctx *cli.Context) error {
				endpoint := ctx.Args().Get(0)
				if ctx.NArg() == 0 || len(endpoint) == 0 {
					return fmt.Errorf("%w: %s", ErrCommand, "provide endpoint as argument e.g. /tip")
				}
				res, err := api.HEAD(callctx, endpoint, nil, nil)
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

func attachAPINetworkCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "tip",
			Category: "NETWORK",
			Usage:    "Get the tip info about the latest block seen by chain.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTip(callctx)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "genesis",
			Category: "NETWORK",
			Usage:    "Get the Genesis parameters used to start specific era on chain.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetGenesis(callctx)
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
					Name:  "epoch",
					Usage: "Epoch Number to fetch details for",
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint("epoch"))
					epoch = &v
				}

				res, err := api.GetTotals(callctx, epoch)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIScriptCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "script-list",
			Category: "SCRIPT",
			Usage:    "List of all existing script hashes along with their creation transaction hashes.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetScriptList(callctx)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "script-redeemers",
			Category:  "SCRIPT",
			Usage:     "List of all redeemers for a given script hash.",
			ArgsUsage: "[script_hash]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "script-redeemers requires single script-hash as arg")
				}
				res, err := api.GetScriptRedeemers(callctx, koios.ScriptHash(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPITransactionsCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:      "txs-infos",
			Category:  "TRANSACTIONS",
			Usage:     "Get detailed information about transaction(s).",
			ArgsUsage: "[tx-hashes...]",
			Action: func(ctx *cli.Context) error {
				var txs []koios.TxHash
				for _, a := range ctx.Args().Slice() {
					txs = append(txs, koios.TxHash(a))
				}
				res, err := api.GetTxsInfos(callctx, txs)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "tx-info",
			Category:  "TRANSACTIONS",
			Usage:     "Get detailed information about single transaction.",
			ArgsUsage: "[tx-hash]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "tx-info requires single transaction hash")
				}
				res, err := api.GetTxInfo(callctx, koios.TxHash(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "tx-utxos",
			Category:  "TRANSACTIONS",
			Usage:     "Get UTxO set (inputs/apiOutputs) of transactions.",
			ArgsUsage: "[tx-hashes...]",
			Action: func(ctx *cli.Context) error {
				var txs []koios.TxHash
				for _, a := range ctx.Args().Slice() {
					txs = append(txs, koios.TxHash(a))
				}
				res, err := api.GetTxsUTxOs(callctx, txs)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "txs-metadata",
			Category:  "TRANSACTIONS",
			ArgsUsage: "[tx-hashes...]",
			Usage:     "Get metadata information (if any) for given transaction(s).",
			Action: func(ctx *cli.Context) error {
				var txs []koios.TxHash
				for _, a := range ctx.Args().Slice() {
					txs = append(txs, koios.TxHash(a))
				}
				res, err := api.GetTxsMetadata(callctx, txs)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "tx-metadata",
			Category:  "TRANSACTIONS",
			ArgsUsage: "[tx-hash]",
			Usage:     "Get metadata information (if any) for given transaction.",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "tx-metadata requires single transaction hash")
				}
				res, err := api.GetTxMetadata(callctx, koios.TxHash(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "tx-metalabels",
			Category: "TRANSACTIONS",
			Usage:    "Get a list of all transaction metalabels.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetTxMetaLabels(callctx)
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

				txfile, err := ioutil.ReadFile(ctx.Args().Get(0))
				if err != nil {
					return err
				}

				if err = json.Unmarshal(txfile, &stx); err != nil {
					return err
				}
				res, err := api.SubmitSignedTx(callctx, stx)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "txs-statuses",
			Category:  "TRANSACTIONS",
			Usage:     "Get the number of block confirmations for a given transaction hash list",
			ArgsUsage: "[tx-hashes...]",
			Action: func(ctx *cli.Context) error {
				var txs []koios.TxHash
				for _, a := range ctx.Args().Slice() {
					txs = append(txs, koios.TxHash(a))
				}
				res, err := api.GetTxsStatuses(callctx, txs)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "tx-status",
			Category:  "TRANSACTIONS",
			Usage:     "Get the number of block confirmations for a given transaction hash",
			ArgsUsage: "[tx-hash]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "tx-status requires single transaction hash")
				}
				res, err := api.GetTxStatus(callctx, koios.TxHash(ctx.Args().Get(0)))
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}

func attachAPIPoolCommmands(apicmd *cli.Command, api *koios.Client) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "pool-list",
			Category: "POOL",
			Usage:    "A list of all currently registered/retiring (not retired) pools.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolList(callctx)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "pool-infos",
			Category:  "POOL",
			Usage:     "Current pool statuses and details for a specified list of pool ids.",
			ArgsUsage: "[pool-id...]",
			Action: func(ctx *cli.Context) error {
				var pids []koios.PoolID
				for _, pid := range ctx.Args().Slice() {
					pids = append(pids, koios.PoolID(pid))
				}
				res, err := api.GetPoolInfos(callctx, pids)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:      "pool-info",
			Category:  "POOL",
			Usage:     "Current pool status and details for a specified pool by pool id.",
			ArgsUsage: "[pool-id]",
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "pool-info requires single pool id")
				}
				res, err := api.GetPoolInfo(callctx, koios.PoolID(ctx.Args().Get(0)))
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
					Name:  "epoch",
					Usage: "Epoch Number to fetch details for",
					Value: uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "pool-delegators requires single pool id")
				}
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetPoolDelegators(callctx, koios.PoolID(ctx.Args().Get(0)), epoch)
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
					Name:  "epoch",
					Usage: "Epoch Number to fetch details for",
					Value: uint64(0),
				},
			},
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 1 {
					return fmt.Errorf("%w: %s", ErrCommand, "pool-blocks requires single pool id")
				}
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetPoolBlocks(callctx, koios.PoolID(ctx.Args().Get(0)), epoch)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-updates",
			Category: "POOL",
			Usage:    "Return all pool updates for all pools or only updates for specific pool if specified.",
			Action: func(ctx *cli.Context) error {
				var pool *koios.PoolID
				if ctx.NArg() == 1 {
					v := koios.PoolID(ctx.Args().Get(0))
					pool = &v
				}

				res, err := api.GetPoolUpdates(callctx, pool)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-relays",
			Category: "POOL",
			Usage:    "A list of registered relays for all currently registered/retiring (not retired) pools.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolRelays(callctx)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-metadata",
			Category: "POOL",
			Usage:    "Metadata(on & off-chain) for all currently registered/retiring (not retired) pools.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolMetadata(callctx)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}
