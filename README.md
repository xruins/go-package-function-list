# Overview

A command-line tool to get functions of go source file on specific directory.

# Usage

```
Usage:
  go-package-function-list [OPTIONS] DIRECTORY

Application Options:
  -x= regexp to filter functions
  -p= a prefix to filter functions
  -s= a suffix to filter functions
  -d= delimiter among function names (default:  )
  -o  whether shows only public methods or not
  -r  parses directory recursively if true
```

# Build

```
GO111MODULE=on go build ./...
```

# License

MIT
