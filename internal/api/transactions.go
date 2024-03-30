// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
	"slices"

	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/pkg/vars/varflag"
)

const categoryTransactions = "transactions"

// Transactions:
// https://api.koios.rest/#tag--Transactions
func transactions(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryTransactions, "Query blockchain transaction details")
	cmd.AddSubCommand(cmdTransactionsUtxoInfo(c))
	cmd.AddSubCommand(cmdTransactionsTxInfo(c))
	cmd.AddSubCommand(cmdTransactionsTxMetadata(c))
	cmd.AddSubCommand(cmdTransactionsTxMetalabels(c))
	cmd.AddSubCommand(cmdTransactionsSubmittx(c))
	cmd.AddSubCommand(cmdTransactionsTxStatus(c))
	cmd.AddSubCommand(cmdTransactionsTxUtxos(c))
}

func cmdTransactionsUtxoInfo(c *client) *happy.Command {
	cmd := happy.NewCommand("utxo_info",
		happy.Option("description", "UTxO Info"),
		happy.Option("category", categoryTransactions),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api utxo_info [_utxo_refs...] // max 50"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(
		varflag.BoolFunc("extended", false, "Controls whether or not certain optional fields supported by a given endpoint are populated as a part of the call"),
	))...)

	cmd.AddInfo("Array of Cardano UTxO references in the form \"hash#index\" with extended flag to toggle additional fields")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/utxo_info

    Example: koios-cli api utxo_info \
      f144a8264acf4bdfe2e1241170969c930d64ab6b0996a4a45237b623f1dd670e#0 \
      0b8ba3bed976fa4913f19adc9f6dd9063138db5b4dd29cecde369456b5155e94#0

    Example: koios-cli api utxo_info --extended \
      f144a8264acf4bdfe2e1241170969c930d64ab6b0996a4a45237b623f1dd670e#0 \
      0b8ba3bed976fa4913f19adc9f6dd9063138db5b4dd29cecde369456b5155e94#0
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var utxos []koios.UTxORef
		for _, arg := range args.Args() {
			utxos = append(utxos, koios.UTxORef(arg.String()))
		}

		res, err := c.koios().GetUTxOInfo(sess, utxos, args.Flag("extended").Var().Bool(), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdTransactionsTxInfo(c *client) *happy.Command {
	cmd := happy.NewCommand("tx_info",
		happy.Option("description", "Transaction Information"),
		happy.Option("category", categoryTransactions),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api tx_info [_tx_hashes...] // max 50"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Get detailed information about transaction(s)")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/tx_info

    Example: koios-cli api tx_info \
      f144a8264acf4bdfe2e1241170969c930d64ab6b0996a4a45237b623f1dd670e \
      0b8ba3bed976fa4913f19adc9f6dd9063138db5b4dd29cecde369456b5155e94

    Example: koios-cli api tx_info --page 1 --page-size 5 \
      f144a8264acf4bdfe2e1241170969c930d64ab6b0996a4a45237b623f1dd670e \
      0b8ba3bed976fa4913f19adc9f6dd9063138db5b4dd29cecde369456b5155e94
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var txs []koios.TxHash
		for _, arg := range args.Args() {
			txs = append(txs, koios.TxHash(arg.String()))
		}

		res, err := c.koios().GetTxInfo(sess, txs, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdTransactionsTxMetadata(c *client) *happy.Command {
	return notimplCmd(categoryTransactions, "tx_metadata")
}

func cmdTransactionsTxMetalabels(c *client) *happy.Command {
	return notimplCmd(categoryTransactions, "tx_metalabels")
}

func cmdTransactionsSubmittx(c *client) *happy.Command {
	return notimplCmd(categoryTransactions, "submittx")
}

func cmdTransactionsTxStatus(c *client) *happy.Command {
	return notimplCmd(categoryTransactions, "tx_status")
}

func cmdTransactionsTxUtxos(c *client) *happy.Command {
	return notimplCmd(categoryTransactions, "tx_utxos")
}
