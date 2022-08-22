# paste-cli

This is a simple CLI tool written in Go to interact with a
[paste-server](https://github.com/h5law/paste-server) either a self hosted
instance or the main hosted site.

## Install

To install and run this project simply clone the repo:
```
git clone https://github.com/h5law/paste-cli
```

Then install the dependencies and build the binary:
```
cd paste-cli
go mod tidy
go build -o paste
```

You can then run the command with `./paste` or move it into your `$PATH` and
execute it via `paste` anywhere on the computer.

## Config

The paste command will look for a config file (by default `$HOME/.paste.yaml`)
which can control the url of the paste-server instance to interact with.

The document should look like this:
```
url: "<paste-server instance url goes here>:<port to use here>"
```

So for local instances on port 3000 it would look like:
```
url: "http://127.0.0.1:3000"
```

You can also use the `--config` flag to use a config file elsewhere or if no
config file is found / given the paste command will default to using the main
server URL.

## TODO
- Add update subcommand and functionality
- Add delete subcommand and functionality
