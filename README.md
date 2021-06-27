# Go CLI Template

A template for writing Go CLI

## Requirements

1. Python 3.6+
   1. Install [cookiecutter](https://github.com/cookiecutter/cookiecutter)
2. Go 1.16+

## Usage

In your terminal type in

```sh
> cookiecutter https://github.com/akshaybabloo/go-cli-template
```

Then follow the options on the terminal.

Once done, you will have to refactor the module path in `go.mod` to suit your project name

## Features

1. Checks for updates once every 24 hours
2. Uses [Cobra](https://github.com/spf13/cobra) package for CLI
3. Factory based approach, functions are available on every command if needed
4. Global debug flag - uses [logrus](https://github.com/sirupsen/logrus)
5. Custom help output that also displays aliases
6. Uses custom color that's available via factory - uses [color](https://github.com/fatih/color)
7. Default global configuration location - `<user folder>/config.yaml`
   1. Disable colour usage globally
