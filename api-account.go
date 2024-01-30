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

func attachAPIAccountCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "account-list",
			Category: "ACCOUNT",
			Usage:    "Get a list of all accounts returns array of stake addresses.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccounts(callctx, opts)
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
				&cli.BoolFlag{
					Value:    false,
					Name:     "cached",
					Usage:    "Get the cached account information for given stake addresses (accounts)",
					Required: false,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountInfo(callctx, koios.Address(ctx.String("address")), ctx.Bool("cached"), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-rewards",
			Category: "ACCOUNT",
			Usage:    "Get the full rewards history (including MIR) for a stake address, or certain epoch if specified.",
			Flags: []cli.Flag{
				epochFlag,
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
				res, err := api.GetAccountRewards(callctx, koios.Address(ctx.String("address")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-updates",
			Category: "ACCOUNT",
			Usage:    "Get the account updates (registration, deregistration, delegation and withdrawals).",
			Flags: []cli.Flag{
				epochFlag,
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
				res, err := api.GetAccountUpdates(callctx, koios.Address(ctx.String("address")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-addresses",
			Category: "ACCOUNT",
			Usage:    "Get all addresses associated with an account payment or staking address",
			Flags: []cli.Flag{
				epochFlag,
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
				res, err := api.GetAccountAddresses(callctx, koios.Address(ctx.String("address")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-assets",
			Category: "ACCOUNT",
			Usage:    "Get the native asset balance of an account.",
			Flags: []cli.Flag{
				epochFlag,
				&cli.StringFlag{
					Name:     "address",
					Aliases:  []string{"a"},
					Usage:    "Cardano payment address in bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetAccountAssets(callctx, koios.Address(ctx.String("address")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "account-history",
			Category: "ACCOUNT",
			Usage:    "Get the staking history of an account.",
			Flags: []cli.Flag{
				epochFlag,
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
				res, err := api.GetAccountHistory(callctx, koios.Address(ctx.String("address")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}
