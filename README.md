# 🪝 GetHooky
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
![latest version](https://img.shields.io/github/v/release/ezpieco/gethooky?style=for-the-badge)
![latest release](https://img.shields.io/github/downloads/ezpieco/gethooky/total?style=for-the-badge)
![GitHub License](https://img.shields.io/github/license/ezpieco/gethooky?style=for-the-badge)

Inspired from [husky](https://github.com/typicode/husky) but built for **every stack**. Because git is git whether you code in python, rust, go, or nodejs.

![logo](./assets/getHooky.svg)

> *Because the intern shouldn't `git push` broken, untested, unlinted, code to production* - wise old programmer, 2025

Get hooks but for every stack - Python, Node, Rust, Go, anything.

GetHooky will make sure that you run your commands before(or after) doing any git command.

**NOTE:** GetHooky doesn't have control over what you do, so use wisely.

## 🚀 Features

-  Cross platform.
- ✅ Works with any language stack(python, go, rust, node)
- 👶 Very easy to use, even a 5 year old can use it.
-  Store hooks in a version controlable manner, inside `.hooky` directory
- 🧠 manages only hooks it has access to hooks with `# hookie ya rookie`
- 🛠 Easy to install, update and share hooks with your team

## ⚡ Blazingly Fast Example

Prevent interns from pushing broken commits

```
$ hooky init
$ hooky add pre-commit "pytest"
$ hooky install
```

And then try running commit and let the black magic work!

##  Installation

| Method | Command |
| ------ | ------- |
| curl   | `sh -c "$(curl -fsSL https://raw.githubusercontent.com/ezpieco/gethooky/master/tools/install.sh)"`|
| wget   | `sh -c "$(wget -O- https://raw.githubusercontent.com/ezpieco/gethooky/master/tools/install.sh)"`|
| scoop(windows recommended)   | `scoop bucket add hooky https://github.com/ezpieco/gethooky && scoop install hooky`|

## ⚡ Quick Start

```bash
hooky init
```
Creates a `.hooky` directory in your current directory
```bash
hooky add <hook-name> <command>
```
Creates a `.hooky/<hook-name>` file with `<command>` in it.
```bash
hooky install
```
Installs all `.hooky/*` hooks into `.git/hooks/*` with custom hooks in mind.

## Why not X?

|    Tool    |   Language   | Git hook sharing |
|------------|--------------|------------------|
|    Husky   | Node.js only |        Yes       |
| pre-commit |  Python only |        Yes       |
|  GetHooky  | Any language |        Yes       |

## 🧠 How It Works
GetHooky stores all your hooks inside of `.hooky` where you can version control them and share it with your team.

When you run `hooky install` it generates a `.git/hooky/<hook-name>` with the following content:

```bash
#!/bin/sh
# hooky ya rookie

# your command here

if [ $? -ne 0 ]; then
  echo ""
  echo "🚫 Hook '<hook-name>' failed."
  echo "👉 To bypass, use: git commit --no-verify"
  echo ""
  exit 1
fi
```

Only files with `# hooky ya rookie` are controlled by GetHooky, your custom ones are always skipped.

## 🌐 Docs
You can read the full docs at [here](https://ezpieco.github.io/GetHooky/)

⭐ If you find GetHooky useful, consider starring the repo!

## Thank you

[![Forkers repo roster for @ezpieco/gethooky](https://reporoster.com/stars/dark/ezpieco/gethooky)](https://github.com/ezpieco/gethooky/network/members)

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/flowXM"><img src="https://avatars.githubusercontent.com/u/54032222?v=4?s=100" width="100px;" alt="flowXM"/><br /><sub><b>flowXM</b></sub></a><br /><a href="https://github.com/EzpieCo/GetHooky/commits?author=flowXM" title="Code">💻</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/thatonecodes"><img src="https://avatars.githubusercontent.com/u/114317937?v=4?s=100" width="100px;" alt="Maher"/><br /><sub><b>Maher</b></sub></a><br /><a href="https://github.com/EzpieCo/GetHooky/commits?author=thatonecodes" title="Code">💻</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->
