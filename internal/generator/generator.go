package generator

import (
	"fmt"

	"github.com/NesoHQ/kraken/config"
)

func GenerateService(serviceName string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Step 1: Create folders
	if err := CreateFolders(serviceName, cfg.PackageName); err != nil {
		panic(err)
	}
	fmt.Println("✅ Created folders")

	// Step 2: Generate core files
	GenerateMain(serviceName, cfg)

	// Step 3: Generate entities
	for _, entity := range cfg.Entities {
		GenerateEntity(serviceName, cfg.PackageName, entity)
	}

	fmt.Println("✅ Generated service:", serviceName)
}

// GenerateMain generates the main.go from template
func GenerateMain(serviceName string, cfg *config.ServiceConfig) {
	err := GenerateFileFromTemplate(
		"templates/main.tmpl",
		fmt.Sprintf("%s/main/main.go", serviceName),
		cfg,
	)
	if err != nil {
		panic(fmt.Errorf("failed to generate main.go: %w", err))
	}
}

// GenerateEntity generates the entity structs
func GenerateEntity(serviceName, domainName string, entity config.Entity) {
	data := struct {
		PackageName string
		Entity      config.Entity
	}{
		PackageName: domainName,
		Entity:      entity,
	}

	err := GenerateFileFromTemplate(
		"templates/entity.tmpl",
		fmt.Sprintf("%s/%s/entity/%s.go", serviceName, domainName, entity.Name),
		data,
	)
	if err != nil {
		panic(fmt.Errorf("failed to generate entity: %w", err))
	}
}
