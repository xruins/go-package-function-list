package function

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"strings"
)

// ParseDir parses files in given directory and returns the list of defined functions.
func ParseDir(dirpath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, fmt.Errorf("failed to list directory. directory: %s", dirpath)
	}
	filePaths := make([]string, len(files))
	for _, f := range files {
		path := singleJoiningSlash(dirpath, f.Name())
		filePaths = append(filePaths, path)
	}

	var funcNames [][]string
	eg := errgroup.Group{}
	for i, file := range filePaths {
		i, file := i, file
		eg.Go(func() error {
			fns, err := ParseFile(file)
			if err != nil {
				return fmt.Errorf("failed to parse file. file: %s, err: %v", file, err)
			}
			funcNames[i] = fns
			return nil
		})

	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return flat(funcNames), nil
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// flat flats slice of string slice to single string slice.
func flat(slice [][]string) []string {
	var entireLen int
	for _, s := range slice {
		entireLen += len(s)
	}

	ret := make([]string, entireLen)
	for _, s := range slice {
		ret = append(ret, s...)
	}
	return ret
}

// ParseFile parses given file and returns the list of defined functions.
func ParseFile(filepath string) ([]string, error) {
	set := token.NewFileSet()
	astFile, err := parser.ParseFile(set, filepath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file. path: %s, err: %v", filepath, err)
	}
	var funcNames []string
	for _, decl := range astFile.Decls {
		if fn, isFn := decl.(*ast.FuncDecl); isFn {
			funcName := fn.Name.Name
			funcNames = append(funcNames, funcName)
		}
	}
	return funcNames, nil
}
