// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
	"fmt"
	"strings"

	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
	"github.com/happy-sdk/happy/pkg/vars/varflag"
)

const categoryAddress = "address"

// Address:
// https://api.koios.rest/#tag--Address
func address(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryAddress, "Query information about specific address(es)")
	cmd.AddSubCommand(cmdAddressAddressInfo(c))
	cmd.AddSubCommand(cmdAddressAddressUtxos(c))
	cmd.AddSubCommand(cmdAddressCredentialUtxos(c))
	cmd.AddSubCommand(cmdAddressAddressTxs(c))
	cmd.AddSubCommand(cmdAddressCredentialTxs(c))
	cmd.AddSubCommand(cmdAddressAddressAssets(c))
}

func cmdAddressAddressInfo(c *client) *happy.Command {
	cmd := happy.NewCommand("address_info",
		happy.Option("description", "Address Information"),
		happy.Option("category", categoryAddress),
		happy.Option("argn.max", 100),
		happy.Option("usage", "koios api address_info [_addresses...] // max 100"),
	).WithFalgs(queryFlag)
	cmd.AddInfo("Get address info - balance, associated stake address (if any) and UTxO set for given addresses")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#post-/address_info

  _addresses query parameter is constructed from command line arguments,

  Example: koios-cli api address_info \
    addr1qy2jt0qpqz2z2z9zx5w4xemekkce7yderz53kjue53lpqv90lkfa9sgrfjuz6uvt4uqtrqhl2kj0a9lnr9ndzutx32gqleeckv \
    addr1q9xvgr4ehvu5k5tmaly7ugpnvekpqvnxj8xy50pa7kyetlnhel389pa4rnq6fmkzwsaynmw0mnldhlmchn2sfd589fgsz9dd0y

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		if args.Argn() == 0 {
			return fmt.Errorf("atleast one address required")
		}

		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		var addresses []koios.Address
		for _, addr := range args.Args() {
			addresses = append(addresses, koios.Address(addr.String()))
		}
		res, err := c.koios().GetAddressesInfo(sess, addresses, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func cmdAddressAddressUtxos(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "address_utxos")
}

func cmdAddressCredentialUtxos(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "credential_utxos")
}

func cmdAddressAddressTxs(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "address_txs")
}

func cmdAddressCredentialTxs(c *client) *happy.Command {
	return notimplCmd(categoryAddress, "credential_txs")
}

func cmdAddressAddressAssets(c *client) *happy.Command {
	cmd := happy.NewCommand("address_assets",
		happy.Option("description", "Address Assets"),
		happy.Option("category", categoryAddress),
	).WithFalgs(queryFlag)
	cmd.AddInfo("Get the list of all the assets (policy, name and quantity) for given addresses")
	cmd.AddInfo(`
Docs: https://api.koios.rest/#post-/address_assets`)

	cmd.AddFlag(varflag.StringFunc("addresses", "", "Comma separated list of addresses to query", "a"))
	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		if !args.Flag("addresses").Present() {
			return fmt.Errorf("missing required flag: addresses")
		}

		addrs := strings.Split(args.Flag("addresses").String(), ",")

		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		var addresses []koios.Address
		for _, addr := range addrs {
			addresses = append(addresses, koios.Address(addr))
		}
		res, err := c.koios().GetAddressesAssets(sess, addresses, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}
