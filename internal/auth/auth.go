// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package auth

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"log/slog"

	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/pkg/strings/textfmt"
	"github.com/happy-sdk/happy/pkg/vars"
	"github.com/happy-sdk/happy/sdk/cli"
)

func Command() *happy.Command {
	cmd := happy.NewCommand("auth",
		happy.Option("description", "Manage you subscription with Koios API"),
	)

	cmd.AddSubCommand(cmdAuthAddToken())
	return cmd
}

func cmdAuthAddToken() *happy.Command {
	cmd := happy.NewCommand("add-token",
		happy.Option("description", "Add a new JWT Bearer Auth token generated via https://koios.rest Profile page."),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios auth add-token <token>"),
	)

	cmd.AddInfo("Add a new API token to your subscription")
	cmd.AddInfo(`
    This will create new --profile with name of project ID.

    If same project ID already exists, it will ask before overwriting the existing profile.

    Example: koios-cli auth add-token eyJhbGciOiJ...-JGHedpgsQVsI
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		token, err := koios.GetTokenAuthInfo(args.Arg(0).String())
		if err != nil {
			return err
		}

		infotbl := textfmt.Table{
			Title: "Provided Token Information",
			// WithHeader: true,
		}
		// infotbl.AddRow("KEY", "VALUE")
		infotbl.AddRow("Tier", token.Tier.String())
		infotbl.AddRow("Project ID", token.ProjID)
		infotbl.AddRow("Address", token.Addr)
		infotbl.AddRow("Expires", token.Expires.String())
		infotbl.AddDivider()
		infotbl.AddRow("Max Requests/day", fmt.Sprint(token.MaxRequests))
		infotbl.AddRow("Max Request/s", fmt.Sprint(token.MaxRPS))
		infotbl.AddRow("Max Query Timeout", token.MaxQueryTimeout.String())
		infotbl.AddRow("CORS Restricted", fmt.Sprint(token.CORSRestricted))
		fmt.Println(infotbl.String())

		ok, profileDir := happy.HasProfile(sess, token.ProjID)
		koiosAuthFile := filepath.Join(profileDir, "koios-auth.jwt")

		if ok {
			if _, err := os.Stat(koiosAuthFile); err != nil {
				if os.IsNotExist(err) {
					goto writeToken
				} else {
					return err
				}
			}
			if !cli.AskForConfirmation("Profile already exists. Do you want to overwrite it? [y/N]: ") {
				return fmt.Errorf("profile %s already exists, skipping", token.ProjID)
			}
		}
	writeToken:

		if err := os.MkdirAll(profileDir, 0700); err != nil {
			return err
		}

		if err := os.WriteFile(koiosAuthFile, []byte(args.Arg(0).String()), 0600); err != nil {
			return err
		}
		profileFile := filepath.Join(profileDir, "profile.preferences")
		profile := sess.Settings().All()
		pd := vars.Map{}
		for _, setting := range profile {
			if setting.Persistent() {
				if err := pd.Store(setting.Key(), setting.Value().String()); err != nil {
					return err
				}
			}
		}
		pddata := pd.ToKeyValSlice()
		var dest bytes.Buffer
		enc := gob.NewEncoder(&dest)
		if err := enc.Encode(pddata); err != nil {
			return err
		}

		if err := os.WriteFile(profileFile, dest.Bytes(), 0600); err != nil {
			return err
		}
		sess.Log().Ok("Token added successfully", slog.String("profile", token.ProjID))

		return nil
	})

	return cmd
}
