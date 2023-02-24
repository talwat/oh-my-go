# oh-my-go

Make your terminal stylish, functional, and **fast**.

## What is oh-my-go?

oh-my-go is a shell engine that is designed to be fast and opinionated.

It is **not** customizable (unless you know go), which allows it to be faster than other shell engines.

It supports a wide variety of shells.

## Installation

### Obtaining the binary

First, download a suitable binary from the releases, or clone the repository and compile it yourself by running `make`.

Then, move the binary into PATH.

#### Global install

```sh
sudo mkdir -p /usr/local/bin
sudo mv oh-my-go* /usr/local/bin/oh-my-gos
sudo chmod +x /usr/local/bin/oh-my-go
```

#### Local install

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

#### ZSH

```sh
echo 'set -o PROMPT_SUBST' >> .zshrc
echo 'PS1='"'\$( oh-my-go prompt zsh \$PWD \$? )'" >> ~/.zshrc
```

#### Bash

```sh
echo 'PS1='"'\$( oh-my-go prompt bash \$PWD \$? )'" >> ~/.bashrc
```

#### Fish

```sh
mkdir -p ~/.config/fish/

echo >> ~/.config/fish/config.fish 'if status is-interactive
    function fish_prompt
        printf (oh-my-go prompt fish $PWD $status)
    end
end
'
```

#### Powershell

You need to add:

```powershell
function prompt {
    "$(oh-my-go prompt powershell $PWD $LASTEXITCODE)"
}
```

to your `$PROFILE`.

I didn't create a specific command because powershell is garbage.

#### Other POSIX shells

You should add:

```sh
PS1='$( oh-my-go prompt sh $PWD $? )'
```

to your shell's rc/profile file.

Please open an [issue](https://github.com/talwat/oh-my-go) and I might add instructions/support for that shell.

#### CMD

Nope. You're on your own.
