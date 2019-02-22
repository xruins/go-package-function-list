package function

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

func filterGoFiles(dirpath string, files []os.FileInfo) []string {
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
	return filePaths
}

// ParseDir parses files in given directory and returns the list of defined functions.
func ParseDir(dirpath string, recursive bool) ([]string, error) {
	var filePaths []string
	var files []os.FileInfo
	if recursive {
		err := filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				files = append(files, info)
				return nil
			})
		if err != nil {
			return nil, fmt.Errorf("failed to list directory. directory: %s, err: %s", dirpath, err)
		}
	} else {
		files, err := ioutil.ReadDir(dirpath)
		if err != nil {
			return nil, fmt.Errorf("failed to list directory. directory: %s", dirpath)
		}
		filePaths = filterGoFiles(dirpath, files)
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
