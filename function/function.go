package function

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"strings"
	"sync"
)

// ParseDir parses files in given directory and returns the list of defined functions.
func ParseDir(dirpath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, fmt.Errorf("failed to list directory. directory: %s", dirpath)
	}
	var filePaths []string
	for _, f := range files {
		fn := f.Name()
		// skip non *.go files
		if !strings.HasSuffix(fn, ".go") {
			continue
		}
		path := singleJoiningSlash(dirpath, fn)
		filePaths = append(filePaths, path)
	}

	var funcNames []string
	eg := errgroup.Group{}
	mutex := &sync.Mutex{}
	
	for _, file := range filePaths {
		file := file
		eg.Go(func() error {
			fns, err := ParseFile(file)
			if err != nil {
				return fmt.Errorf("failed to parse file. file: %s, err: %v", file, err)
			}
			mutex.Lock()
			funcNames = append(funcNames, fns...)
			mutex.Unlock()
			return nil
		})

	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return funcNames, nil
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
