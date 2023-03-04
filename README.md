# oh-my-go

Make your terminal stylish, functional, and **fast**.

None of that **bloated** nonsense.

Made to replace [oh-my-posh](https://ohmyposh.dev/) because it was too slow and complicated.

## Table of contents

- [oh-my-go](#oh-my-go)
  - [Table of contents](#table-of-contents)
  - [What is oh-my-go?](#what-is-oh-my-go)
  - [Installation](#installation)
    - [Obtaining the binary](#obtaining-the-binary)
      - [Global install (unix)](#global-install-unix)
      - [Local install (unix)](#local-install-unix)
    - [Installing it to your shell](#installing-it-to-your-shell)
      - [ZSH](#zsh)
      - [Bash](#bash)
      - [Fish](#fish)
      - [Powershell](#powershell)
      - [Other POSIX shells](#other-posix-shells)
      - [CMD](#cmd)
  - [Customization](#customization)
    - [Plugins](#plugins)
    - [List of plugins](#list-of-plugins)
  - [Hostname \& User](#hostname--user)
  - [Prompt customization](#prompt-customization)

## What is oh-my-go?

oh-my-go is a shell engine that is designed to be fast and opinionated.

It is **not** customizable (unless you know go), which allows it to be faster than other shell engines.

It supports a wide variety of shells.

Additionally, oh-my-go changes your PS1 and **that's it**. No confusing aliases,
no unexplained flags, nothing.

It supports one singular theme: robbyrussels, because it's functional and pretty.

## Installation

### Obtaining the binary

First, download a suitable binary from the releases, or clone the repository and compile it yourself by running `make`.

Then, move the binary into PATH.

#### Global install (unix)

```sh
sudo mkdir -p /usr/local/bin
sudo mv oh-my-go* /usr/local/bin/oh-my-gos
sudo chmod +x /usr/local/bin/oh-my-go
```

#### Local install (unix)

> **Warning**
> You may see an error that oh-my-go wasn't found, if you see this you may not have `~/.local/bin/` in your PATH.
>
> [Here](https://github.com/talwat/pap#local-installation-not-found) there are instructions on how to add it from another one of my projects.

```sh
sudo mkdir -p ~/.local/bin
sudo mv oh-my-go* ~/.local/bin/oh-my-gos
sudo chmod +x ~/.local/bin/oh-my-go
```

### Installing it to your shell

Keep and mind if you don't have `oh-my-go` in your path you can replace it with the location.

So for zsh, instead of:

```sh
echo 'set -o PROMPT_SUBST' >> .zshrc
echo 'PS1='"'\$( oh-my-go prompt --shell=zsh --pwd=\$PWD --user=\$USER --exitcode=\$? --hostname=\$HOST )'" >> ~/.zshrc
```

you could do:

```sh
echo 'set -o PROMPT_SUBST' >> .zshrc
echo 'PS1='"'\$( /path/to/oh-my-go prompt --shell=zsh --pwd=\$PWD --user=\$USER --exitcode=\$? --hostname=\$HOST )'" >> ~/.zshrc
```

#### ZSH

```sh
echo 'set -o PROMPT_SUBST' >> .zshrc
echo 'PS1='"'\$( oh-my-go prompt --shell=zsh --pwd=\$PWD --user=\$USER --exitcode=\$? --hostname=\$HOST )'" >> ~/.zshrc
```

#### Bash

```sh
echo 'PS1='"'\$( oh-my-go prompt --shell=bash --pwd=\$PWD --user=\$USER --exitcode=\$? --hostname=\$HOSTNAME )'" >> ~/.bashrc
```

#### Fish

```sh
mkdir -p ~/.config/fish/

echo >> ~/.config/fish/config.fish 'if status is-interactive
    function fish_prompt
        printf (oh-my-go prompt --shell=fish --pwd=\$PWD --user=\$USER --exitcode=\$status --hostname=(prompt_hostname))
    end
end
'
```

#### Powershell

You need to add:

```powershell
function prompt {
  "$(\Location\To\oh-my-go prompt --shell=powershell --pwd=$PWD --user=$UserName --exitcode=$LASTEXITCODE) --hostname=$COMPUTERNAME"
}
```

to your `$PROFILE`.

I didn't create a specific command because powershell is garbage.

#### Other POSIX shells

You should add:

```sh
PS1='$( oh-my-go prompt --shell=sh --pwd=$PWD --user=$USER --exitcode=$? --hostname=$HOST )'
```

to your shell's rc/profile file.

Please open an [issue](https://github.com/talwat/oh-my-go) and I might add instructions/support for that shell.

#### CMD

Nope. You're on your own.

## Customization

### Plugins

You can use specific plugins by setting `OMGO_PLUGINS` to a string which has all your plugins
seperated by spaces. Like so:

```sh
export OMGO_PLUGINS="git node python"
```

### List of plugins

- `git`
- `python`
- `ruby`
- `node`
- `rust`

## Hostname & User

If you use a lot of different devices, you can set `OMGO_SHOW_USER_HOSTNAME=1` which will add an extra bit
at the beggining of your prompt which displays your username and hostname.

```sh
export OMGO_SHOW_USER_HOSTNAME=1
```

> **Note** This is an "unofficial" modification to the theme which was not explicitly defined by
> robbyrussell himself.

## Prompt customization

If you're even slightly familiar with golang, this shouldn't be too bad.

You need to edit `internal/prompt/prompt.go` and from there you can make small adjustments like change
some characters or add a newline and so on so fourth.

Then, recompile & install oh-my-go.
