package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s directory", os.Args[0])
	}

	dir := os.Args[1]
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, dir, nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	funcs := []*ast.FuncDecl{}
	for _, pack := range packs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					funcs = append(funcs, fn)
				}
			}
		}
	}

	var functionNames []string
	for _, f := range funcs {
		functionNames = append(functionNames, f.Name.Name)
	}

	fmt.Print(strings.Join(functionNames, ","))
}
