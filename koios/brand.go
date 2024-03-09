// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2024 The Cardano Community Authors

package koios

import (
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/pkg/branding"
	"github.com/happy-sdk/happy/pkg/cli/ansicolor"
)

func Brand() happy.BrandFunc {

	ansi := ansicolor.New()
	ansi.Primary = ansicolor.RGB(67, 174, 124)
	ansi.Secondary = ansicolor.RGB(53, 145, 101)
	ansi.Accent = ansicolor.RGB(32, 141, 188)

	builder := branding.New(branding.Info{
		Name:        "koios",
		Slug:        "koios",
		Version:     "1.0",
		Description: "Koios is a distributed & open-source public API query layer for Cardano, that is elastic in nature and addresses ever-demanding requirements from Cardano Blockchain. It provides a easy to query RESTful layer that has a lot of flexibility to cater for different data consumption requirements from blockchain.",
	}).WithANSI(ansi)

	return func() (happy.Brand, error) {
		return builder.Build()
	}
}
