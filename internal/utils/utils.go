package utils

import (
	"os"
	"strings"
)

func FindFilesInDirectory(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
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
