![Title Example](title.gif)

# nuke

![Go Build](https://github.com/Matt-Gleich/nuke/workflows/Go%20Build/badge.svg) ![Go Format](https://github.com/Matt-Gleich/nuke/workflows/Go%20Format/badge.svg) ![goreleaser format](https://github.com/Matt-Gleich/nuke/workflows/goreleaser%20format/badge.svg) ![goreleaser](https://github.com/Matt-Gleich/nuke/workflows/goreleaser/badge.svg)

☢️ Force quit all applications with one terminal command.

## 🚀 Install

Simply run `brew tap Matt-Gleich/homebrew-taps` and then `brew install nuke`

## 🏃 Running

Just run `nuke` and answer the one question to confirm. A window will pop up when you run it from your terminal for the first time asking if your terminal is allowed to close finder windows. Click `ok` to continue. You should only be asked this once.

## ⚙️ Configuration

You can configure a list of apps you want nuke to ignore by creating a directory called `nuke` inside the `~/.config/` directory and then a file inside of it called `config.yml`. Below is an example config:

```yml
ignored:
  - Google Chrome
  - Terminal
```

## 🙋‍♀️🙋‍♂️ Contributing

All contributions are welcome! Just make sure that its not an already existing issue or pull request.
