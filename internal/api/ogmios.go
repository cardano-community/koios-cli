// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package api

import "github.com/happy-sdk/happy"

const categoryOgmios = "ogmios"

// Ogmios:
// https://api.koios.rest/#tag--Ogmios
func ogmios(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryOgmios, "Various stateless queries against Ogmios v6 instance")
	cmd.AddSubCommand(notimplCmd(categoryOgmios, "ogmios"))
}
