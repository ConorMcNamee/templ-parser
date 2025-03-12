package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type FunctionMetaData struct {
	Name       string
	Parameters []string
}

func FindFilesInDirectory(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	fmt.Println(files)
	if err != nil {
		return nil, err
	}

	goFiles := []string{}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), "_templ.go") {
			filePath := dir + "/" + file.Name()
			goFiles = append(goFiles, filePath)
		}
	}

	return goFiles, nil
}

func ParseGoFileForTemplateFunctions(filePath string) ([]string, error) {
	// Open the Go file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Parse the Go file
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filePath, file, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	// Step 2.1: Extract function names that return templ.Component
	var funcNames []string
	for _, decl := range node.Decls {
		// Only interested in functions
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check if the function returns templ.Component
		if fn.Type.Results != nil {
			for _, result := range fn.Type.Results.List {
				if result.Type.(*ast.SelectorExpr).Sel.String() == "Component" {
					funcNames = append(funcNames, fn.Name.Name)
				}
			}
		}
	}

	return funcNames, nil
}
