// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package api

import (
	"slices"

	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
)

func (c *client) cmdNetworkTip() *happy.Command {
	cmd := happy.NewCommand("tip",
		happy.Option("description", "Query Chain Tip"),
		happy.Option("category", "network"),
	).WithFalgs(queryFlag)
	cmd.AddInfo("Get the tip info about the latest block seen by chain")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/tip

  Example: koios-cli api tip
    {
      ...
      "data": {
       "abs_slot": 117326575,
       "block_no": 9979917,
       "block_time": 1708892866,
       "epoch_no": 469,
       "epoch_slot": 81775,
       "hash": "7b934bad3611a5c7f8500b56a273c6467c120771cc60f9a1bd3ace2a9a81e548"
      }
     }

  Example: koios-cli api tip --query="select=epoch_no"
    {
      ...
      "data": {
        "abs_slot": null,
        "block_no": null,
        "block_time": null,
        "epoch_no": 469,
        "epoch_slot": null,
        "hash": null
      }
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.kc.GetTip(sess, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func (c *client) cmdNetworkGenesis() *happy.Command {
	cmd := happy.NewCommand("genesis",
		happy.Option("description", "Get Genesis info"),
		happy.Option("category", "network"),
	).WithFalgs(queryFlag)
	cmd.AddInfo("Get the Genesis parameters used to start specific era on chain")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/genesis

  Example: koios-cli api genesis --query="select=networkmagic,networkid"
    {
      ...
      "data": {
        "activeslotcoeff": "0",
        "alonzogenesis": null,
        "epochlength": "0",
        "maxkesrevolutions": "0",
        "maxlovelacesupply": "0",
        "networkid": "Mainnet",
        "networkmagic": "764824073",
        "securityparam": "0",
        "slotlength": "0",
        "slotsperkesperiod": "0",
        "systemstart": null,
        "updatequorum": "0"
      }
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.kc.GetGenesis(sess, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func (c *client) cmdNetworkTotals() *happy.Command {
	cmd := happy.NewCommand("totals",
		happy.Option("description", "Get historical tokenomic stats for the network."),
		happy.Option("category", "network"),
	).WithFalgs(
		slices.Concat(clientPagingFlags, flagSlice(epochFlag))...,
	)
	cmd.AddInfo(`
  Get the circulating utxo, treasury rewards, supply and reserves in lovelace
  for specified epoch, all epochs if --epoch is not set`)
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/totals

  Example: koios-cli api totals
  Example: koios-cli api totals --page 1 --page-size 5
  Example: koios-cli api totals --epoch 469
    {
      ...
      "data": [
        {
         "circulation": "34463774233630959",
         "epoch_no": 469,
         "reserves": "8343083956124781",
         "reward": "688313147260771",
         "supply": "36656916043875219",
         "treasury": "1500366891487428"
        }
      ]
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		var epoch *koios.EpochNo

		if args.Flag("epoch").Present() {
			if e := args.Flag("epoch").Var().Uint(); e > 0 {
				v := koios.EpochNo(e)
				epoch = &v
			}
		}

		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		res, err := c.kc.GetTotals(sess, epoch, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func (c *client) cmdNetworkParamUpdates() *happy.Command {
	cmd := happy.NewCommand("param_updates",
		happy.Option("description", "Param Update Proposals"),
		happy.Option("category", "network"),
	).WithFalgs(
		slices.Concat(clientPagingFlags, flagSlice(queryFlag))...,
	)
	cmd.AddInfo("Get all parameter update proposals submitted to the chain starting Shelley era")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/param_updates

  Example:
    koios-cli api param_updates --page 1 --page-size 3
    koios-cli api param_updates --query="order=block_height.desc&limit=1"
    koios-cli api param_updates --query="epoch_no=eq.444"
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.kc.GetParamUpdates(sess, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func (c *client) cmdNetworkReserveWithdrawals() *happy.Command {
	cmd := happy.NewCommand("reserve_withdrawals",
		happy.Option("description", "Reserve Withdrawals"),
		happy.Option("category", "network"),
	).WithFalgs(
		slices.Concat(clientPagingFlags, flagSlice(queryFlag))...,
	)
	cmd.AddInfo("List of all withdrawals from reserves against stake accounts")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/reserve_withdrawals

  Example:
    koios-cli api reserve_withdrawals --page 1 --page-size 3
    koios-cli api reserve_withdrawals --page-size 3 --query="epoch_no=eq.285"
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		res, err := c.kc.GetReserveWithdrawals(sess, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func (c *client) cmdNetworkTreasuryWithdrawals() *happy.Command {
	cmd := happy.NewCommand("treasury_withdrawals",
		happy.Option("description", "Treasury Withdrawals"),
		happy.Option("category", "network"),
	).WithFalgs(
		slices.Concat(clientPagingFlags, flagSlice(queryFlag))...,
	)
	cmd.AddInfo("List of all withdrawals from treasury against stake accounts")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/treasury_withdrawals

  Example:
    koios-cli api treasury_withdrawals --page 1 --page-size 3
    koios-cli api treasury_withdrawals --page 1 --page-size 3 --query "order=block_height.desc"
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		res, err := c.kc.GetTreasuryWithdrawals(sess, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}
