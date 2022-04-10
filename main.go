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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"

	"github.com/urfave/cli/v2"

	"github.com/cardano-community/koios-go-client"
)

//nolint: gochecknoglobals
var (
	ErrCommand = errors.New("command error")

	callctx context.Context
	cancel  context.CancelFunc

	// Populated by goreleaser during build.
	version = "dev"
	date    = ""

	//nolint: gochecknoglobals
	api  *koios.Client
	opts *koios.RequestOptions
)

//nolint:gochecknoinits
func init() {
	if info, available := debug.ReadBuildInfo(); available && version == "dev" && date == "" {
		version = info.Main.Version
	}
}

// application main.
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	callctx, cancel = context.WithCancel(context.Background())
	go func() {
		<-c
		cancel()
	}()

	app := &cli.App{
		Version: version,
		Authors: []*cli.Author{
			{
				Name: "The Cardano Community Authors",
			},
		},
		Copyright:            "(c) 2022",
		Usage:                "Koios CLI Client.",
		EnableBashCompletion: true,
	}

	attachAPICommmand(app)
	attachHealthcheckCommmand(app)
	handleErr(app.Run(os.Args))
}

// common error handler.
func handleErr(err error) {
	if err == nil {
		return
	}
	cancel()
	// trace := err
	// for errors.Unwrap(trace) != nil {
	// 	trace = errors.Unwrap(trace)
	// 	log.Println(trace)
	// }
	log.Fatal(err)
}

// helper to print json response body.
func printJSON(ctx *cli.Context, body []byte) {
	if ctx.Bool("no-format") {
		fmt.Println(string(body))
		return
	}
	var pretty bytes.Buffer
	handleErr(json.Indent(&pretty, body, "", "    "))
	fmt.Println(pretty.String())
}

// output koios api client responses.
func apiOutput(ctx *cli.Context, data interface{}, err error) {
	handleErr(err)

	if ctx.Bool("no-format") {
		out, err := json.Marshal(data)
		handleErr(err)
		fmt.Println(string(out))
		return
	}

	out, merr := json.MarshalIndent(data, "", " ")
	handleErr(merr)
	fmt.Println(string(out))
}
