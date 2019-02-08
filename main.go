package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xruins/go-package-function-list/function"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s directory", os.Args[0])
	}

	dir := os.Args[1]
	fnames, err := function.ParseDir(dir)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	fmt.Println(strings.Join(fnames, ","))
}
