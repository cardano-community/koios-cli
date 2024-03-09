// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

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
	return notimplCmd(categoryScript, "script_info")
}

func cmdScriptNativeScriptList(c *client) *happy.Command {
	return notimplCmd(categoryScript, "native_script_list")
}

func cmdScriptPlutusScriptList(c *client) *happy.Command {
	return notimplCmd(categoryScript, "plutus_script_list")
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
