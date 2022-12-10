# mcli

Simple shortcut menu for shell commands

![](./.github/demo.gif)

## Features

* One simple JSON as storage(`~/.config/mcli/config.json`)
* Search by Name/Command
* Fast and Lightweight

## Installation

#### Arch Linux

```bash
paru -S mcli-bin
```

#### DEB/RPM/Alpine

Packages available on the [Releases](https://github.com/mazzz1y/mcli/releases) page

#### macOS

```bash
brew tap mazzz1y/tap
brew install mcli
```

## Usage

```
USAGE:
   mcli [global options] command [command options] [arguments...]

COMMANDS:
   add, a          Add shortcut
   delete, d       Remove shortcut
   prompt-size, p  Set prompt size
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

