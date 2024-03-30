// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
)

const categoryScript = "script"

// Script:
// https://api.koios.rest/#tag--Script
func script(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryScript, "Query information about specific scripts (Smart Contracts)")
	cmd.AddSubCommand(cmdScriptScriptInfo(c))
	cmd.AddSubCommand(cmdScriptNativeScriptList(c))
	cmd.AddSubCommand(cmdScriptPlutusScriptList(c))
	cmd.AddSubCommand(cmdScriptScriptRedeemers(c))
	cmd.AddSubCommand(cmdScriptScriptUtxos(c))
	cmd.AddSubCommand(cmdScriptDatumInfo(c))
}

func cmdScriptScriptInfo(c *client) *happy.Command {
	cmd := happy.NewCommand("script_info",
		happy.Option("description", "Script Information"),
		happy.Option("category", categoryScript),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api script_info [_script_hashes...] // max 50"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("List of script information for given script hashes")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#post-/script_info

    Example: koios-cli api script_info \
      bd2119ee2bfb8c8d7c427e8af3c35d537534281e09e23013bca5b138 \
      c0c671fba483641a71bb92d3a8b7c52c90bf1c01e2b83116ad7d4536
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		var hashes []koios.ScriptHash
		for _, arg := range args.Args() {
			hashes = append(hashes, koios.ScriptHash(arg.String()))
		}

		res, err := c.koios().GetScriptInfo(sess, hashes, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdScriptNativeScriptList(c *client) *happy.Command {
	cmd := happy.NewCommand("native_script_list",
		happy.Option("description", "Native Script List"),
		happy.Option("category", categoryScript),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("List of all existing native script hashes along with their creation transaction hashes.")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/native_script_list

    Example: koios-cli api native_script_list
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetNativeScripts(sess, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdScriptPlutusScriptList(c *client) *happy.Command {
	cmd := happy.NewCommand("plutus_script_list",
		happy.Option("description", "Plutus Script List"),
		happy.Option("category", categoryScript),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("List of all existing Plutus script hashes along with their creation transaction hashes.")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/plutus_script_list

    Example: koios-cli api plutus_script_list
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		res, err := c.koios().GetPlutusScripts(sess, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdScriptScriptRedeemers(c *client) *happy.Command {
	return notimplCmd(categoryScript, "script_redeemers")
}

func cmdScriptScriptUtxos(c *client) *happy.Command {
	return notimplCmd(categoryScript, "script_utxos")
}

func cmdScriptDatumInfo(c *client) *happy.Command {
	return notimplCmd(categoryScript, "datum_info")
}
