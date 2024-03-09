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
	return notimplCmd(categoryPool, "pool_list")
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
