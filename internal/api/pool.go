// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

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
	return notimplCmd(categoryPool, "pool_info")
}

func cmdPoolPoolStakeSnapshot(c *client) *happy.Command {
	return notimplCmd(categoryPool, "pool_stake_snapshot")
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
