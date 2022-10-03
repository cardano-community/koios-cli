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
	"github.com/cardano-community/koios-go-client/v2"
	"github.com/urfave/cli/v2"
)

func attachAPIAssetsCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "asset-list",
			Category: "ASSET",
			Usage:    "Get the list of all native assets (paginated).",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssets(callctx, opts)
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
				res, err := api.GetAssetAddresses(
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
				&cli.Uint64Flag{
					Name:     "after-block-height",
					Usage:    "Only fetch information after specific block height",
					Required: false,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAssetTxs(
					callctx,
					koios.PolicyID(ctx.String("policy")),
					koios.AssetName(ctx.String("name")),
					ctx.Int("after-block-height"),
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
