# Overview

A commandline tool to get functions on specified directory.

# Usage

```
Usage:
  go-package-function-list [OPTIONS]

Application Options:
  -d, --dir=         directory to parse.
  -r, --regex=       regexp to filter results. it applies after filter by 'suffix'.
  -s, --suffix=      suffix to filter results. it applies before filter by 'regexp'. (default:  )
  -b, --delimiter=   delimiter among function names (default:  )
  -p, --public-only  shows only public methods

Help Options:
  -h, --help         Show this help message
```

# Build

```
GO111MODULES=on go build ./...
```

# License

MIT