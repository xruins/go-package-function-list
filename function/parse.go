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

// filterGoFiles filters out non go-source file from given string slice.
func filterGoFiles(paths []string) []string {
	var filePaths []string
	for _, p := range paths {
		// skip non *.go files
		if !strings.HasSuffix(p, ".go") {
			continue
		}
		filePaths = append(filePaths, p)
	}
	return filePaths
}

// ParseDir parses files in given directory and returns the list of defined functions.
func ParseDir(dirpath string, recursive bool) ([]string, error) {
	var filePaths []string
	var files []os.FileInfo
	var err error
	if recursive {
		err = filepath.Walk(dirpath,
			func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				// skip hidden files
				if strings.HasPrefix(".", info.Name()) {
					return nil
				}
				if err != nil {
					return err
				}
				filePaths = append(filePaths, path)
				return nil
			})
		if err != nil {
			return nil, fmt.Errorf("failed to list directory. directory: %s, err: %s", dirpath, err)
		}
	} else {
		files, err = ioutil.ReadDir(dirpath)
		if err != nil {
			return nil, fmt.Errorf("failed to list directory. directory: %s, err: %s", dirpath, err)
		}

		for _, f := range files {
			filePaths = append(filePaths, singleJoiningSlash(dirpath, f.Name()))
		}
	}

	var funcNames []string
	eg := errgroup.Group{}
	mutex := &sync.Mutex{}

	for _, file := range filterGoFiles(filePaths) {
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
