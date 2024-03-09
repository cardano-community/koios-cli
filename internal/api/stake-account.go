// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

const categoryStakeAccount = "stake account"

// Stake Account:
// https://api.koios.rest/#tag--Stake-Account
func stakeAccount(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryStakeAccount, "Query details about specific stake account addresses")
	cmd.AddSubCommand(cmdStakeAccountAccountList(c))
	cmd.AddSubCommand(cmdStakeAccountAccountInfo(c))
	cmd.AddSubCommand(cmdStakeAccountAccountInfoCached(c))
	cmd.AddSubCommand(cmdStakeAccountAccountUtxos(c))
	cmd.AddSubCommand(cmdStakeAccountAccountTxs(c))
	cmd.AddSubCommand(cmdStakeAccountAccountRewards(c))
	cmd.AddSubCommand(cmdStakeAccountAccountUpdates(c))
	cmd.AddSubCommand(cmdStakeAccountAccountAddresses(c))
	cmd.AddSubCommand(cmdStakeAccountAccountAssets(c))
	cmd.AddSubCommand(cmdStakeAccountAccountHistory(c))

}

func cmdStakeAccountAccountList(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_list")
}

func cmdStakeAccountAccountInfo(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_info")
}

func cmdStakeAccountAccountInfoCached(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_info_cached")
}

func cmdStakeAccountAccountUtxos(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_utxos")
}

func cmdStakeAccountAccountTxs(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_txs")
}

func cmdStakeAccountAccountRewards(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_rewards")
}

func cmdStakeAccountAccountUpdates(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_updates")
}

func cmdStakeAccountAccountAddresses(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_addresses")
}

func cmdStakeAccountAccountAssets(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_assets")
}

func cmdStakeAccountAccountHistory(c *client) *happy.Command {
	return notimplCmd(categoryStakeAccount, "account_history")
}
