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
	args, err := p.Parse()
	if err != nil {
		fmt.Errorf("failed to parse arguments. err: %s", err)
		os.Exit(1)
	}

	fnames, err := function.ParseDir(opts.Dir, opts.Recursive)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	regex := opts.Regex
	if regex != "" {
		newFnames, err := function.FilterByRegexp(fnames, regex)
		if err != nil {
			fmt.Printf("failed to filter by regexp. err: %s", err)
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

	if opts.publicOnly {
		fnames = function.FilterPublicMethod(fnames)
	}

	fmt.Println(strings.Join(fnames, opts.Delimiter))
}
