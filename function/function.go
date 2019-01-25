package function

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// PackageMap is an alias of map[string]*ast.Package
type PackageMap map[string]*ast.Package

// ParseDir parsed specified directory and returns FuncDeclList
func ParseDir(dir string) (map[string]*ast.Package, error) {
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, dir, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to parse directory. directory: %s", dir)
	}
	return packs, nil
}

// ToList returns string list of function names
func (f PackageMap) ToList() []string {
	funcs := []*ast.FuncDecl{}
	for _, pack := range f {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					funcs = append(funcs, fn)
				}
			}
		}
	}

	var funcNames []string
	for _, f := range funcs {
		funcNames = append(funcNames, f.Name.Name)
	}
	return funcNames
}

// ToString returns string representative delimited with specified delimiter
func (f PackageMap) ToString(delimiter string) string {
	l := f.ToList()
	return strings.Join(l, delimiter)
}
