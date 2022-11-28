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
	"os"

	"github.com/cardano-community/koios-go-client/v3"
	"github.com/urfave/cli/v2"
)

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
