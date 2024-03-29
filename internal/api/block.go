// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
)

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
	cmd := happy.NewCommand("blocks",
		happy.Option("description", "Block List"),
		happy.Option("category", categoryBlock),
	).WithFalgs(pagingFlags...)

	cmd.AddInfo("Get summarised details about all blocks (paginated - latest first)")
	cmd.AddInfo("Docs: https://api.koios.rest/#get-/blocks")

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetBlocks(sess, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdBlockBlockInfo(c *client) *happy.Command {
	cmd := happy.NewCommand("block-info",
		happy.Option("description", "Block Info"),
		happy.Option("category", categoryBlock),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
	).WithFalgs(pagingFlags...)

	cmd.AddInfo("Get detailed information about a blocks")
	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/block_info

    Example: koios-cli api block-info \
      fb9087c9f1408a7bbd7b022fd294ab565fec8dd3a8ef091567482722a1fa4e30 \
      60188a8dcb6db0d80628815be2cf626c4d17cb3e826cebfca84adaff93ad492a \
      c6646214a1f377aa461a0163c213fc6b86a559a2d6ebd647d54c4eb00aaab015
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var hashes []koios.BlockHash
		for _, arg := range args.Args() {
			hashes = append(hashes, koios.BlockHash(arg.String()))
		}
		res, err := c.koios().GetBlockInfos(sess, hashes, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdBlockBlockTxs(c *client) *happy.Command {
	return notimplCmd(categoryBlock, "block_txs")
}
