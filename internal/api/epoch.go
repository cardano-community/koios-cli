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
	cmd := happy.NewCommand("epoch_info",
		happy.Option("description", "Epoch Information"),
		happy.Option("category", categoryEpoch),
		happy.Option("argn.min", 0),
		happy.Option("argn.max", 1),
	).WithFlags(
		slices.Concat(
			pagingFlags,
			flagSlice(
				varflag.BoolFunc("include-next-epoch", false, "Include information about nearing but not yet started epoch, to get access to active stake snapshot information if available"),
			),
		)...,
	)

	cmd.AddInfo("Get the epoch information, all epochs if no epoch specified")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/epoch_info

    Example: koios-cli api epoch_info
    Example: koios-cli api epoch_info 320
    Example: koios-cli api epoch_info --include-next-epoch
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		// 0  when value is invalid
		epochNo, _ := args.Arg(0).Uint()
		res, err := c.koios().GetEpochInfo(sess, koios.EpochNo(epochNo), args.Flag("include-next-epoch").Var().Bool(), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdEpochParams(c *client) *happy.Command {
	cmd := happy.NewCommand("epoch_params",
		happy.Option("description", "Epoch Parameters"),
		happy.Option("category", categoryEpoch),
		happy.Option("argn.min", 0),
		happy.Option("argn.max", 1),
	).WithFlags(
		pagingFlags...,
	)

	cmd.AddInfo("Get the protocol parameters for specific epoch, returns information about all epochs if no epoch specified")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/epoch_params

    Example: koios-cli api epoch_params
    Example: koios-cli api epoch_params 320
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		// 0  when value is invalid
		epochNo, _ := args.Arg(0).Uint()
		res, err := c.koios().GetEpochParams(sess, koios.EpochNo(epochNo), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdEpochBlockProtocols(c *client) *happy.Command {
	cmd := happy.NewCommand("epoch_block_protocols",
		happy.Option("description", "Epoch Block Protocols"),
		happy.Option("category", categoryEpoch),
		happy.Option("argn.min", 0),
		happy.Option("argn.max", 1),
	)

	cmd.AddInfo("Get the information about block protocol distribution in epoch")
	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/epoch_block_protocols

    Example: koios-cli api epoch_block_protocols
    Example: koios-cli api epoch_block_protocols 320
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		// 0  when value is invalid
		epochNo, _ := args.Arg(0).Uint()
		res, err := c.koios().GetEpochBlockProtocols(sess, koios.EpochNo(epochNo), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}
