// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
)

const categoryPool = "pool"

// Pool:
// https://api.koios.rest/#tag--Pool
func pool(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryPool, "Query information about specific pools")
	cmd.AddSubCommand(cmdPoolPoolList(c))
	cmd.AddSubCommand(cmdPoolPoolInfo(c))
	cmd.AddSubCommand(cmdPoolPoolStakeSnapshot(c))
	cmd.AddSubCommand(cmdPoolPoolDelegators(c))
	cmd.AddSubCommand(cmdPoolPoolDelegatorsHistory(c))
	cmd.AddSubCommand(cmdPoolPoolBlocks(c))
	cmd.AddSubCommand(cmdPoolPoolHistory(c))
	cmd.AddSubCommand(cmdPoolPoolUpdates(c))
	cmd.AddSubCommand(cmdPoolPoolRegistrations(c))
	cmd.AddSubCommand(cmdPoolPoolRetirements(c))
	cmd.AddSubCommand(cmdPoolPoolRelays(c))
	cmd.AddSubCommand(cmdPoolPoolMetadata(c))
}

func cmdPoolPoolList(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_list",
		happy.Option("description", "List of brief info for all pools"),
		happy.Option("category", categoryPool),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Get the list of all pools")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_list

    Example: koios-cli api pool_list
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPoolList(sess, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolInfo(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_info",
		happy.Option("description", "Pool Information"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api pool_info [_pool_bech32_ids...] // max 50"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Current pool statuses and details for a specified list of pool ids")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/pool_info

    Example: koios-cli api pool_info \
      pool100wj94uzf54vup2hdzk0afng4dhjaqggt7j434mtgm8v2gfvfgp \
      pool102s2nqtea2hf5q0s4amj0evysmfnhrn4apyyhd4azcmsclzm96m \
      pool102vsulhfx8ua2j9fwl2u7gv57fhhutc3tp6juzaefgrn7ae35wm
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var poolIDs []koios.PoolID
		for _, id := range args.Args() {
			poolIDs = append(poolIDs, koios.PoolID(id.String()))
		}

		res, err := c.koios().GetPoolInfos(sess, poolIDs, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolStakeSnapshot(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_stake_snapshot",
		happy.Option("description", "Pool Stake Snapshot"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_stake_snapshot _pool_bech32"),
	)

	cmd.AddInfo("Returns Mark, Set and Go stake snapshots for the selected pool, useful for leaderlog calculation")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_stake_snapshot

    Example: koios-cli api pool_stake_snapshot pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPoolStakeSnapshot(sess, koios.PoolID(args.Arg(0).String()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolDelegators(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_delegators")
}

func cmdPoolPoolDelegatorsHistory(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_delegators_history")
}

func cmdPoolPoolBlocks(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_blocks")
}

func cmdPoolPoolHistory(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_history")
}

func cmdPoolPoolUpdates(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_updates")
}

func cmdPoolPoolRegistrations(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_registrations")
}

func cmdPoolPoolRetirements(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_retirements")
}

func cmdPoolPoolRelays(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_relays")
}

func cmdPoolPoolMetadata(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_metadata")
}
