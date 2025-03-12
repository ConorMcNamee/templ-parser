package app

import (
	"fmt"
	"templparser/internal/utils"
)

func ParseTemplFiles() {
	files, err := utils.FindFilesInDirectory("./internal/views")
	if err != nil {
		fmt.Errorf("Error finding templ files", err)
	}
	fmt.Println(len(files))

	if len(files) > 0 {
		for _, file := range files {
			funcNames, err := utils.ParseGoFileForTemplateFunctions(file)
			if err != nil {
				fmt.Errorf("Error parsing templ files", err)
			}
			for _, funcName := range funcNames {
				fmt.Println(funcName)
			}
		}
	} else {
		fmt.Println("No files found in directory", files)
		return
	}
}
