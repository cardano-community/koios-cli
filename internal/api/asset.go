// SPDX-License-Identifier: Apache-2.0
//
// Copyright © 2022 The Cardano Community Authors

package api

import (
	"strings"

	"github.com/cardano-community/koios-go-client/v4"
	"github.com/happy-sdk/happy"
)

const categoryAsset = "asset"

// Asset:
// https://api.koios.rest/#tag--Asset
func asset(cmd *happy.Command, c *client) {
	cmd.DescribeCategory(categoryAsset, "Query Asset related informations")

	cmd.AddSubCommand(cmdAssetAddresses(c))
	cmd.AddSubCommand(cmdAssetHistory(c))
	cmd.AddSubCommand(cmdAssetInfo(c))
	cmd.AddSubCommand(cmdAssetList(c))
	cmd.AddSubCommand(cmdAssetNftAddress(c))
	cmd.AddSubCommand(cmdAssetSummary(c))
	cmd.AddSubCommand(cmdAssetTokenRegistry(c))
	cmd.AddSubCommand(cmdAssetTxs(c))
	cmd.AddSubCommand(cmdAssetUtxos(c))
	cmd.AddSubCommand(cmdAssetPolicyAssetAddresses(c))
	cmd.AddSubCommand(cmdAssetPolicyAssetInfo(c))
	cmd.AddSubCommand(cmdAssetPolicyAssetList(c))
}

func cmdAssetAddresses(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_addresses")
}

func cmdAssetHistory(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_history")
}

func cmdAssetInfo(c *client) *happy.Command {
	cmd := happy.NewCommand("asset_info",
		happy.Option("description", "Asset Information (Bulk)"),
		happy.Option("category", categoryAsset),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 50),
	).WithFalgs(pagingFlags...)

	cmd.AddInfo("Get the information of an asset including first minting & token registry metadata.")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/asset_info

  Example: koios-cli api asset_info \
    750900e4999ebe0d58f19b634768ba25e525aaf12403bfe8fe130501.424f4f4b \
    f0ff48bbb7bbe9d59a40f1ce90e9e9d0ff5002ec48f232b49ca0fb9a.6b6f696f732e72657374

  Example: koios-cli api asset_info
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		var assets []koios.Asset
		for _, arg := range args.Args() {
			policy, asset, _ := strings.Cut(arg.String(), ".")
			assets = append(assets, koios.Asset{
				PolicyID:  koios.PolicyID(policy),
				AssetName: koios.AssetName(asset),
			})
		}

		res, err := c.koios().GetAssetInfo(sess, assets, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdAssetList(c *client) *happy.Command {
	cmd := happy.NewCommand("asset_list",
		happy.Option("description", "Asset List"),
		happy.Option("category", categoryAsset),
	).WithFalgs(pagingFlags...)
	cmd.AddInfo("Get the list of all native assets (paginated)")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/asset_list

  Example: koios-cli api asset_list
    {
      ...
      "data": [
        {
          "asset": {
            "asset_name": "6e7574636f696e",
            "policy_id": "00000002df633853f6a47465c9496721d2d5b1291b8398016c0e87ae",
            "fingerprint": "asset12h3p5l3nd5y26lr22am7y7ga3vxghkhf57zkhd"
          }
        }
      ]
    }
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		res, err := c.koios().GetAssets(sess, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})
	return cmd
}

func cmdAssetNftAddress(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_nft_address")
}

func cmdAssetSummary(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_summary")
}

func cmdAssetTokenRegistry(c *client) *happy.Command {
	cmd := happy.NewCommand("asset_token_registry",
		happy.Option("description", "Asset Token Registry"),
		happy.Option("category", categoryAsset),
	).WithFalgs(pagingFlags...)
	cmd.AddInfo("Get a list of assets registered via token registry on github")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/asset_token_registry

  Example: koios-cli api asset_token_registry
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return err
		}
		res, err := c.koios().GetAssetTokenRegistry(sess, opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}

func cmdAssetTxs(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_txs")
}

func cmdAssetUtxos(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "asset_utxos")
}

func cmdAssetPolicyAssetAddresses(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "policy_asset_addresses")
}

func cmdAssetPolicyAssetInfo(c *client) *happy.Command {
	return notimplCmd(categoryAsset, "policy_asset_info")
}

func cmdAssetPolicyAssetList(c *client) *happy.Command {
	cmd := happy.NewCommand("policy_asset_list",
		happy.Option("description", "Policy Asset List"),
		happy.Option("category", categoryAsset),
		happy.Option("argn.min", 1),
		happy.Option("argn.max", 1),
		happy.Option("usage", "koios api policy_asset_list [policy_id]"),
	).WithFalgs(pagingFlags...)
	cmd.AddInfo("Get the list of all assets minted under a given policy (paginated)")
	cmd.AddInfo(`
  Docs: https://api.koios.rest/#get-/policy_asset_list

  Example: koios-cli api policy_asset_list 750900e4999ebe0d58f19b634768ba25e525aaf12403bfe8fe130501
  `)

	cmd.Do(func(sess *happy.Session, args happy.Args) error {
		opts, err := c.newRequestOpts(sess, args)
		if err != nil {
			return nil
		}
		res, err := c.koios().GetPolicyAssetList(sess, koios.PolicyID(args.Arg(0).String()), opts)
		apiOutput(c.noFormat, res, err)
		return err
	})

	return cmd
}
