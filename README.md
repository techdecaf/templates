# Template Functions

- [Template Functions](#template-functions)
  - [Examples](#examples)
    - [Operating System Helpers](#operating-system-helpers)
    - [File and Path Functions](#file-and-path-functions)
    - [String Functions](#string-functions)
    - [Execution Functions](#execution-functions)

## Examples

### Operating System Helpers

- OS `{{OS}}` resolves to current running os `runtime.GOOS`
- ARCH `{{ARCH}}` resolves to current processor architecture `runtime.GOARCH`

### File and Path Functions

- PWD `{{PWD}}` resolves to current working directory
- FromSlash `{{FromSlash}}` `filepath.FromSlash`
- ToSlash `{{ToSlash}}` `filepath.ToSlash`
- ReadFile `{{ReadFile}}` read file contents, supports relative or absolute paths.
- WriteFile `{{WriteFile 'file/path' 'string data'}}` write file contents, supports relative or absolute paths.
- MkdirAll `{{MkdirAll}}`
- Touch `{{Touch}}`

### String Functions

- ToTitle `{{ToTitle}}` uses go function `strings.Title`
- ToUpper `{{ToUpper}}` uses go function `strings.ToUpper`
- ToLower `{{ToLower}}` uses go function `strings.ToLower`
- Replace `{{Replace}}` uses go function `strings.Replace`

### Execution Functions

- EXEC `{{EXEC 'echo hello'}}`
- TRY `{{TRY 'fails' | default 'default value'}}`
- Expand `{{Expand}}`
- ExpandFile `{{ExpandFile}}`
- JQ `{{JQ 'key.value[0]', '{"key": {"value":[1]} }'}}` === "1"
