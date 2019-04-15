package main

import (
	"fmt"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/xruins/go-package-function-list/function"
)

// CmdOptions represents command-line options for go-package-function-list
type cmdOptions struct {
	Regex      string `short:"x" description:"regexp to filter functions"`
	Prefix     string `short:"p" description:"a prefix to filter functions"`
	Suffix     string `short:"s" description:"a suffix to filter functions"`
	Delimiter  string `short:"d" default:" " description:"delimiter among function names"`
	PublicOnly bool   `short:"o" description:"whether shows only public methods or not"`
	Recursive  bool   `short:"r" description:"parses directory recursively if true"`
}

func main() {
	opts := &cmdOptions{}
	p := flags.NewParser(opts, flags.PrintErrors)
	p.Name = "go-package-function-list"
	p.Usage = "[OPTIONS] DIRECTORY"
	args, err := p.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse arguments. err: %s", err)
		os.Exit(1)
	}

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Wrong number of arguments.")
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}
	dir := args[0]

	fnames, err := function.ParseDir(dir, opts.Recursive)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse package:", err)
		os.Exit(1)
	}

	regex := opts.Regex
	if regex != "" {
		newFnames, err := function.FilterByRegexp(fnames, regex)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to filter by regexp. err: %s", err)
			os.Exit(1)
		}
		fnames = newFnames
	}

	prefix := opts.Prefix
	if prefix != "" {
		fnames = function.FilterByPrefix(fnames, prefix)
	}

	suffix := opts.Suffix
	if suffix != "" {
		fnames = function.FilterBySuffix(fnames, suffix)
	}

	if opts.PublicOnly {
		fnames = function.FilterPublicMethod(fnames)
	}

	if len(fnames) > 0 {
		fmt.Println(strings.Join(fnames, opts.Delimiter))
	}
}
