<!-- DO NOT REMOVE - contributor_list:data:start:["gleich", "dependabot-preview[bot]", "fharper", "khrj"]:end -->

![Title Example](title.gif)

# nuke ![GitHub release (latest by date)](https://img.shields.io/github/v/release/gleich/nuke)

![build](https://github.com/gleich/nuke/workflows/build/badge.svg)
![lint](https://github.com/gleich/nuke/workflows/lint/badge.svg)
![release](https://github.com/gleich/nuke/workflows/release/badge.svg)

☢️ Force quit all applications with one terminal command.

## 🚀 Install

### 🍎 MacOS

Simply run `brew install gleich/homebrew-taps/nuke`

### 🐧 Linux

Download the binary from the [releases page](https://github.com/gleich/nuke/releases) and copy it to your desired location.

#### Requirements

- Please install the `wmctrl` CLI tool using your system's package manager

## 🏃 Running

Just run `nuke` and answer the one question to confirm. A window will pop up when you run it from your terminal for the first time asking if your terminal is allowed to close finder windows. Click `ok` to continue. You should only be asked this once.

## ⚙️ Configuration

### 🙈 Ignoring apps

You can configure a list of apps you want nuke to ignore in `~/.config/nuke/config.yml`. You will need to create this file. Below is an example config:

```yml
ignoredApps:
  - Google Chrome
  - Terminal
```

_Ignores Google Chrome and Terminal_

If you want to ignore apps on the fly then you can pass them in as arguments. Keep in mind that spaces and other characters might need to be escaped. Below are two examples

```bash
$ nuke "Google Chrome" "Visual Studio Code"
```

_Ignores Google Chrome and Visual Studio Code_

```bash
$ nuke Music Slack Notion
```

_Ignores Music, Slack, and Notion_

### 🚀 Update Checks

By default nuke checks if there is an update every time you run it. If you want to turn it off add the following to your config:

```yaml
ignoredUpdates: true
```

## 🙋‍♀️🙋‍♂️ Contributing

All contributions are welcome! Just make sure that its not an already existing issue or pull request.

<!-- DO NOT REMOVE - contributor_list:start -->
## 👥 Contributors


- **[@gleich](https://github.com/gleich)**

- **[@dependabot-preview[bot]](https://github.com/apps/dependabot-preview)**

- **[@fharper](https://github.com/fharper)**

- **[@khrj](https://github.com/khrj)**

<!-- DO NOT REMOVE - contributor_list:end -->
