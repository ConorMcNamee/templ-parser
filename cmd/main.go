package main

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"templparser/internal/utils"

	"github.com/a-h/templ"
	templruntime "github.com/a-h/templ/runtime"
)

func main() {

	files, err := utils.FindFilesInDirectory("./")
	if err != nil {
		fmt.Errorf("Error finding templ files", err)
	}
	// Step 2: Parse the Go files to find the template functions
	for _, file := range files {
		funcNames, err := parseGoFileForTemplateFunctions(file)
		if err != nil {
			fmt.Errorf("Error parsing templ files", err)
		}
		for _, funcName := range funcNames {
			// Use reflection to dynamically call the function
			component, err := getTemplateComponentByName(funcName)
			if err != nil {
				log.Printf("Error getting template function %s: %v", funcName, err)
				continue
			}

			// Create context and writer (output buffer)
			ctx := context.Background()
			var outputBuffer bytes.Buffer

			// Step 4: Render the component
			err = component.Render(ctx, &outputBuffer)
			if err != nil {
				log.Printf("Error rendering template %s: %v", funcName, err)
				continue
			}

			// Print the output (this should contain the rendered HTML)
			fmt.Printf("Rendered output from %s:\n", funcName)
			fmt.Println(outputBuffer.String())
		}
	}
}

// Step 2: Parse the Go file to extract template functions
func parseGoFileForTemplateFunctions(filePath string) ([]string, error) {
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
				if result.Type.(*ast.Ident).Name == "Component" {
					funcNames = append(funcNames, fn.Name.Name)
				}
			}
		}
	}

	return funcNames, nil
}

// Step 3: Retrieve the template component by function name using reflection
func getTemplateComponentByName(funcName string) (templ.Component, error) {
	// Define the functions to be dynamically called
	funcMap := map[string]interface{}{
		"Function":  Function,
		"Function2": Function2,
		// Add all other template functions here
	}

	// Use reflection to get the function from the map
	fn, exists := funcMap[funcName]
	if !exists {
		return nil, fmt.Errorf("function %s not found", funcName)
	}

	// Call the function using reflection
	return fn.(func(string) templ.Component)("John"), nil
}

// Example template functions
func Function(name string) templ.Component {
	return templruntime.GeneratedTemplate(func(input templruntime.GeneratedComponentInput) error {
		writer := input.Writer
		_, err := writer.Write([]byte("<div>Hello " + name + "</div>"))
		return err
	})
}

func Function2(name string) templ.Component {
	return templruntime.GeneratedTemplate(func(input templruntime.GeneratedComponentInput) error {
		writer := input.Writer
		_, err := writer.Write([]byte("<div>Goodbye " + name + "</div>"))
		return err
	})
}
