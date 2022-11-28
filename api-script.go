// Copyright 2022 The Cardano Community Authors
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//   or LICENSE file in repository root.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/cardano-community/koios-go-client/v3"
	"github.com/urfave/cli/v2"
)

func attachAPIScriptCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "native-script-list",
			Category: "SCRIPT",
			Usage:    "List of all existing native script hashes along with their creation transaction hashes.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetNativeScripts(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "plutus-script-list",
			Category: "SCRIPT",
			Usage:    "List of all existing Plutus script hashes along with their creation transaction hashes.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPlutusScripts(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "script-redeemers",
			Category: "SCRIPT",
			Usage:    "List of all redeemers for a given script hash.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "script-hash",
					Aliases:  []string{"s"},
					Usage:    "Script hash in hexadecimal format (hex)",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetScriptRedeemers(callctx, koios.ScriptHash(ctx.String("script-hash")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}
