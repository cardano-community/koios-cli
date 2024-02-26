// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

const categoryBlock = "block"

// Block:
// https://api.koios.rest/#tag--Block
func block(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryBlock, "Query information about particular block on chain")
	cmd.AddSubCommand(cmdBlockBlocks(c))
	cmd.AddSubCommand(cmdBlockBlockInfo(c))
	cmd.AddSubCommand(cmdBlockBlockTxs(c))
}

func cmdBlockBlocks(c *client) *happy.Command {
	return notimplCmd(categoryBlock, "blocks")
}

func cmdBlockBlockInfo(c *client) *happy.Command {
	return notimplCmd(categoryBlock, "block_info")
}

func cmdBlockBlockTxs(c *client) *happy.Command {
	return notimplCmd(categoryBlock, "block_txs")
}
