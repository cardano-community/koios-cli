// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

const categoryEpoch = "epoch"

// Epoch:
// https://api.koios.rest/#tag--Epoch
func epoch(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryEpoch, "Query epoch-specific details")
	cmd.AddSubCommand(cmdEpochInfo(c))
	cmd.AddSubCommand(cmdEpochParams(c))
	cmd.AddSubCommand(cmdEpochBlockProtocols(c))
}

func cmdEpochInfo(c *client) *happy.Command {
	return notimplCmd(categoryEpoch, "epoch_info")
}

func cmdEpochParams(c *client) *happy.Command {
	return notimplCmd(categoryEpoch, "epoch_params")
}

func cmdEpochBlockProtocols(c *client) *happy.Command {
	return notimplCmd(categoryEpoch, "epoch_block_protocols")
}
