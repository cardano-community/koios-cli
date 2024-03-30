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
	cmd := happy.NewCommand("script_redeemers",
		happy.Option("description", "Script Redeemers"),
		happy.Option("category", categoryScript),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api script_redeemers [_script_hash]"),
	).WithFlags(pagingFlags...)

	cmd.AddInfo("List of all redeemers for a given script hash.")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/script_redeemers

    Example: koios-cli api script_redeemers d8480dc869b94b80e81ec91b0abe307279311fe0e7001a9488f61ff8

    Example: koios-cli api script_redeemers d8480dc869b94b80e81ec91b0abe307279311fe0e7001a9488f61ff8 --page 1 --page-size 3
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		hash := koios.ScriptHash(args.Arg(0).String())
		res, err := c.koios().GetScriptRedeemers(sess, hash, opts)
		apiOutput(c.noFormat, res, err)
		return err

	})

	return cmd
}

func cmdScriptScriptUtxos(c *client) *happy.Command {
	cmd := happy.NewCommand("script_utxos",
		happy.Option("description", "Script Utxos"),
		happy.Option("category", categoryScript),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api script_utxos [_script_hash]"),
	).WithFlags(
		slices.Concat(
			pagingFlags,
			flagSlice(
				varflag.BoolFunc("extended", false, "Controls whether or not certain optional fields supported by a given endpoint are populated as a part of the call"),
			))...)

	cmd.AddInfo("List of all UTXOs for a given script hash.")

	cmd.AddInfo(`
    Docs: https://api.koios.rest/#get-/script_utxos

    Example: koios-cli api script_utxos d8480dc869b94b80e81ec91b0abe307279311fe0e7001a9488f61ff8
    Example: koios-cli api script_utxos d8480dc869b94b80e81ec91b0abe307279311fe0e7001a9488f61ff8 --extended
    Example: koios-cli api script_utxos d8480dc869b94b80e81ec91b0abe307279311fe0e7001a9488f61ff8 --page 1 --page-size 3

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}

		hash := koios.ScriptHash(args.Arg(0).String())
		res, err := c.koios().GetScriptUtxos(sess, hash, args.Flag("extended").Var().Bool(), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdScriptDatumInfo(c *client) *happy.Command {
	return notimplCmd(categoryScript, "datum_info")
}
