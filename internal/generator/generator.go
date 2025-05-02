package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateProject(serviceName string) {
	dirs := []string{
		".",
		"apm",
		"cache",
		"cmd",
		"config",
		"docs",
		"healthcheck",
		"logger",
		"metrics",
		"repo",
		"rest/handlers",
		"rest/middlewares",
		"rest/swagger",
		"rest/utils",
		"types",
		"util",
		serviceName,
	}

	for _, dir := range dirs {
		os.MkdirAll(filepath.Join(serviceName, dir), os.ModePerm)
	}

	files := []struct {
		Template string
		Output   string
	}{
		{"Dockerfile.tmpl", "Dockerfile"},
		{"docker-compose.yml.tmpl", "docker-compose.yml"},
		{"go.mod.tmpl", "go.mod"},
		{"main.tmpl", "main.go"},
		{"Makefile.tmpl", "Makefile"},
		{"air.toml.tmpl", ".air.toml"},
		{"env.tmpl", ".env"},
		{"env.tmpl", ".env.local"},
		{"env.tmpl", ".env.example"},
		{"gitignore.tmpl", ".gitignore"},
		// Add more template mappings as needed
	}

	for _, f := range files {
		renderTemplate(filepath.Join("templates", f.Template), filepath.Join(serviceName, f.Output), serviceName)
	}

	fmt.Printf("✅ Project '%s' initialized successfully.\n", serviceName)
}
