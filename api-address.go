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
