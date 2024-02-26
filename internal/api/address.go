// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

const categoryAddress = "address"

// Address:
// https://api.koios.rest/#tag--Address
func address(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryAddress, "Query information about specific address(es)")
	cmd.AddSubCommand(cmdAddressAddressInfo(c))
	cmd.AddSubCommand(cmdAddressAddressUtxos(c))
	cmd.AddSubCommand(cmdAddressCredentialUtxos(c))
	cmd.AddSubCommand(cmdAddressAddressTxs(c))
	cmd.AddSubCommand(cmdAddressCredentialTxs(c))
	cmd.AddSubCommand(cmdAddressAddressAssets(c))
}

func cmdAddressAddressInfo(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "address_info")
}

func cmdAddressAddressUtxos(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "address_utxos")
}

func cmdAddressCredentialUtxos(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "credential_utxos")
}

func cmdAddressAddressTxs(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "address_txs")
}

func cmdAddressCredentialTxs(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "credential_txs")
}

func cmdAddressAddressAssets(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "address_assets")
}
