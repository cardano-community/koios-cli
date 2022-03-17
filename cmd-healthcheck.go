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
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/cardano-community/koios-go-client"
)

type HealthcheckResponse struct {
	Tip struct {
		koios.Response
		Data            *koios.Tip    `json:"data,omitempty"`
		LastBlockAgeNs  time.Duration `json:"last_block_age_ns"`
		LastBlockAgeStr string        `json:"last_block_age_str"`
		TipTimoutNs     time.Duration `json:"tip_timeout_ns"`
		TipTimoutStr    string        `json:"tip_timeout_str"`
	} `json:"chk_tip"`

	RPC struct {
	} `json:"chk_rpcs"`
	Cache struct {
	} `json:"chk_cache_status"`
	Limit struct {
	} `json:"chk_limit"`
	Endpoints []struct {
		Endpoint string `json:"endpoint"`
	} `json:"chk_endpt_get"`
}

func attachHealthcheckCommmand(app *cli.App) {
	hcmd := &cli.Command{
		Name:     "healthcheck",
		Category: "GENERAL",
		Usage:    "perform healthcheck against Koios instances",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "timeout",
				Usage:   "Seconds when request will timeout.",
				Aliases: []string{"t"},
				Value:   2,
			},
			&cli.IntFlag{
				Name:    "age",
				Usage:   "Maximum allowed age of last block in seconds.",
				Aliases: []string{"a"},
				Value:   600,
			},
			&cli.BoolFlag{
				Name:  "quiet",
				Usage: "do not print nothing to sdtout",
				Value: false,
			},

			&cli.UintFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Set port",
				Value:   uint(koios.DefaultPort),
			},
			&cli.StringFlag{
				Name:  "host",
				Usage: "Set host",
				Value: koios.MainnetHost,
			},
			&cli.StringFlag{
				Name:  "api-version",
				Usage: "Set API version",
				Value: koios.DefaultAPIVersion,
			},
			&cli.StringFlag{
				Name:  "scheme",
				Usage: "Set URL scheme",
				Value: koios.DefaultScheme,
			},
			&cli.BoolFlag{
				Name:  "no-format",
				Usage: "prints response json strings directly without calling json pretty.",
				Value: false,
			},
		},
		Before: func(c *cli.Context) error {
			var (
				hostopt koios.Option
				err     error
			)
			if c.Bool("testnet") {
				hostopt = koios.Host(koios.TestnetHost)
			} else {
				hostopt = koios.Host(c.String("host"))
			}

			api, err = koios.New(
				hostopt,
				koios.APIVersion(c.String("api-version")),
				koios.Port(uint16(c.Uint("port"))),
				koios.Scheme(c.String("scheme")),
				koios.CollectRequestsStats(true),
			)

			return err
		},
		Action: func(ctx *cli.Context) error {
			tiptimeout := time.Second * time.Duration(ctx.Int("timeout"))
			tipage := time.Second * time.Duration(ctx.Int("age"))
			tipctx, tipcancel := context.WithTimeout(callctx, tiptimeout)
			defer tipcancel()

			res := HealthcheckResponse{}
			var failed bool

			// CHECK TIP
			// replace ./scripts/grest-helper-scripts/grest-poll.sh#L80
			// function chk_tip()
			res.Tip.TipTimoutNs = tiptimeout
			res.Tip.TipTimoutStr = tiptimeout.String()
			tipres, err := api.GetTip(tipctx)

			//nolint: nestif
			if err != nil {
				if err.Error() == "context deadline exceeded" {
					tipres.StatusCode = http.StatusRequestTimeout
					tipres.Status = http.StatusText(http.StatusRequestTimeout)
					tipres.Error.Code = fmt.Sprint(http.StatusRequestTimeout)
					tipres.Error.Message = fmt.Sprintf("instance did not respond within %s", tiptimeout)
					tipres.Error.Details = "to increase the timeout value adjust the --timeout flag value (n) seconds"
				}
				failed = true
			} else {
				reqtime, err := time.Parse(time.RFC1123, tipres.Date)
				if err != nil {
					return err
				}
				blocktime, err := time.Parse("2006-01-02T15:04:05", tipres.Data.BlockTime)
				if err != nil {
					return err
				}
				res.Tip.LastBlockAgeNs = reqtime.Sub(blocktime)
				res.Tip.LastBlockAgeStr = res.Tip.LastBlockAgeNs.String()
				res.Tip.Response = tipres.Response
				res.Tip.Data = tipres.Data

				if res.Tip.LastBlockAgeNs > tipage {
					failed = true
				}
			}

			// Output if quiet flag is not present
			if !ctx.Bool("quiet") {
				apiOutput(ctx, res, nil)
			}

			// exit with code 1 if any checks failed
			if failed {
				os.Exit(1)
			}

			return nil
		},
	}

	app.Commands = append(app.Commands, hcmd)
}
