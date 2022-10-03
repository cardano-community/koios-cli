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
	"github.com/cardano-community/koios-go-client/v2"
	"github.com/urfave/cli/v2"
)

func attachAPIEpochCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "epoch-info",
			Category: "EPOCH",
			Usage:    "Get the epoch information, all epochs if no epoch specified.",
			Flags: []cli.Flag{
				epochFlag,
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetEpochInfo(callctx, epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "epoch-params",
			Category: "EPOCH",
			Usage: "Get the protocol parameters for specific epoch, " +
				"returns information about all epochs if no epoch specified.",
			Flags: []cli.Flag{
				epochFlag,
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetEpochParams(callctx, epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "epoch-block-protocols",
			Category: "EPOCH",
			Usage:    "Get the information about block protocol distribution in epoch",
			Flags: []cli.Flag{
				epochFlag,
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetEpochBlockProtocols(callctx, epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}
