package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/xruins/go-package-function-list/function"
)

// CmdOptions represents command-line options for go-package-function-list
type CmdOptions struct {
	Dir        string `short:"d"`
	Regex      string `short:"x" default:"" description:"directory to parse"`
	Suffix     string `short:"s" default:"" description:"shows only functions which has given suffix if specified"`
	Delimiter  string `short:"d" default:" " description:"delimiter among function names"`
	PublicOnly bool   `short:"p" default:false description:"whether shows only public methods or not"`
	Recursive  bool   `short:"r" default:false description:"parses directory recursively if true"`
}

func main() {
	dir := flag.String("dir", "", "directory to parse")
	regex := flag.String("regex", "", "regexp to filter results")
	suffix := flag.String("suffix", "", "if specified, show only functions has given suffix")
	delimiter := flag.String("delimiter", " ", "delimiter among function names")
	publicOnly := flag.Bool("public-only", false, "whether shows only public methods or not")
	recursive := flag.Bool("recursive", false, "parses directory recursively")
	flag.Parse()

	fnames, err := function.ParseDir(*dir, *recursive)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	if *regex != "" {
		newFnames, err := function.FilterByRegexp(fnames, *regex)
		if err != nil {
			fmt.Println("Failed to parse regexp: ", err)
			os.Exit(1)
		}
		fnames = newFnames
	}
	if *suffix != "" {
		fnames = function.FilterBySuffix(fnames, *suffix)
	}
	if *publicOnly {
		fnames = function.FilterPublicMethod(fnames)
	}

	fmt.Println(strings.Join(fnames, *delimiter))
}
