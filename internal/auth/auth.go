// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package auth

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/pkg/strings/textfmt"
	"github.com/happy-sdk/happy/sdk/cli"
)

func Command() *happy.Command {
	cmd := happy.NewCommand("auth",
		happy.Option("description", "Manage you subscription with Koios API"),
	)

	cmd.AddInfo("Manage your subscription with Koios API")

	cmd.AddInfo(`
    Koios API provides subscription to access the API endpoints.
    You can create a subscription by visiting https://koios.rest
    and creating a profile. Once you have a profile, you can add
    the subscription token to Koios CLI using this command.

    Example: koios-cli auth add <jwt-token>

    This will will create local profile with Porject ID read from the JWT Token.
    You can then use the profile to access Koios API endpoints.

    Example: koios-cli --profile <project-id> api tip

    To see your current usage add --stats flag to the api command. e.g.

    Example: koios-cli --profile <project-id> api --stats tip

    If you want to use Koios APi with your token without saving it to disk,
    you can use --auth flag with the token.

    Example: koios-cli api --auth <jwt-token> tip
  `)

	cmd.AddSubCommand(cmdAuthAdd())
	cmd.AddSubCommand(cmdAuthRemove())
	cmd.AddSubCommand(cmdList())
	return cmd
}

func cmdAuthAdd() *happy.Command {
	cmd := happy.NewCommand("add",
		happy.Option("description", "Create profile from JWT Bearer Auth token generated via https://koios.rest Profile page."),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios auth add-token <token>"),
	)

	cmd.AddInfo("Add a new API token to your subscription")
	cmd.AddInfo(`
    This will create new --profile with name of project ID.

    If same project ID already exists, it will ask before overwriting the existing profile.

    Example: koios-cli auth add eyJhbGciOiJ...-JGHedpgsQVsI
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		token := strings.TrimSpace(args.Arg(0).String())
		authInfo, err := koios.GetTokenAuthInfo(token)
		if err != nil {
			return err
		}

		infotbl := textfmt.Table{
			Title: "Provided Token Information",
			// WithHeader: true,
		}
		// infotbl.AddRow("KEY", "VALUE")
		infotbl.AddRow("Tier", authInfo.Tier.String())
		infotbl.AddRow("Project ID", authInfo.ProjID)
		infotbl.AddRow("Address", authInfo.Addr)
		infotbl.AddRow("Expires", authInfo.Expires.String())
		infotbl.AddDivider()
		infotbl.AddRow("Max Requests/day", fmt.Sprint(authInfo.MaxRequests))
		infotbl.AddRow("Max Request/s", fmt.Sprint(authInfo.MaxRPS))
		infotbl.AddRow("Max Query Timeout", authInfo.MaxQueryTimeout.String())
		infotbl.AddRow("CORS Restricted", fmt.Sprint(authInfo.CORSRestricted))
		fmt.Println(infotbl.String())

		var configRootDir string
		if sess.Get("app.profile.name").String() == "public" {
			configRootDir = filepath.Join(sess.Get("app.fs.path.config").String(), "profiles")
		} else {
			configRootDir = filepath.Dir(sess.Get("app.fs.path.config").String())
		}

		profileName := authInfo.ProjID
		if sess.Get("app.devel").Bool() {
			profileName += "-devel"
		}

		configDir := filepath.Join(configRootDir, profileName)
		koiosAuthFile := filepath.Join(configDir, "koios.subscription")

		if _, err := os.Stat(koiosAuthFile); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("failed to check auth token file: %w", err)
			}
		} else {
			if !cli.AskForConfirmation(fmt.Sprintf("Profile (%s) already exists. Do you want to overwrite it?", authInfo.ProjID)) {
				return fmt.Errorf("profile %s already exists, skipping", authInfo.ProjID)
			}
		}
		if os.MkdirAll(configDir, 0700); err != nil {
			return fmt.Errorf("failed to create profile directory: %w", err)
		}

		sess.Log().Notice("token will be saved", slog.String("path", koiosAuthFile))
		if !cli.AskForConfirmation(`
      The token will be stored in unencrypted form, which is a security risk similar
      to keeping tokens in .env files. If you opt to save, it'll be located at:
      ` + koiosAuthFile + `

      You can then call Koios API with the token using --profile flag.
      e.g.
        koios-cli --profile "` + authInfo.ProjID + `" api tip

      Do you wish to save the token?`) {
			return errors.New("cancelled by user")
		}

		subscription := &Subscription{
			path:          koiosAuthFile,
			JWT:           token,
			RequestsToday: 0,
		}

		if err := subscription.Save(); err != nil {
			return fmt.Errorf("failed to save subscription: %w", err)
		}

		sess.Log().Ok("token saved", slog.String("path", koiosAuthFile))
		return nil
	})

	return cmd
}

func cmdAuthRemove() *happy.Command {
	cmd := happy.NewCommand("remove",
		happy.Option("description", "Remove profile and subscription token from local system"),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios auth remove <profile-name>"),
	)

	cmd.AddInfo("Remove a profile from your subscription")
	cmd.AddInfo(`
    This will remove the subscription profile from your local environment.

    Example: koios-cli auth remove my-project
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {

		profileName := strings.TrimSpace(args.Arg(0).String())
		if len(profileName) == 0 {
			return errors.New("profile name is required")
		}
		if sess.Get("app.devel").Bool() {
			profileName += "-devel"
		}

		configRootDir := filepath.Dir(sess.Get("app.fs.path.config").String())
		configDir := filepath.Join(configRootDir, profileName)
		koiosAuthFile := filepath.Join(configDir, "koios.subscription")

		if _, err := os.Stat(koiosAuthFile); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("profile %s does not exist", profileName)
			}
			return fmt.Errorf("failed to check auth token file: %w", err)
		}

		if !cli.AskForConfirmation("Are you sure you want to remove profile " + profileName + "?") {
			return errors.New("cancelled by user")
		}

		if err := os.Remove(koiosAuthFile); err != nil {
			return fmt.Errorf("failed to remove profile: %w", err)
		}

		sess.Log().Ok("profile removed", slog.String("profile", profileName))
		return nil
	})

	return cmd
}

func cmdList() *happy.Command {
	cmd := happy.NewCommand("list",
		happy.Option("description", "List all saved profiles and their stats"),
	)

	cmd.AddInfo("List all profiles and their subscription tokens")
	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		configRootDir := filepath.Dir(sess.Get("app.fs.path.config").String())
		files, err := os.ReadDir(configRootDir)
		if err != nil {
			return fmt.Errorf("failed to read config directory: %w", err)
		}

		if len(files) == 0 {
			fmt.Println("No profiles found")
			return nil
		}

		profiles := textfmt.Table{
			Title:      "Profiles",
			WithHeader: true,
		}
		profiles.AddRow("Profile", "Tier", "Expires", "Requests Today")

		for _, file := range files {
			if !file.IsDir() {
				continue
			}
			koiosAuthFile := filepath.Join(configRootDir, file.Name(), "koios.subscription")
			if _, err := os.Stat(koiosAuthFile); err != nil {
				continue
			}

			sub, err := LoadSubscriptionFile(koiosAuthFile)
			if err != nil {
				return err
			}

			profileName := file.Name()
			listitem := subscriptionListItem{
				Name: profileName,
			}
			listitem.Devel = strings.HasSuffix(profileName, "-devel")
			listitem.RequestsToday = sub.RequestsToday

			authInfo, err := koios.GetTokenAuthInfo(sub.JWT)
			if err != nil {
				return err
			}
			listitem.Expires = authInfo.Expires.String()
			listitem.Tier = authInfo.Tier.String()

			profiles.AddRow(profileName, listitem.Tier, listitem.Expires, fmt.Sprint(sub.RequestsToday))
		}
		fmt.Println(profiles.String())
		return nil
	})

	return cmd
}

type subscriptionListItem struct {
	Name          string
	Tier          string
	Expires       string
	RequestsToday uint
	Devel         bool
}

type Subscription struct {
	path          string
	JWT           string
	RequestsToday uint
	Date          string
	History       []RequestHistoryEntry
}

func (s *Subscription) Save() error {
	if s.path == "" {
		return errors.New("subscription path is not set")
	}
	s.update()

	today := time.Now().Format("2006-01-02")
	if s.Date != today {
		s.RequestsToday = 0
		s.Date = today
	}

	var dest bytes.Buffer
	enc := gob.NewEncoder(&dest)
	if err := enc.Encode(s); err != nil {
		return err
	}
	return os.WriteFile(s.path, dest.Bytes(), 0600)
}

func LoadSubscription(sess *happy.Session) (*Subscription, error) {
	configDir := filepath.Join(sess.Get("app.fs.path.config").String())
	koiosAuthFile := filepath.Join(configDir, "koios.subscription")
	return LoadSubscriptionFile(koiosAuthFile)
}

func LoadSubscriptionFile(koiosAuthFile string) (*Subscription, error) {
	if _, err := os.Stat(koiosAuthFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("auth token file not found %s", koiosAuthFile)
		}
		return nil, fmt.Errorf("failed to check auth token file: %w", err)
	}

	data, err := os.ReadFile(koiosAuthFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read auth token file: %w", err)
	}

	sub := &Subscription{}
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(sub); err != nil {
		return nil, fmt.Errorf("failed to decode auth token file: %w", err)
	}
	sub.path = koiosAuthFile
	sub.update()
	return sub, nil
}

func (s *Subscription) update() {
	today := time.Now().Format("2006-01-02")
	if s.Date != today {
		if s.Date != "" {
			s.History = append(s.History, RequestHistoryEntry{
				Date:     s.Date,
				Requests: s.RequestsToday,
			})
			sort.Slice(s.History, func(i, j int) bool {
				return s.History[i].Date < s.History[j].Date
			})
		}
		s.RequestsToday = 0
		s.Date = today
	}
}

type RequestHistoryEntry struct {
	Date     string
	Requests uint
}
