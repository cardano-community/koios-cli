// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
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

  _addresses is constructed from command line arguments,

  Example: koios-cli api address_info \
    addr1qy2jt0qpqz2z2z9zx5w4xemekkce7yderz53kjue53lpqv90lkfa9sgrfjuz6uvt4uqtrqhl2kj0a9lnr9ndzutx32gqleeckv \
    addr1q9xvgr4ehvu5k5tmaly7ugpnvekpqvnxj8xy50pa7kyetlnhel389pa4rnq6fmkzwsaynmw0mnldhlmchn2sfd589fgsz9dd0y

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
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

func cmdAddressAddressAssets(c *client) *happy.Command {
	cmd := happy.NewCommand("address_assets",
		happy.Option("description", "Address Assets"),
		happy.Option("category", categoryAddress),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api address_assets [addresses...] // max 50"),
	).WithFalgs(queryFlag)
	cmd.AddInfo("Get the list of all the assets (policy, name and quantity) for given addresses")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#post-/address_assets

  _addresses is constructed from command line arguments,

  Example: koios-cli api address_assets \
    addr1qy2jt0qpqz2z2z9zx5w4xemekkce7yderz53kjue53lpqv90lkfa9sgrfjuz6uvt4uqtrqhl2kj0a9lnr9ndzutx32gqleeckv \
    addr1q9xvgr4ehvu5k5tmaly7ugpnvekpqvnxj8xy50pa7kyetlnhel389pa4rnq6fmkzwsaynmw0mnldhlmchn2sfd589fgsz9dd0y

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		var addresses []koios.Address
		for _, addr := range args.Args() {
			addresses = append(addresses, koios.Address(addr.String()))
		}
		res, err := c.koios().GetAddressesAssets(sess, addresses, opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func cmdAddressAddressTxs(c *client) *happy.Command {
	cmd := happy.NewCommand("address_txs",
		happy.Option("description", "Address Transactions"),
		happy.Option("category", categoryAddress),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api address_txs --after-block-height 9945516 [addresses...] // max 50"),
	).
		WithFalgs(varflag.UintFunc("after-block-height", 0, "Only fetch information after specific block height"))

	cmd.AddInfo("Get the transaction hash list of input address array, optionally filtering after specified block height (inclusive)")
	cmd.AddInfo(`
  Docs:https://api.koios.rest/#post-/address_txs

  _addresses is constructed from command line arguments,

  Example: koios-cli api address_txs \
    --after-block-height 8000000 \
    addr1qy2jt0qpqz2z2z9zx5w4xemekkce7yderz53kjue53lpqv90lkfa9sgrfjuz6uvt4uqtrqhl2kj0a9lnr9ndzutx32gqleeckv \
    addr1q9xvgr4ehvu5k5tmaly7ugpnvekpqvnxj8xy50pa7kyetlnhel389pa4rnq6fmkzwsaynmw0mnldhlmchn2sfd589fgsz9dd0y

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		var addresses []koios.Address
		for _, addr := range args.Args() {
			addresses = append(addresses, koios.Address(addr.String()))
		}

		res, err := c.koios().GetAddressTxs(sess, addresses, args.Flag("after-block-height").Var().Uint64(), opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func cmdAddressAddressUtxos(c *client) *happy.Command {
	cmd := happy.NewCommand("address_utxos",
		happy.Option("description", "Address UTxOs"),
		happy.Option("category", categoryAddress),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api address_utxos [addresses...] // max 50"),
	).
		WithFalgs(varflag.BoolFunc("extended", false, "Controls whether or not certain optional fields supported by a given endpoint are populated as a part of the call", "e"))

	cmd.AddInfo("Get the UTxO set for given addresses")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#post-/address_utxos

  _addresses constructed from command line arguments,

  Example: koios-cli api address_utxos \
    --extended \
    addr1qy2jt0qpqz2z2z9zx5w4xemekkce7yderz53kjue53lpqv90lkfa9sgrfjuz6uvt4uqtrqhl2kj0a9lnr9ndzutx32gqleeckv \
    addr1q9xvgr4ehvu5k5tmaly7ugpnvekpqvnxj8xy50pa7kyetlnhel389pa4rnq6fmkzwsaynmw0mnldhlmchn2sfd589fgsz9dd0y

    `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		var addresses []koios.Address
		for _, addr := range args.Args() {
			addresses = append(addresses, koios.Address(addr.String()))
		}

		res, err := c.koios().GetAddressUTxOs(sess, addresses, args.Flag("extended").Var().Bool(), opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}

func cmdAddressCredentialTxs(c *client) *happy.Command {
	cmd := happy.NewCommand("credential_txs",
		happy.Option("description", "Transactions from payment credentials"),
		happy.Option("category", categoryAddress),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api credential_txs --after-block-height 6238675 [credentials...] // max 50"),
	).
		WithFalgs(varflag.UintFunc("after-block-height", 0, "Only fetch information after specific block height (inclusive)"))

	cmd.AddInfo("Get the transaction hash list of input payment credential array, optionally filtering after specified block height (inclusive)")

	cmd.AddInfo(`
  Docs: https://api.koios.rest/#post-/credential_txs

  _payment_credentials parameter is constructed from command line arguments,

  Example: koios-cli api credential_txs \
    --after-block-height 6238675 \
    025b0a8f85cb8a46e1dda3fae5d22f07e2d56abb4019a2129c5d6c52 \
    13f6870c5e4f3b242463e4dc1f2f56b02a032d3797d933816f15e555

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return nil
		}
		var credentials []koios.PaymentCredential
		for _, cred := range args.Args() {
			credentials = append(credentials, koios.PaymentCredential(cred.String()))
		}
		res, err := c.koios().GetCredentialTxs(sess, credentials, args.Flag("after-block-height").Var().Uint64(), opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})
	return cmd
}

func cmdAddressCredentialUtxos(c *client) *happy.Command {
	cmd := happy.NewCommand("credential_utxos",
		happy.Option("description", "UTxOs from payment credentials"),
		happy.Option("category", categoryAddress),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
		happy.Option("usage", "koios api credential_utxos [credentials...] // max 50"),
	).
		WithFalgs(varflag.BoolFunc("extended", false, "Controls whether or not certain optional fields supported by a given endpoint are populated as a part of the call", "e"))

	cmd.AddInfo("Get UTxO details for requested payment credentials")

	cmd.AddInfo(`
  Docs: https://api.koios.rest/#post-/credential_utxos

  _payment_credentials parameter is constructed from command line arguments,

  Example: koios-cli api credential_utxos \
    --extended \
    025b0a8f85cb8a46e1dda3fae5d22f07e2d56abb4019a2129c5d6c52 \
    13f6870c5e4f3b242463e4dc1f2f56b02a032d3797d933816f15e555

  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return nil
		}
		var credentials []koios.PaymentCredential
		for _, cred := range args.Args() {
			credentials = append(credentials, koios.PaymentCredential(cred.String()))
		}
		res, err := c.koios().GetCredentialUTxOs(sess, credentials, args.Flag("extended").Var().Bool(), opts)
		apiOutput(c.noFormat, res, err)
		return nil
	})

	return cmd
}
