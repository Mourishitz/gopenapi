package main

import (
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func MergeYAMLFiles(basePath string, filePaths map[string][]string) (OpenAPI, error) {
	var result OpenAPI

	for _, path := range filePaths["base"] {
		file := OpenFile(filepath.Join(basePath, path))
		defer file.Close()

		content := ReadFromFile(file)
		if err := yaml.Unmarshal(content, &result); err != nil {
			return result, fmt.Errorf("failed to unmarshal YAML from file %s: %w", path, err)
		}
	}

	for field, paths := range filePaths {
		if field == "base" {
			continue
		}
		for _, path := range paths {
			file := OpenFile(filepath.Join(basePath, path))
			defer file.Close()

			content := ReadFromFile(file)
			var node interface{}
			if err := yaml.Unmarshal(content, &node); err != nil {
				return result, fmt.Errorf("failed to unmarshal YAML from file %s: %w", path, err)
			}

			switch field {
			case "info":
				if result.Info == nil {
					result.Info = make(map[string]interface{})
				}
				for k, v := range node.(map[string]interface{}) {
					result.Info[k] = v
				}
			case "security":
				if security, ok := node.([]interface{}); ok {
					for _, security := range security {
						if securityMap, ok := security.(map[string]interface{}); ok {
							result.Security = append(result.Security, securityMap)
						}
					}
				}
			case "tags":
				if tags, ok := node.([]interface{}); ok {
					for _, tag := range tags {
						if tagMap, ok := tag.(map[string]interface{}); ok {
							result.Tags = append(result.Tags, tagMap)
						}
					}
				}
			case "paths":
				if result.Paths == nil {
					result.Paths = make(map[string]interface{})
				}
				for k, v := range node.(map[string]interface{}) {
					result.Paths[k] = v
				}
			case "components":
				if result.Components == nil {
					result.Components = make(map[string]interface{})
				}
				for k, v := range node.(map[string]interface{}) {
					result.Components[k] = v
				}
			case "requests":
				if result.Components == nil {
					result.Components = make(map[string]interface{})
				}
				if result.Components["requests"] == nil {
					result.Components["requests"] = make(map[string]interface{})
				}
				for k, v := range node.(map[string]interface{}) {
					result.Components["requests"].(map[string]interface{})[k] = v
				}
			case "schemas":
				if result.Components == nil {
					result.Components = make(map[string]interface{})
				}
				if result.Components["schemas"] == nil {
					result.Components["schemas"] = make(map[string]interface{})
				}
				for k, v := range node.(map[string]interface{}) {
					result.Components["schemas"].(map[string]interface{})[k] = v
				}
			case "definitions":
				if result.Definitions == nil {
					result.Definitions = make(map[string]interface{})
				}
				for k, v := range node.(map[string]interface{}) {
					result.Definitions[k] = v
				}
			case "servers":
				if servers, ok := node.([]interface{}); ok {
					for _, server := range servers {
						if serverMap, ok := server.(map[string]interface{}); ok {
							result.Servers = append(result.Servers, serverMap)
						}
					}
				}
			}
		}
	}

	return result, nil
}
