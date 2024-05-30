// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package main

import (
	"github.com/cardano-community/koios-cli/v2/internal/api"
	"github.com/cardano-community/koios-cli/v2/internal/auth"
	"github.com/cardano-community/koios-cli/v2/koios"
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/sdk/logging"
)

func main() {
	logOpts := logging.ConsoleDefaultOptions()
	logOpts.Level = logging.LevelOk
	logOpts.AddSource = false

	app := happy.New(happy.Settings{
		Name: "KOIOS CLI",
		Slug: "koios",
		Description: `
      Koios is a distributed & open-source public API query layer for Cardano,
      that is elastic in nature and addresses ever-demanding requirements from Cardano Blockchain.
      It provides a easy to query RESTful layer that has a lot of flexibility to cater for
      different data consumption requirements from blockchain.`,
		CopyrightBy:    "The Cardano Community",
		CopyrightSince: 2022,
		License:        "Apache-2.0",
		TimeLocation:   "Local",
		StatsEnabled:   false,
	}).
		WithLogger(logging.Console(logOpts)).
		WithBrand(koios.Brand()).
		WithCommand(api.Command()).
		WithCommand(auth.Command())

	app.AddInfo(`
      Example: Usage with Public Tier
        koios-cli api tip

      Example: Usage with subscription token without saving token to disk
        koios-cli api --auth <jwt-token> tip

      Example: Usage with saved profile (see koios-cli auth -h)
        koios-cli --profile <project-id> api tip

      Example: Usage with saved profile and stats
        koios-cli --profile <project-id> api --stats tip
    `)

	app.Run()
}
