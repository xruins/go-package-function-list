package main

import (
	"fmt"
	"os"
	"strings"
	"flag"

	"github.com/xruins/go-package-function-list/function"
)

func main() {
	dir := flag.String("dir", "", "directory to parse")
	regex := flag.String("regex", "", "regexp to filter results")
	suffix := flag.String("suffix", "", "if specified, show only functions has given suffix")
	delimiter := flag.String("delimiter", " ", "delimiter among function names")
	publicOnly := flag.Bool("public-only", false, "whether shows only public methods or not")
	flag.Parse()
	
	fnames, err := function.ParseDir(*dir)
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
