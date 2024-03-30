// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
	"slices"

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
		happy.Option("description", "Pool List"),
		happy.Option("category", categoryPool),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("List of brief info for all pools")

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
		happy.Option("usage", "koios api pool_stake_snapshot [_pool_bech32]"),
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
	cmd := happy.NewCommand("pool_delegators",
		happy.Option("description", "Pool Delegators"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_delegators [_pool_bech32]"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Return information about live delegators for a given pool.")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_delegators

    Example: koios-cli api pool_delegators pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc

    `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPoolDelegators(sess, koios.PoolID(args.Arg(0).String()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolDelegatorsHistory(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_delegators_history",
		happy.Option("description", "Pool Delegators History"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_delegators_history [_pool_bech32]"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(epochNoFlag))...)

	cmd.AddInfo("Return information about active delegators (incl. history) for a given pool and epoch number (all epochs if not specified).")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_delegators_history

    Example: koios-cli api pool_delegators_history pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc
    Example: koios-cli api pool_delegators_history pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc --epoch 320

    `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPoolDelegatorsHistory(sess, koios.PoolID(args.Arg(0).String()), koios.EpochNo(args.Flag("epoch").Var().Uint()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolBlocks(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_blocks",
		happy.Option("description", "Pool Blocks"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_blocks [_pool_bech32]"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(epochNoFlag))...)

	cmd.AddInfo("Return information about blocks minted by a given pool for all epochs (or _epoch_no if provided)")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_blocks

    Example: koios-cli api pool_blocks pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc
    Example: koios-cli api pool_blocks pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc --epoch 320

    `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPoolBlocks(sess, koios.PoolID(args.Arg(0).String()), koios.EpochNo(args.Flag("epoch").Var().Uint()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolHistory(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_history",
		happy.Option("description", "Pool History"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_history [_pool_bech32]"),
	).WithFlags(slices.Concat(pagingFlags, flagSlice(epochNoFlag))...)

	cmd.AddInfo("Return information about pool stake, block and reward history in a given epoch _epoch_no (or all epochs that pool existed for, in descending order if no _epoch_no was provided)")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_history

    Example: koios-cli api pool_history pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc
    Example: koios-cli api pool_history pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc --epoch 320

    `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPoolHistory(sess, koios.PoolID(args.Arg(0).String()), koios.EpochNo(args.Flag("epoch").Var().Uint()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolUpdates(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_updates",
		happy.Option("description", "Pool Updates (History)"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_updates [_pool_bech32]"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Return all pool updates for all pools or only updates for specific pool if specified")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_updates

    Example: koios-cli api pool_updates
    Example: koios-cli api pool_updates pool155efqn9xpcf73pphkk88cmlkdwx4ulkg606tne970qswczg3asc
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPoolUpdates(sess, koios.PoolID(args.Arg(0).String()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolRegistrations(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_registrations",
		happy.Option("description", "Pool Registrations"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 0),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_registrations [_epoch_no]"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Return all pool registrations initiated in the requested epoch")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_registrations

    Example: koios-cli api pool_registrations
    Example: koios-cli api pool_registrations 320
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		// 0 when value is invalid
		epochNo, _ := args.Arg(0).Uint()

		res, err := c.koios().GetPoolRegistrations(sess, koios.EpochNo(epochNo), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolRetirements(c *client) *happy.Command {
	cmd := happy.NewCommand("pool_retirements",
		happy.Option("description", "Pool Retirements"),
		happy.Option("category", categoryPool),
		happy.Option("argn.min", 0),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api pool_retirements [_epoch_no]"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("Return all pool retirements initiated in the requested epoch")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/pool_retirements

    Example: koios-cli api pool_retirements
    Example: koios-cli api pool_retirements 320

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		// 0 when value is invalid
		epochNo, _ := args.Arg(0).Uint()

		res, err := c.koios().GetPoolRetirements(sess, koios.EpochNo(epochNo), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdPoolPoolRelays(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_relays")
}

func cmdPoolPoolMetadata(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_metadata")
}
