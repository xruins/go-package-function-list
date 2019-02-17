package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/xruins/go-package-function-list/function"
)

// CmdOptions defines commanline options for jessevdk/go-flags.
type CmdOptions struct {
	Dir        string `short:"d" long:"dir" description:"directory to parse."`
	Regex      string `short:"r" long:"regex" default:"" description:"regexp to filter results. it applies after filter by 'suffix'."`
	Suffix     string `short:"s" long:"suffix" default:" " description:"suffix to filter results. it applies before filter by 'regexp'."`
	Bound      string `short:"b" long:"delimiter" default:" " description:"delimiter among function names"`
	PublicOnly bool   `short:"p" long:"public-only" description:"shows only public methods"`
}

func do(opts *CmdOptions) (string, error) {
	fnames, err := function.ParseDir(opts.Dir)
	if err != nil {
		return "", fmt.Errorf("failed to parse package. err: %s", err)
	}

	if opts.PublicOnly {
		fnames = function.FilterPublicMethod(fnames)
	}
	if opts.Suffix != "" {
		fnames = function.FilterBySuffix(fnames, opts.Suffix)
	}
	if opts.Regex != "" {
		newFnames, err := function.FilterByRegexp(fnames, opts.Regex)
		if err != nil {
			return "", fmt.Errorf("failed to parse regexp: %s", err)
		}
		fnames = newFnames
	}
	bounded := strings.Join(fnames, opts.Bound)
	return bounded, nil
}

func main() {
	var opts CmdOptions
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			panic(err)
		}
	}

	out, err := do(&opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(out)
	os.Exit(0)
}
