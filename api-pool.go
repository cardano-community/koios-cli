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

func attachAPIPoolCommmands(apicmd *cli.Command) {
	apicmd.Subcommands = append(apicmd.Subcommands, []*cli.Command{
		{
			Name:     "pool-list",
			Category: "POOL",
			Usage:    "A list of all currently registered/retiring (not retired) pools.",
			Flags: []cli.Flag{
				epochFlag,
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPools(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-infos",
			Category: "POOL",
			Usage:    "Current pool statuses and details for a specified list of pool ids.",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format, can be used multiple times for list of transactions",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var pids []koios.PoolID
				for _, pid := range ctx.StringSlice("pool-id") {
					pids = append(pids, koios.PoolID(pid))
				}
				res, err := api.GetPoolInfos(callctx, pids, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-info",
			Category: "POOL",
			Usage:    "Current pool status and details for a specified pool by pool id.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolInfo(callctx, koios.PoolID(ctx.String("pool-id")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-delegators",
			Category: "POOL",
			Usage:    "Return information about delegators by a given pool and optional epoch (current if omitted).",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolDelegators(callctx, koios.PoolID(ctx.String("pool-id")), opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-history",
			Category: "POOL",
			Usage:    "Return information about delegators by a given pool and optional epoch (current if omitted).",
			Flags: []cli.Flag{
				epochFlag,
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetPoolHistory(callctx, koios.PoolID(ctx.String("pool-id")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-blocks",
			Category: "POOL",
			Usage:    "Return information about blocks minted by a given pool in current epoch (or _epoch_no if provided).",
			Flags: []cli.Flag{
				epochFlag,
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var epoch *koios.EpochNo
				if ctx.Uint("epoch") > 0 {
					v := koios.EpochNo(ctx.Uint64("epoch"))
					epoch = &v
				}

				res, err := api.GetPoolBlocks(callctx, koios.PoolID(ctx.String("pool-id")), epoch, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-updates",
			Category: "POOL",
			Usage:    "Return all pool updates for all pools or only updates for specific pool if specified.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				pool := koios.PoolID(ctx.String("pool-id"))

				res, err := api.GetPoolUpdates(callctx, &pool, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-relays",
			Category: "POOL",
			Usage:    "A list of registered relays for all currently registered/retiring (not retired) pools.",
			Action: func(ctx *cli.Context) error {
				res, err := api.GetPoolRelays(callctx, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
		{
			Name:     "pool-metadata",
			Category: "POOL",
			Usage:    "Metadata(on & off-chain) for all currently registered/retiring (not retired) pools.",
			Flags: []cli.Flag{
				&cli.StringSliceFlag{
					Name:     "pool-id",
					Aliases:  []string{"p"},
					Usage:    "Pool ids bech32 format, can be used multiple times for list of transactions",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var pids []koios.PoolID
				for _, pid := range ctx.StringSlice("pool-id") {
					pids = append(pids, koios.PoolID(pid))
				}

				res, err := api.GetPoolMetadata(callctx, pids, opts)
				apiOutput(ctx, res, err)
				return nil
			},
		},
	}...)
}
