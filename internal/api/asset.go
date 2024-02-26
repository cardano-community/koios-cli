// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

const categoryAsset = "asset"

// Asset:
// https://api.koios.rest/#tag--Asset
func asset(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryAsset, "Query Asset related informations")
	cmd.AddSubCommand(cmdAssetAssetList(c))
	cmd.AddSubCommand(cmdAssetPolicyAssetList(c))
	cmd.AddSubCommand(cmdAssetAssetTokenRegistry(c))
	cmd.AddSubCommand(cmdAssetAssetInfo(c))
	cmd.AddSubCommand(cmdAssetAssetUtxos(c))
	cmd.AddSubCommand(cmdAssetAssetHistory(c))
	cmd.AddSubCommand(cmdAssetAssetAddresses(c))
	cmd.AddSubCommand(cmdAssetAssetNftAddress(c))
	cmd.AddSubCommand(cmdAssetPolicyAssetAddresses(c))
	cmd.AddSubCommand(cmdAssetPolicyAssetInfo(c))
	cmd.AddSubCommand(cmdAssetAssetSummary(c))
	cmd.AddSubCommand(cmdAssetAssetTxs(c))
}

func cmdAssetAssetList(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_list")
}

func cmdAssetPolicyAssetList(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "policy_asset_list")
}

func cmdAssetAssetTokenRegistry(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_token_registry")
}

func cmdAssetAssetInfo(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_info")
}

func cmdAssetAssetUtxos(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_utxos")
}

func cmdAssetAssetHistory(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_history")
}

func cmdAssetAssetAddresses(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_addresses")
}

func cmdAssetAssetNftAddress(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_nft_address")
}

func cmdAssetPolicyAssetAddresses(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "policy_asset_addresses")
}

func cmdAssetPolicyAssetInfo(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "policy_asset_info")
}

func cmdAssetAssetSummary(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_summary")
}

func cmdAssetAssetTxs(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_txs")
}
