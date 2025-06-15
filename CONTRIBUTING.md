# Contributing to GetHooky

Thanks for wanting to add value to this tool!

## Getting started

1. Clone repo:
```bash
git clone https://github.com/ezpieco/gethooky.git
cd gethooky
```
2. Install dependencies:
```bash
go mod tidy
```
3. Try it out
```bash
go run main.go init
go run main.go add pre-commit "echo 'help me'"
go run main.go install
```
4. Build the CLI(optional)
```bash
go build -o hooky .
```

## Project structure
```
cmd/               Cobra commands
  init.go
  add.go
  install.go
  root.go
internal/core/     Logic
  core.go
  core_test.go
builds/            cross platform binaries(git ignored)
```

## Running test

By default, `hooky` is installed for this repo, and assuming you have installed hooky too, just run `hooky install` to setup the git hook for running tests before committing.
But if you still want, run test with:

```bash
go test ./... -v
```

## Your First Contribution
Here are some great ways to start:

- Fix a bug or add a feature (see Issues)
- Improve logging or messages
- Enhance this file ;)
- Or do anything you want! I'll just reject it ;)

## Code of Conduct
Be nice. Be helpful. We're all rookies at most things.
