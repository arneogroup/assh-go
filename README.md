# *A*rneo*SSH* command

Allows you to fuzzy search your servers from a consul inventory.

Generates arguments for the UNIX `ssh` command, based on a consul node inventory
Custom meta fields can be used to automatically update ssh arguments

## Features

By default when calling the binary without parameters, you will be promped with a fuzzy search (powered by fzf) to search through your consul inventory.
Once you've selected the server you want to connect to, `assh-go` will prompt you the needed parameters for the `ssh` command.
If you pass two parameters to the binary instead of generating you default `ssh` params, it will allow you to generate basic `rsync` params. You'll just have to replace your server with `_srv_` and `assh-go` will understand you want to insert it here.

## How to integrate

This tool generates ssh arguments, therefore you have to inject them to a real `ssh` command
You can do this in many ways. The simplest way is inline of your command such as:
```bash
ssh $(assh-go)
```
But you can easily export an alias or bash script under your `$PATH` like so:
```bash
#!/usr/bin/env bash

SSHARGS=$(assh-go)
echo ssh $SSHARGS
ssh $SSHARGS
```

## Conf file structure

A simple configuration file is required for `assh-go` to run.

You need to specify a `map[string]string` of jumpServers with name. This is made so if you have a custom meta with your customer's name it will try to find the right jumpServer to use.
As for now this feature is mendatory, but I plan to make it optionnal in the future as not anyone needs a jumpServer each time.


Here is an example of a basic conf file. It should be located under `~/.assh`
```
{
  "jumpServer": {
  "firstJump": "jumpUser@1.1.1.1",
  "default": "jumpyjack@2.2.2.2"
  },
  "defaultUsername": "ubuntu"
}

```

## Consul custom meta fields

- `meta.private_ip`: defines a private IP to use for connection instead of the default `address` field
- `meta.jump_server`: defines a custom jump server to use (blab and adfab jump servers are identified via the `meta.customer` field

## Contribution

You are more than welcome to open an Issue or a PR if you have any idea of improvement.
Many features in `assh-go` may seems a bit weired but keep in mind that I developed this first for our needs in arneo. It is primarly made for us. I try to make `assh-go` more generic by allowing custom configuration file to enable or not some features, but it'll take time. If you want to give a hand, feel free to contact me or open an issue.

Cheers.
