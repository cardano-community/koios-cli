<h1>Koios CLI</h1>

**[Koios API] is Elastic Cardano Query Layer!**

> A consistent query layer for developers to build upon Cardano, with   
> multiple, redundant endpoints that allow $$for easy scalability.

**Development Status**

![GitHub last commit](https://img.shields.io/github/last-commit/cardano-community/koios-cli)
<!-- coverge -->
[![codeql](https://github.com/cardano-community/koios-cli/workflows/codeql/badge.svg)](https://github.com/cardano-community/koios-cli/actions/workflows/codeql.yml)
[![misspell](https://github.com/cardano-community/koios-cli/workflows/misspell/badge.svg)](https://github.com/cardano-community/koios-cli/actions/workflows/misspell.yml)



- [Usage](#usage)
  - [List of API commands](#list-of-api-commands)
  - [Example Usage](#example-usage)
    - [Example to query testnet tip from cli](#example-to-query-testnet-tip-from-cli)
- [Install](#install)
- [Install from Source](#install-from-source)
- [Contributing](#contributing)
  - [Code of Conduct](#code-of-conduct)
  - [Got a Question or Problem?](#got-a-question-or-problem)
  - [Issues and Bugs](#issues-and-bugs)
  - [Feature Requests](#feature-requests)
  - [Submission Guidelines](#submission-guidelines)
    - [Submitting an Issue](#submitting-an-issue)
    - [Submitting a Pull Request (PR)](#submitting-a-pull-request-pr)
    - [After your pull request is merged](#after-your-pull-request-is-merged)
  - [Coding Rules](#coding-rules)
  - [Commit Message Guidelines](#commit-message-guidelines)
    - [Commit Message Format](#commit-message-format)
    - [Revert](#revert)
    - [Type](#type)
    - [Scope](#scope)
- [| **markdown** | Markdown files |](#-markdown--markdown-files-)
    - [Subject](#subject)
    - [Body](#body)
    - [Footer](#footer)
  - [Development Documentation](#development-documentation)
    - [Setup your machine](#setup-your-machine)
- [Credits](#credits)

---

## Usage

**see `koios-cli -h` for available commands**

### List of API commands

<details>
  <summary><code>koios-cli api --help</code></summary>

```
  KOIOS CLI - v2.0.0
  Copyright © 2022 - 2024 The Cardano Community
  License: Apache-2.0
  
  Interact with Koios API REST endpoints

  koios api [flags] [subcommand]

 COMMANDS:

  ADDRESS - Query information about specific address(es)

  address_assets             Address Assets
  address_info               Address Information
  address_txs                Address Transactions
  address_utxos              Address UTxOs
  credential_txs             Transactions from payment credentials
  credential_utxos           UTxOs from payment credentials

  ASSET - Query Asset related informations

  asset_addresses            Asset Addresses
  asset_history              Asset History
  asset_info                 Asset Information (Bulk)
  asset_list                 Asset List
  asset_nft_address          Asset NFT Address
  asset_summary              Asset Summary
  asset_token_registry       Asset Token Registry
  asset_txs                  Asset Transactions
  asset_utxos                Asset UTXOs
  policy_asset_addresses     Policy Asset Address List
  policy_asset_info          Policy Asset Information
  policy_asset_list          Policy Asset List
  policy_asset_mints         Policy Asset Mints

  BLOCK - Query information about particular block on chain

  block-info                 Block Info
  block-txs                  Block Txs
  blocks                     Block List

  EPOCH - Query epoch-specific details

  epoch_block_protocols      Epoch Block Protocols
  epoch_info                 Epoch Information
  epoch_params               Epoch Parameters

  NETWORK - Query information about the network

  genesis                    Get Genesis info
  param_updates              Param Update Proposals
  reserve_withdrawals        Reserve Withdrawals
  tip                        Query Chain Tip
  totals                     Get historical tokenomic stats
  treasury_withdrawals       Treasury Withdrawals

  OGMIOS - Various stateless queries against Ogmios v6 instance

  ogmios                     NOT IMPLEMENTED

  POOL - Query information about specific pools

  pool_blocks                Pool Blocks
  pool_delegators            Pool Delegators
  pool_delegators_history    Pool Delegators History
  pool_history               Pool History
  pool_info                  Pool Information
  pool_list                  Pool List
  pool_metadata              Pool Metadata
  pool_registrations         Pool Registrations
  pool_relays                Pool Relays
  pool_retirements           Pool Retirements
  pool_stake_snapshot        Pool Stake Snapshot
  pool_updates               Pool Updates (History)

  SCRIPT - Query information about specific scripts (Smart Contracts)

  datum_info                 Datum Information
  native_script_list         Native Script List
  plutus_script_list         Plutus Script List
  script_info                Script Information
  script_redeemers           Script Redeemers
  script_utxos               Script Utxos

  STAKE ACCOUNT - Query details about specific stake account addresses

  account_addresses          Account Addresses
  account_assets             Account Assets
  account_history            Account History
  account_info               Account Information
  account_info_cached        Account Information Cached
  account_list               Account List
  account_rewards            Account Rewards
  account_txs                Account Transactions
  account_updates            Account Updates
  account_utxos              Account Utxos

  TRANSACTIONS - Query blockchain transaction details

  submittx                   NOT IMPLEMENTED
  tx_info                    Transaction Information
  tx_metadata                Transaction Metadata
  tx_metalabels              Transaction Metadata Labels
  tx_status                  Transaction Status
  utxo_info                  UTxO Info

 FLAGS:

  --api-version       Set API version - default: "v1"
  --auth              JWT Bearer Auth token generated via https://koios.rest Profile page.
  --host              Set host for the API server - default: "api.koios.rest"
  --host-eu           Use eu mainet network host - default: "false"
  --host-guildnet     Use guildnet network host - default: "false"
  --host-preprod      Use preprod network host - default: "false"
  --host-preview      Use preview network host - default: "false"
  --no-format         prints response as machine readable json string - default: "false"
  --origin            Set origin for the API server - default:
                      "https://github.com/cardano-community/koios-cli/v2"
  --port              Set port number for the API server - default: "443"
  --rate-limit        Set rate limit for the API server - default: "10"
  --scheme            Set scheme for the API server - default: "https"
  --stats             Enable request stats - default: "false"
  --timeout           Set timeout for the API server - default: "1m0s"

 GLOBAL FLAGS:

  --debug              enable debug log level. when debug flag is after the command then debug level
                       will be enabled only for that command - default: "false"
  --help         -h    display help or help for the command. [...command --help] - default: "false"
  --profile            session profile to be used - default: "public"
  --system-debug       enable system debug log level (very verbose) - default: "false"
  --verbose      -v    enable verbose log level - default: "false"
  --version            print application version - default: "false"
  -x                   the -x flag prints all the cli commands as they are executed. - default: "false"
```

</details>

### Example Usage

```shell
koios-cli api --stats tip
```

**response**

```json
{
  "request_url": "https://api.koios.rest/api/v1/tip",
  "request_method": "GET",
  "status_code": 200,
  "status": "200 OK",
  "date": "Fri, 31 May 2024 08:06:48 GMT",
  "content_range": "0-0/*",
  "stats": {
    "req_started_at": "2024-05-31T08:06:48.846603938Z",
    "req_dns_lookup_dur": 78489018,
    "tls_hs_dur": 109035482,
    "est_cxn_dur": 62920181,
    "ttfb": 636448677,
    "req_dur": 636774353,
    "req_dur_str": "636.774353ms",
    "auth": {
      "tier": "public",
      "expires": "No expiration date",
      "max_requests": 5000,
      "max_rps": 10,
      "query_timeout": "30s",
      "cors_restricted": true
    }
  },
  "data": {
    "abs_slot": 125576504,
    "block_no": 10385051,
    "block_time": 1717142795,
    "epoch_no": 488,
    "epoch_slot": 123704,
    "hash": "1dd29324ce46038be9afb9011348530e6ff2e57ad0f818dcbe12bbf498abd157"
  }
}
```

#### Example to query testnet tip from cli

```shell
koios-cli api --stats --host-preprod tip
# OR
koios-cli api --stats --host preprod.koios.rest tip
```

**response**

```json
{
  "request_url": "https://preprod.koios.rest/api/v1/tip",
  "request_method": "GET",
  "status_code": 200,
  "status": "200 OK",
  "date": "Fri, 31 May 2024 08:09:31 GMT",
  "content_range": "0-0/*",
  "stats": {
    "req_started_at": "2024-05-31T08:09:30.905348607Z",
    "req_dns_lookup_dur": 1873265,
    "tls_hs_dur": 63867504,
    "est_cxn_dur": 28092739,
    "ttfb": 159446743,
    "req_dur": 159811589,
    "req_dur_str": "159.811589ms",
    "auth": {
      "tier": "public",
      "expires": "No expiration date",
      "max_requests": 5000,
      "max_rps": 10,
      "query_timeout": "30s",
      "cors_restricted": true
    }
  },
  "data": {
    "abs_slot": 61459716,
    "block_no": 2304547,
    "block_time": 1717142916,
    "epoch_no": 146,
    "epoch_slot": 29316,
    "hash": "1f5df7af623cfc29539ee48a3cfd5ac9c93a5b85fd84066566c762a968d0f04c"
  }
}
```

## Install

It's highly recommended installing a latest version of `koios-cli` available on the [releases page](https://github.com/cardano-community/koios-cli/releases/latest).

## Install from Source

```shell
go install github.com/cardano-community/koios-cli/v2@latest
```

**verify installation**

`koios-cli --version`


## Contributing

We would love for you to contribute to [Koios API Client Library for Go][github] and help make it even better than it is today! As a contributor, here are the guidelines we would like you to follow:

 - [Code of Conduct](#code-of-conduct)
 - [Question or Problem?](#got-a-question-or-problem)
 - [Found a Bug?](#issues-and-bugs)
 - [Missing a Feature?](#feature-requests)
 - [Submission Guidelines](#submission-guidelines)
 - [Coding Rules](#coding-rules)
 - [Commit Message Guidelines](#commit-message-guidelines)
 - [Development Documentation](#development-documentation)

### Code of Conduct

Help us keep [Koios API Client Library for Go][github] open and inclusive. Please read and follow our [Code of Conduct][coc]

---

### Got a Question or Problem?

Do not open issues for general support questions as we want to keep GitHub issues for bug reports and feature requests. You've got much better chances of getting your question answered on [Koios Telegram Group](https://t.me/joinchat/+zE4Lce_QUepiY2U1)

---

### Issues and Bugs

If you find a bug in the source code, you can help us by
[submitting an issue](#submitting-an-issue) to our [GitHub Repository][github]. Even better, you can
[submit a Pull Request](#submitting-a-pull-request-pr) with a fix.

---

### Feature Requests

You can *request* a new feature by [submitting an issue](#submitting-an-issue) to our GitHub
Repository. If you would like to *implement* a new feature, please submit an issue with
a proposal for your work first, to be sure that we can use it.
Please consider what kind of change it is:

* For a **Major Feature**, first open an issue and outline your proposal so that it can be
discussed. This will also allow us to better coordinate our efforts, prevent duplication of work,
and help you to craft the change so that it is successfully accepted into the project.
* **Small Features** can be crafted and directly [submitted as a Pull Request](#submitting-a-pull-request-pr).

---

### Submission Guidelines

#### Submitting an Issue

Before you submit an issue, please search the issue tracker, maybe an issue for your problem already exists and the discussion might inform you of workarounds readily available.

You can file new issues by filling out our [new issue form](https://github.com/cardano-community/koios-cli/issues/new).

---

#### Submitting a Pull Request (PR)

Before you submit your Pull Request (PR) consider the following guidelines:

1. Search [GitHub](https://github.com/cardano-community/koios-cli/pulls) for an open or closed PR that relates to your submission. You don't want to duplicate effort.
2. Fork the [cardano-community/koios-cli][github] repo.
3. Setup you local repository

    ```shell
    git@github.com:<your-github-username>/koios-cli.git
    cd koios-cli
    git remote add upstream git@github.com:cardano-community/koios-cli.git
    ```
4. Make your changes in a new git branch and ensure that you always start from up to date main branch. **Repeat this step every time you are about to start woking on new PR**.

    e.g. Start new change work to update readme:
    ```shell
    # if you are not in main branch e.g. still on previous work branch
    git checkout main
    git pull --ff upstream main
    git checkout -b update-readme main
    ```
5. Create your patch, **including appropriate test cases**.
6. Follow our [Coding Rules](#coding-rules).
7. If changes are in source code except documentations then run the full test suite, as described in the [developer documentation](#development-documentation), and ensure that all tests pass.
8.  Commit your changes using a descriptive commit message that follows our
  [commit message conventions](#commit-message-guidelines). Adherence to these conventions
  is necessary because release notes are automatically generated from these messages.

     ```shell
     git add -A
     git commit --signoff
     # or in short
     git commit -sm"docs(markdown): update readme examples"
     ```
9. Push your branch to GitHub:

    ```shell
    git push -u origin update-readme
    ```
10. In GitHub, send a pull request to `main` branch.
* If we suggest changes then:
  * Make the required updates.
  * Re-run the test suites to ensure tests are still passing.
  * Rebase your branch and force push to your GitHub repository (this will update your Pull Request):

     ```shell
    git fetch --all
    git rebase upstream main
    git push -uf origin update-readme
    ```
That's it! Thank you for your contribution!

---

#### After your pull request is merged

After your pull request is merged, you can safely delete your branch and pull the changes from the main (upstream) repository:

* Delete the remote branch on GitHub either through the GitHub web UI or your local shell as follows:
  
    ```shell
    git push origin --delete update-readme
    ```
* Check out the main branch:
  
    ```shell
    git checkout main -f
    ```

* Delete the local branch:

    ```shell
    git branch -D update-readme
    ```
* Update your master with the latest upstream version:

    ```shell
    git pull --ff upstream main
    ```
---

### Coding Rules

To ensure consistency throughout the source code, keep these rules in mind as you are working:

* All features or bug fixes **must be tested** by one or more specs (unit-tests).
* All public API methods **must be documented**.

---

### Commit Message Guidelines

[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

We have very precise rules over how our git commit messages can be formatted. This leads to **more readable messages** that are easy to follow when looking through the **project history**. Commit messages should be well formatted, and to make that "standardized", we are using Conventional Commits. Our release workflow uses these rules to generate changelogs.

---

#### Commit Message Format

Each commit message consists of a **header**, a **body** and a **footer**.  The header has a special format that includes a **type**, a **scope** and a **subject**:

```
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

*When maintainers are merging PR merge commit should be edited:*

```
<type>(<scope>): <subject> (#pr)
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

The **header** is mandatory and the **scope** of the header is optional.

Any line of the commit message cannot be longer 100 characters! This allows the message to be easier to read on GitHub as well as in various git tools.

The footer should contain a [closing reference to an issue](https://help.github.com/articles/closing-issues-via-commit-messages/) if any.

Samples:

```
docs(markdown): update readme examples
```

```
fix(endpoint): update Tip endpoint to latest specs.

description of your change.
```

```
refactor(client): change Client GET function signature

change order of client GET method arguments.

BREAKING CHANGE: Clien.Get signature has changed
```

---

#### Revert

If the commit reverts a previous commit, it should begin with `revert: `, followed by the header of the reverted commit. In the body it should say: `This reverts commit <hash>.`, where the hash is the SHA of the commit being reverted.

---

#### Type

Must be one of the following:

* **build**: Changes that affect the build system or external dependencies (example scopes: goreleaser, taskfile)
* **chore**: Other changes that don't modify src or test files.
* **ci**: Changes to our CI configuration files and scripts.
* **dep**: Changes related to dependecies e.g. `go.mod`
* **docs**: Documentation only changes (example scopes: markdown, godoc)
* **feat**: A new feature
* **fix**: A bug fix
* **perf**: A code change that improves performance
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **revert**: Reverts a previous commit
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* **test**: Adding missing tests or correcting existing tests

---

#### Scope

The following is the list of supported scopes:

| scope | description |
| --- | --- |
| **client** | API client related changes |
| **endpoint** | Changes related to api endpoints |
| **godoc** | Go documentation |
| **markdown** | Markdown files |
---

#### Subject

The subject contains a succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize the first letter
* no dot (.) at the end
  
#### Body
Just as in the **subject**, use the imperative, present tense: "change" not "changed" nor "changes".
The body should include the motivation for the change and contrast this with previous behavior.

#### Footer
The footer should contain any information about **Breaking Changes** and is also the place to
reference GitHub issues that this commit **Closes**.

**Breaking Changes** should start with the word `BREAKING CHANGE:` with a space or two newlines. The rest of the commit message is then used for this.

A detailed explanation can be found in this [document][commit-message-format].

---

### Development Documentation

#### Setup your machine

**Prerequisites:**

* Working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).
* [golangci-lint](https://golangci-lint.run/usage/install/#local-installation) - Go linters aggregator should be installed
* [taskfile](https://taskfile.dev/#/installation) - task runner / build tool should be installed
* [svu](https://github.com/caarlos0/svu#install) - Semantic Version Util tool should be installed
* Fork the [cardano-community/koios-cli][github] repo.
* Setup you local repository

    ```shell
    git@github.com:<your-github-username>/koios-cli.git
    cd koios-cli
    git remote add upstream git@github.com:cardano-community/koios-cli.git
    ```

**Setup local env**

```shell
task setup
```

**Lint your code**

```shell
task lint
```

**Test your change**

```shell
task test
```


**View code coverage report from in browser (results from `task test`)**

```shell
task cover
```

## Credits

[![GitHub contributors](https://img.shields.io/github/contributors/cardano-community/koios-cli?style=flat-square)](https://github.com/cardano-community/koios-cli/graphs/contributors)

<sub>**Original author.**</sub>  
<sup>koios-cli was moved under Cardano Community from <a href="https://github.com/howijd/koios-rest-go-client">howijd/koios-rest-go-client</a></sup>

<!-- LINKS -->
[Koios API]: https://koios.rest "Koios API"
[coc]: https://github.com/cardano-community/.github/blob/main/CODE_OF_CONDUCT.md
[github]: https://github.com/cardano-community/koios-cli
[koios-cli]: https://github.com/cardano-community/koios-cli "cardano-community/koios-cli"

<!-- release -->
