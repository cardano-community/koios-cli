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
	cmd := happy.NewCommand("account_list",
		happy.Option("description", "Account List"),
		happy.Option("category", categoryStakeAccount),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Get a list of all stake addresses that have atleast 1 transaction")
	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/account_list

    Example: koios-cli api account_list

    {
      "data": [
        "stake1uyfmzu5qqy70a8kq4c8rw09q0w0ktfcxppwujejnsh6tyrg5c774g",
        "stake1uy9crcqratu65rklv0v7eyt4hnkpewkjgtwgwmkwzl573msyk9gjl",
        "stake1uydhlh7f2kkw9eazct5zyzlrvj32gjnkmt2v5qf6t8rut4qwch8ey",
        ...
      ]
    }

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetAccountList(sess, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
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
