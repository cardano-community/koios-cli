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
	cmd := happy.NewCommand("account_info",
		happy.Option("description", "Account Information"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_info_cached [_stake_addresses...] // max 50"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Get the account information for given stake addresses")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/account_info

    Example: koios-cli api account_info \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address
		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		res, err := c.koios().GetAccountInfo(sess, addresses, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountInfoCached(c *client) *happy.Command {
	cmd := happy.NewCommand("account_info_cached",
		happy.Option("description", "Account Information Cached"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_info_cached [_stake_addresses...] // max 50"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Get the cached account information for given stake addresses (effective for performance query against registered accounts)")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/account_info_cached

    Example: koios-cli api account_info_cached \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address
		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		res, err := c.koios().GetAccountInfoCached(sess, addresses, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountUtxos(c *client) *happy.Command {
	cmd := happy.NewCommand("account_utxos",
		happy.Option("description", "Account Utxos"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_utxos [_stake_addresses...] // max 50"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(
		varflag.BoolFunc("extended", false, "Controls whether or not certain optional fields supported by a given endpoint are populated as a part of the call"),
	))...)

	cmd.AddInfo("Get a list of all UTxOs for given stake addresses (account)s")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/account_utxos

    Example: koios-cli api account_utxos \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address
		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		res, err := c.koios().GetAccountUtxos(sess, addresses, args.Flag("extended").Var().Bool(), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountTxs(c *client) *happy.Command {
	cmd := happy.NewCommand("account_txs",
		happy.Option("description", "Account Transactions"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api account_txs [_stake_address]"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(afterBlockHeightFlag))...)

	cmd.AddInfo("Get a list of all transactions for given stake address (account)")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/account_txs

    Example: koios-cli api account_txs stake1u8yxtugdv63wxafy9d00nuz6hjyyp4qnggvc9a3vxh8yl0ckml2uz
    Example: koios-cli api account_txs stake1u8yxtugdv63wxafy9d00nuz6hjyyp4qnggvc9a3vxh8yl0ckml2uz --after-block-height 50000
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		address := koios.Address(args.Arg(0).String())
		res, err := c.koios().GetAccountTxs(sess, address, args.Flag("after-block-height").Var().Uint64(), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountRewards(c *client) *happy.Command {
	cmd := happy.NewCommand("account_rewards",
		happy.Option("description", "Account Rewards"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_rewards [_stake_addresses...] // max 50"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(epochNoFlag))...)

	cmd.AddInfo("Get the full rewards history (including MIR) for given stake addresses")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/account_rewards

    Example: koios-cli api account_rewards \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy \

    Example: koios-cli api account_rewards --epoch 409 \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address
		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		res, err := c.koios().GetAccountRewards(sess, addresses, koios.EpochNo(args.Flag("epoch").Var().Uint64()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountUpdates(c *client) *happy.Command {
	cmd := happy.NewCommand("account_updates",
		happy.Option("description", "Account Updates"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_updates [_stake_addresses...] // max 50"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Get the account updates (registration, deregistration, delegation and withdrawals) for given stake addresses")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/account_updates

    Example: koios-cli api account_updates \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address

		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		res, err := c.koios().GetAccountUpdates(sess, addresses, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountAddresses(c *client) *happy.Command {
	cmd := happy.NewCommand("account_addresses",
		happy.Option("description", "Account Addresses"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_addresses [_stake_addresses...] // max 50"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(
		varflag.BoolFunc("first-only", false, "Only return the first result"),
		varflag.BoolFunc("empty", false, "Include zero quantity entries"),
	))...)

	cmd.AddInfo("Get all addresses associated with given staking accounts")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/account_addresses

    Example: koios-cli api account_addresses \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy

    Example: koios-cli api account_addresses --first-only --empty \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address
		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		res, err := c.koios().GetAccountAddresses(sess, addresses, args.Flag("first-only").Var().Bool(), args.Flag("empty").Var().Bool(), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountAssets(c *client) *happy.Command {
	cmd := happy.NewCommand("account_assets",
		happy.Option("description", "Account Assets"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_assets [_stake_addresses...] // max 50"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Get the native asset balance for a given stake address(es)")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/account_assets

    Example: koios-cli api account_assets \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address
		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		res, err := c.koios().GetAccountAssets(sess, addresses, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdStakeAccountAccountHistory(c *client) *happy.Command {
	cmd := happy.NewCommand("account_history",
		happy.Option("description", "Account History"),
		happy.Option("category", categoryStakeAccount),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api account_history [_stake_addresses...] // max 50"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(epochNoFlag))...)

	cmd.AddInfo("Get the staking history of given stake addresses (accounts)")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/account_history

    Example: koios-cli api account_history \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy

    Example: koios-cli api account_history --epoch 409 \
      stake1uyrx65wjqjgeeksd8hptmcgl5jfyrqkfq0xe8xlp367kphsckq250 \
      stake1uxpdrerp9wrxunfh6ukyv5267j70fzxgw0fr3z8zeac5vyqhf9jhy

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var addresses []koios.Address
		for _, arg := range args.Args() {
			addresses = append(addresses, koios.Address(arg.String()))
		}

		var epochNo koios.EpochNo
		if args.Flag("epoch").Present() {
			epochNo = koios.EpochNo(args.Flag("epoch").Var().Uint64())
		}

		res, err := c.koios().GetAccountHistory(sess, addresses, &epochNo, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}
