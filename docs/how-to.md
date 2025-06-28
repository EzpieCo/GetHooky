# How To 

## Add hooks

Adding hooks is very simple. Just use the `add` cmd.

```bash
hooky add pre-commit "echo 'hello world\!'"
```

## Uninstall all hooks

Uninstalling hooks is very simple. Just use the `uninstall` cmd.

```bash
hooky uninstall
```

## Ignore perticular hooks

Ignoring a perticular hook which you don't want to run can be achived with
`ignore` cmd.

```bash
hooky ignore pre-commit
```

You can unignore a hook as well using the `unignore` cmd.

```bash
hooky unignore pre-commit
```
