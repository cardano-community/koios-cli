// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

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
	return notimplCmd(categoryTransactions, "utxo_info")
}

func cmdTransactionsTxInfo(c *client) *happy.Command {
	return notimplCmd(categoryTransactions, "tx_info")
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
