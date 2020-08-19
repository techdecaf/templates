# Template Functions

- [Template Functions](#template-functions)
  - [Examples](#examples)
    - [Operating System Helpers](#operating-system-helpers)
    - [File and Path Functions](#file-and-path-functions)
    - [String Functions](#string-functions)
    - [Execution Functions](#execution-functions)
      - [EXEC](#exec)
      - [TRY](#try)
      - [Expand](#expand)
      - [ExpandFile](#expandfile)
      - [JQ](#jq)
      - [YQ](#yq)

## Examples

### Operating System Helpers

- OS `{{OS}}` resolves to current running os `runtime.GOOS`
- ARCH `{{ARCH}}` resolves to current processor architecture `runtime.GOARCH`

### File and Path Functions

- PWD `{{PWD}}` resolves to current working directory
- FromSlash `{{FromSlash}}` uses go function `filepath.FromSlash`
- ToSlash `{{ToSlash}}` uses go function `filepath.ToSlash`
- ReadFile `{{ReadFile}}` read file contents, supports relative or absolute paths.
- WriteFile `{{WriteFile 'file/path' 'string data'}}` write file contents, supports relative or absolute paths.
- RemoveAll `{{RemoveAll 'file/path'}}` **CAUTION** removes file or folder and all children
- MkdirAll `{{MkdirAll}}`
- Touch `{{Touch}}`

### String Functions

- ToTitle `{{ToTitle}}` uses go function `strings.Title`
- ToUpper `{{ToUpper}}` uses go function `strings.ToUpper`
- ToLower `{{ToLower}}` uses go function `strings.ToLower`
- Replace `{{Replace}}` uses go function `strings.Replace`

### Execution Functions

#### EXEC

```text
{{EXEC `echo "hello"`}}
```

#### TRY

```text
{{TRY `fails` | default `default value`}}
```

#### Expand

```text
{{Expand}}
```

#### ExpandFile

```text
{{ReadFile `my-template-file.yaml` | ExpandFile | WriteFile}}
```

#### JQ

```text
{{JQ 'key.value[0]', '{"key": {"value":[1]} }'}}
```

```text
{{ReadFile `my-file.json` | JQ `key.value[0]`}}
```

#### YQ

```text
{{ReadFile `my-file.yaml` | YQ `key.value[0]`}}
```
