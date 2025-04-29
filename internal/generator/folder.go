package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

var folders = []string{
	"apm",
	"cache",
	"cmd",
	"config",
	"grpc",
	"healthcheck",
	"logger",
	"main",
	"metrics",
	"notification",
	"provider",
	"repo",
	"rest",
	"tmp",
	"types",
	"util",
}

func CreateFolders(basePath, domainName string) error {
	for _, folder := range folders {
		path := filepath.Join(basePath, folder)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create folder %s: %w", folder, err)
		}
	}

	// Create domain-driven folders
	domainFolders := []string{"entity", "handler", "request", "response", "service", "transformer", "usecase"}
	for _, folder := range domainFolders {
		path := filepath.Join(basePath, domainName, folder)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create domain folder %s: %w", folder, err)
		}
	}

	return nil
}
