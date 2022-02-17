<h1>Koios CLI</h1>

**[Koios API] is Elastic Cardano Query Layer!**

> A consistent query layer for developers to build upon Cardano, with   
> multiple, redundant endpoints that allow $$for easy scalability.

**Build Status**

[![linux](https://github.com/cardano-community/koios-cli/workflows/linux/badge.svg)](https://github.com/cardano-community/koios-cli/actions/workflows/linux.yml)
[![macos](https://github.com/cardano-community/koios-cli/workflows/macos/badge.svg)](https://github.com/cardano-community/koios-cli/actions/workflows/macos.yml)
[![windows](https://github.com/cardano-community/koios-cli/workflows/windows/badge.svg)](https://github.com/cardano-community/koios-cli/actions/workflows/windows.yml)

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
NAME:
   koios-cli api - Interact with Koios API REST endpoints

USAGE:
   koios-cli api command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command
   ACCOUNT:
     account-list       Get a list of all accounts returns array of stake addresses.
     account-info       Get the account info of any (payment or staking) address.
     account-rewards    Get the full rewards history (including MIR) for a stake address, or certain epoch if specified.
     account-updates    Get the account updates (registration, deregistration, delegation and withdrawals).
     account-addresses  Get all addresses associated with an account payment or staking address
     account-assets     Get the native asset balance of an account.
     account-history    Get the staking history of an account.
   ADDRESS:
     address-info    Get address info - balance, associated stake address (if any) and UTxO set.
     address-txs     Get the transaction hash list of input address array, optionally filtering after specified block height (inclusive).
     address-assets  Get the list of all the assets (policy, name and quantity) for a given address.
     credential-txs  Get the transaction hash list of input payment credential array, optionally filtering after specified block height (inclusive).
   ASSET:
     asset-list          Get the list of all native assets (paginated).
     asset-address-list  Get the list of all addresses holding a given asset.
     asset-info          Get the information of an asset including first minting & token registry metadata.
     asset-summary       Get the summary of an asset (total transactions exclude minting/total wallets include only wallets with asset balance).
     asset-txs           Get the list of all asset transaction hashes (newest first).
   BLOCK:
     blocks      Get summarised details about all blocks (paginated - latest first).
     block-info  Get detailed information about a specific block.
     block-txs   Get a list of all transactions included in a provided block.
   EPOCH:
     epoch-info    Get the epoch information, all epochs if no epoch specified.
     epoch-params  Get the protocol parameters for specific epoch, returns information about all epochs if no epoch specified.
   NETWORK:
     tip      Get the tip info about the latest block seen by chain.
     genesis  Get the Genesis parameters used to start specific era on chain.
     totals   Get the circulating utxo, treasury rewards, supply and reserves in lovelace for specified epoch, all epochs if empty.
   POOL:
     pool-list        A list of all currently registered/retiring (not retired) pools.
     pool-infos       Current pool statuses and details for a specified list of pool ids.
     pool-info        Current pool status and details for a specified pool by pool id.
     pool-delegators  Return information about delegators by a given pool and optional epoch (current if omitted).
     pool-blocks      Return information about blocks minted by a given pool in current epoch (or _epoch_no if provided).
     pool-updates     Return all pool updates for all pools or only updates for specific pool if specified.
     pool-relays      A list of registered relays for all currently registered/retiring (not retired) pools.
     pool-metadata    Metadata(on & off-chain) for all currently registered/retiring (not retired) pools.
   SCRIPT:
     script-list       List of all existing script hashes along with their creation transaction hashes.
     script-redeemers  List of all redeemers for a given script hash.
   TRANSACTIONS:
     txs-infos      Get detailed information about transaction(s).
     tx-info        Get detailed information about single transaction.
     tx-utxos       Get UTxO set (inputs/outputs) of transactions.
     txs-metadata   Get metadata information (if any) for given transaction(s).
     tx-metadata    Get metadata information (if any) for given transaction.
     tx-metalabels  Get a list of all transaction metalabels.
     tx-submit      Submit signed transaction to the network.
     txs-statuses   Get the number of block confirmations for a given transaction hash list
     tx-status      Get the number of block confirmations for a given transaction hash
   UTILS:
     get   get issues a GET request to the specified API endpoint
     head  head issues a HEAD request to the specified API endpoint

OPTIONS:
   --port value, -p value  Set port (default: 443)
   --host value            Set host (default: "api.koios.rest")
   --api-version value     Set API version (default: "v0")
   --schema value          Set URL schema (default: "https")
   --origin value          Set Origin header for requests. (default: "https://github.com/cardano-community/koios-go-client")
   --rate-limit value      Set API Client rate limit for outgoing requests (default: 5)
   --no-format             prints response json strings directly without calling json pretty. (default: false)
   --enable-req-stats      Enable request stats. (default: false)
   --testnet               use default testnet as host. (default: false)
   --help, -h              show help (default: false)
```

</details>

### Example Usage

```shell
koios-cli api --enable-req-stats tip
```

**response**

```json
{
  "request_url": "https://api.koios.rest/api/v0/tip",
  "request_method": "GET",
  "status_code": 200,
  "status": "200 OK",
  "date": "Mon, 07 Feb 2022 12:49:49 GMT",
  "content_range": "0-0/*",
  "stats": {
    "req_started_at": "2022-02-07T12:49:48.565834833Z",
    "req_dns_lookup_dur": 1284269, // dns lookup duration in nanosecons.
    "tls_hs_dur": 208809082, // handshake duration in nanosecons.
    "est_cxn_dur": 159857626, // time it took to establish connection with server in nanosecons.
    "ttfb": 998874037, // time since start of the request it took to recieve first byte.
    "req_dur": 999186595, // total request duration in nanoseconds
    "req_dur_str": "999.186595ms" // string of req_dur
  },
  "data": {
    "abs_slot": 52671876,
    "block_no": 6852764,
    "block_time": "2022-02-07T12:49:27",
    "epoch": 319,
    "epoch_slot": 227076,
    "hash": "1dad134750188460dd48068e655b5935403d2f51afaf53a39337a4c89771754a"
  }

```

#### Example to query testnet tip from cli

```shell
koios-cli api --enable-req-stats --testnet tip
# OR
koios-cli --enable-req-stats --host testnet.koios.rest tip
```

**response**

```json
{
  "request_url": "https://testnet.koios.rest/api/v0/tip",
  "request_method": "GET",
  "status_code": 200,
  "status": "200 OK",
  "date": "Mon, 07 Feb 2022 12:50:04 GMT",
  "content_range": "0-0/*",
  "stats": {
    "req_started_at": "2022-02-07T12:50:03.98615637Z",
    "req_dns_lookup_dur": 1383437,
    "tls_hs_dur": 69093093,
    "est_cxn_dur": 43733700,
    "ttfb": 167423049,
    "req_dur": 167738287,
    "req_dur_str": "167.738287ms"
  },
  "data": {
    "abs_slot": 49868948,
    "block_no": 3300758,
    "block_time": "2022-02-07T12:49:24",
    "epoch": 185,
    "epoch_slot": 318548,
    "hash": "d7623e68cb78f450f42ba4b5a169124b26677f08f676ca4241b27edb6dbf0071"
  }
}
```

## Install

It's highly recommended installing a latest version of `koios-cli` available on the [releases page](https://github.com/cardano-community/koios-cli/releases/latest).

## Install from Source

```shell
go install github.com/cardano-community/koios-cli@latest
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
