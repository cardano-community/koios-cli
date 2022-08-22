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
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	cli "github.com/urfave/cli/v2"

	koios "github.com/cardano-community/koios-go-client"
)

const errstr = "error"

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
	Cache     []HealthcheckResult `json:"chk_cache_status"`
	Limit     *HealthcheckResult  `json:"chk_limit"`
	Endpoints []HealthcheckResult `json:"chk_endpt_get"`
}

type healthcheckTask struct {
	Name string
	Do   func(ctx *cli.Context, res *HealthcheckResponse) (bool, string)
}

type HealthcheckResult struct {
	Task    string `json:"task"`
	Status  string `json:"status"`
	Message string `json:"message"`
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
				Usage: "do not print nothing to stdout",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "print json result to stdout",
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
			opts = api.NewRequestOptions()
			return err
		},
		Action: func(ctx *cli.Context) error {
			res := &HealthcheckResponse{}
			tasks := []healthcheckTask{
				healthcheckCheckTip(),
				healthcheckCheckCacheStatusESDL(),
				healthcheckCheckCacheStatusEPHCLU(),
				healthcheckCheckCacheStatusLASVE(),
				healthcheckCheckLimit(),
				healthcheckEndpoints("/tx_metalabels"),
				healthcheckCheckRPC(),
			}

			for _, task := range tasks {
				performHealthcheckTask(ctx, task, res)
			}

			performHealthcheckTask(
				ctx,
				healthcheckEndpoints(fmt.Sprintf("/epoch_info?_epoch_no=%d", res.Tip.Data.EpochNo-1)),
				res,
			)

			// Output if quiet flag is not present
			if !ctx.Bool("quiet") && ctx.Bool("json") {
				apiOutput(ctx, res, nil)
			}
			return nil
		},
	}

	app.Commands = append(app.Commands, hcmd)
}

func performHealthcheckTask(ctx *cli.Context, task healthcheckTask, res *HealthcheckResponse) {
	fail, msg := task.Do(ctx, res)
	if fail {
		if !ctx.Bool("quiet") {
			if ctx.Bool("json") {
				apiOutput(ctx, res, nil)
			} else {
				log.Println("[ERROR ]: ", task.Name, " - ", msg)
			}
		}
		os.Exit(1)
	}

	if !ctx.Bool("quiet") && !ctx.Bool("json") {
		log.Println("[  OK  ]: ", task.Name, " - ", msg)
	}
}

// CHECK TIP
// replace ./scripts/grest-helper-scripts/grest-poll.sh#L80
// function chk_tip()s.
func healthcheckCheckTip() healthcheckTask {
	return healthcheckTask{
		Name: "check-tip",
		Do: func(ctx *cli.Context, res *HealthcheckResponse) (fail bool, msg string) {
			tiptimeout := time.Second * time.Duration(ctx.Int("timeout"))
			tipage := time.Second * time.Duration(ctx.Int("age"))
			tipctx, tipcancel := context.WithTimeout(callctx, tiptimeout)
			defer tipcancel()

			res.Tip.TipTimoutNs = tiptimeout
			res.Tip.TipTimoutStr = tiptimeout.String()
			tipres, err := api.GetTip(tipctx, nil)

			if err != nil {
				if err.Error() == "context deadline exceeded" {
					tipres.StatusCode = http.StatusRequestTimeout
					tipres.Status = http.StatusText(http.StatusRequestTimeout)
					tipres.Error.Code = fmt.Sprint(http.StatusRequestTimeout)
					tipres.Error.Message = fmt.Sprintf("instance did not respond within %s", tiptimeout)
					tipres.Error.Details = "to increase the timeout value adjust the --timeout flag value (n) seconds"
				}
				return true, tipres.Error.Message
			}

			reqtime, err := time.Parse(time.RFC1123, tipres.Date)
			if err != nil {
				return true, err.Error()
			}

			res.Tip.LastBlockAgeNs = reqtime.Sub(tipres.Data.BlockTime.Time)
			res.Tip.LastBlockAgeStr = res.Tip.LastBlockAgeNs.String()
			res.Tip.Response = tipres.Response
			res.Tip.Data = tipres.Data

			if res.Tip.LastBlockAgeNs > tipage {
				return false, fmt.Sprint("tip has expired - age ", res.Tip.LastBlockAgeStr)
			}
			return false, fmt.Sprint("tip age ", res.Tip.LastBlockAgeStr)
		},
	}
}

func healthcheckCheckCacheStatusESDL() healthcheckTask {
	return healthcheckTask{
		Name: "check-cache-status(eq.stake_distribution_lbh)",
		Do: func(ctx *cli.Context, res *HealthcheckResponse) (fail bool, msg string) {
			status := &HealthcheckResult{}
			status.Task = "eq.stake_distribution_lbh"
			status.Status = errstr
			defer func() {
				res.Cache = append(res.Cache, *status)
			}()

			opts.QuerySet("key", "eq.stake_distribution_lbh")
			rsp, err := api.GET(callctx, "/control_table", opts)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			defer rsp.Body.Close()

			body, err := io.ReadAll(rsp.Body)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}

			var pl = []struct {
				Value string `json:"last_value"`
			}{}
			err = json.Unmarshal(body, &pl)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			if len(pl) != 1 {
				return true, fmt.Sprintf("invalid response length %d", len(pl))
			}

			blockNo, err := strconv.Atoi(pl[0].Value)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			diff := res.Tip.Data.BlockNo - blockNo
			if diff > 1000 {
				msg := fmt.Sprintf("Stake Distribution cache too far from tip  (%d) blocks", diff)
				status.Message = msg
				return true, msg
			}
			ok := fmt.Sprintf("block diff %d", diff)
			status.Status = "ok"
			status.Message = ok
			return false, msg
		},
	}
}

func healthcheckCheckCacheStatusEPHCLU() healthcheckTask {
	return healthcheckTask{
		Name: "check-cache-status(eq.pool_history_cache_last_updated)",
		Do: func(ctx *cli.Context, res *HealthcheckResponse) (fail bool, msg string) {
			status := &HealthcheckResult{}
			status.Task = "eq.pool_history_cache_last_updated"
			status.Status = errstr
			defer func() {
				res.Cache = append(res.Cache, *status)
			}()

			opts.QuerySet("key", "eq.pool_history_cache_last_updated")
			rsp, err := api.GET(callctx, "/control_table", opts)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			defer rsp.Body.Close()

			body, err := io.ReadAll(rsp.Body)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}

			var pl = []struct {
				Value string `json:"last_value"`
			}{}
			err = json.Unmarshal(body, &pl)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			if len(pl) != 1 {
				msg := fmt.Sprintf("invalid response length %d", len(pl))
				status.Message = msg
				return true, msg
			}

			ts, err := time.Parse("2006-01-02 15:04:05.99999", pl[0].Value)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}

			diff := res.Tip.Data.BlockTime.Second() - ts.Second()

			if diff > 1000 {
				msg := fmt.Sprintf("Pool History cache too far from tip (%s)", (time.Duration(diff) * time.Second).String())
				status.Message = msg
				return true, msg
			}
			ok := fmt.Sprintf(
				"time diff %s",
				(time.Duration(diff) * time.Second).String(),
			)
			status.Message = ok
			status.Status = "ok"
			return false, ok
		},
	}
}

func healthcheckCheckCacheStatusLASVE() healthcheckTask {
	return healthcheckTask{
		Name: "check-cache-status(eq.last_active_stake_validated_epoch)",
		Do: func(ctx *cli.Context, res *HealthcheckResponse) (fail bool, msg string) {
			status := &HealthcheckResult{}
			status.Task = "eq.pool_history_cache_last_updated"
			status.Status = errstr
			defer func() {
				res.Cache = append(res.Cache, *status)
			}()

			opts.QuerySet("key", "eq.pool_history_cache_last_updated")
			rsp, err := api.GET(callctx, "/control_table", opts)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			defer rsp.Body.Close()

			body, err := io.ReadAll(rsp.Body)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}

			var pl = []struct {
				Value string `json:"last_value"`
			}{}
			err = json.Unmarshal(body, &pl)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			if len(pl) != 1 {
				status.Message = err.Error()
				return true, "Active Stake cache not populated"
			}

			ts, err := time.Parse("2006-01-02 15:04:05.99999", pl[0].Value)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}

			diff := res.Tip.Data.BlockTime.Time.Second() - ts.Second()

			if diff > 1000 {
				msg := fmt.Sprintf("Active Stake cache too far from tip %s", (time.Duration(diff) * time.Second).String())
				status.Message = msg
				return true, msg
			}
			ok := fmt.Sprintf("time diff %s", (time.Duration(diff) * time.Second).String())
			status.Status = "ok"
			status.Message = ok
			return false, ok
		},
	}
}

func healthcheckCheckLimit() healthcheckTask {
	return healthcheckTask{
		Name: "check-limit",
		Do: func(ctx *cli.Context, res *HealthcheckResponse) (fail bool, msg string) {
			res.Limit = &HealthcheckResult{}
			res.Limit.Task = "check-limit"
			res.Limit.Status = errstr

			rsp, err := api.GET(callctx, "/blocks", nil)
			if err != nil {
				res.Limit.Message = err.Error()
				return true, err.Error()
			}
			defer rsp.Body.Close()

			if rsp.Header.Get("content-range") != "0-999/*" {
				msg := fmt.Sprintf(
					"The PostgREST config for uses a custom limit that does not match monitoring instances expected 999 got %s",
					rsp.Header.Get("content-range"),
				)
				res.Limit.Status = errstr
				res.Limit.Message = msg
				return true, msg
			}
			ok := "PostgREST config limit is 999"
			res.Limit.Status = "ok"
			res.Limit.Message = ok
			return false, ok
		},
	}
}

func healthcheckEndpoints(endpoint string) healthcheckTask {
	e := fmt.Sprintf("check-endpoint(%s)", endpoint)
	return healthcheckTask{
		Name: e,
		Do: func(ctx *cli.Context, res *HealthcheckResponse) (fail bool, msg string) {
			status := &HealthcheckResult{}
			status.Task = e
			status.Status = errstr
			defer func() {
				res.Endpoints = append(res.Endpoints, *status)
			}()
			u, err := url.ParseRequestURI(endpoint)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}

			opts.Page(1)
			opts.PageSize(1)
			opts.QueryApply(u.Query())

			rsp, err := api.GET(callctx, u.Path, opts)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			defer rsp.Body.Close()
			body, err := io.ReadAll(rsp.Body)
			if err != nil {
				status.Message = err.Error()
				return true, err.Error()
			}
			var b = []interface{}{}
			err = json.Unmarshal(body, &b)
			if err != nil {
				status.Message = err.Error()
				return true, status.Message
			}

			if len(b) != 1 {
				status.Message = fmt.Sprintf("worng result (%d)", len(b))
				return true, status.Message
			}
			status.Status = "ok"

			return false, "got valid response"
		},
	}
}

func healthcheckCheckRPC() healthcheckTask {
	return healthcheckTask{
		Name: "check-rpcs",
		Do: func(ctx *cli.Context, res *HealthcheckResponse) (fail bool, msg string) {
			return false, "not implemented"
		},
	}
}
