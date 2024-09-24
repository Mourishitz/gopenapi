package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func main() {
	basePath := "."

	filePaths := make(map[string][]string)

	jsonFile, err := os.Open("dirs.json")
	if err != nil {
		fmt.Printf("Error opening JSON file: %v\n", err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var dirs map[string]string
	if err := json.Unmarshal(byteValue, &dirs); err != nil {
		fmt.Printf("Error unmarshaling JSON data: %v\n", err)
		return
	}

	for key, dir := range dirs {
		if filepath.Ext(dir) == ".yml" { // Single file
			filePaths[key] = []string{dir}
		} else {
			paths, err := ScanDirectory(dir)
			if err != nil {
				fmt.Printf("Error scanning directory %s: %v\n", dir, err)
				return
			}
			filePaths[key] = paths
		}
	}

	mergedOpenAPI, err := MergeYAMLFiles(basePath, filePaths)
	if err != nil {
		fmt.Printf("Error merging YAML files: %v\n", err)
		return
	}

	file := NewFile("swagger.yaml")
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(mergedOpenAPI); err != nil {
		fmt.Printf("Error writing to swagger.yaml: %v\n", err)
		return
	}

	fmt.Println("Successfully merged YAML files into swagger.yaml")
}
