# Getting Started

## Install

::: code-group

```shell [curl]
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ezpieco/gethooky/master/tools/install.sh)"
```

```shell [wget]
sh -c "$(wget -O- https://raw.githubusercontent.com/ezpieco/gethooky/master/tools/install.sh)"
```

```shell [scoop]
scoop bucket add hooky https://github.com/ezpieco/gethooky
scoop install hooky
:::

## `hooky init`

the `init` command creates a directory called `.hooky` in your current directory.
This makes managing, version controling and sharing your git hooks easy.

```shell
hooky init
```

## `hooky add <hook-name> <command>`

Adds a new hook script with your custom command.

```shell
hooky add <hook-name> <command>
```

> 💡 Use any Git supported hook. eg - `pre-commit`, `pre-push`.

## `hooky install`

The `install` command updates your hooks into `.git/hooks/` making sure that
your hooks even work.

```shell
hooky install
```

> ⚠️ Keep in mind that whenever you add or update hooks with `hooky add` 
make sure to run `hooky install` in order for them to work.

## Try it out

Congratulations! You have sucessfully setup GetHooky in your project 🎉!
Now ahead and try it out!

```shell
git commit -m "Don't crash rookie"
```
