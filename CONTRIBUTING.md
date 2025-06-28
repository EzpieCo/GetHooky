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

## Stick to the same page

To make sure we all are on the same page and everybody understands each other.

Please make sure to follow these rules before doing anything:

- Please make sure once you are done coding, make a PR on the `develop` branch **only**.
- Most of the time you should not write test suites unless you have added a feature that needs a new test.
- Please make sure that the logic is separate from the cobra command. Basically, files inside of `cmd/` directory shouldn't handle logic, only connecting.
- Please make sure all logic code goes inside of `internal/core/core.go` file **only**.
- Most of the time when adding a new command please make sure that your code follows this standard:
  ```go
  pwd, err := os.Getwd()
  if err != nil {
    fmt.Printf("❌ Failed to get current directory path:\n %v\n", err)
    return
  }

  if err := core.FunctionName(pwd);
  ```
  This means your logic function should take a basePath, making it easier to write test.
- Make sure that you put this comment on top of the code that you write:
```go
/*
CONTRIBUTOR - @veryCoolUsername <email if you want (optional)>
*/
```
- Please follow **PascalCase** for functions, and **camelCase** for variables.
- Finally the least, most, important thing, name your variables and functions smartly. We are all rookies here.

## Code of Conduct
Be nice. Be helpful. We're all rookies at most things.
