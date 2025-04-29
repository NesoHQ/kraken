```
kraken/
├── cmd/
│   └── root.go             # Cobra root command (entry point)
│   └── init.go             # `kraken init` command logic
├── config/
│   └── sample_config.yml   # Example config file (to copy into microservice repo)
│   └── parser.go           # Code to parse config.yml into Go structs
├── internal/
│   └── generator/
│       ├── generator.go    # Core logic: create folders, files
│       ├── folder.go       # Folder creation logic
│       └── file.go         # File generation logic
├── templates/
│   ├── entity.tmpl         # Entity struct template
│   ├── repo.tmpl           # Repo interface template
│   ├── usecase.tmpl        # Usecase template
│   ├── handler.tmpl        # Handler template
│   ├── request.tmpl        # Request DTO
│   ├── response.tmpl       # Response DTO
│   ├── transformer.tmpl    # Transformer template
│   ├── main.tmpl           # main.go template
│   ├── Dockerfile.tmpl     # Dockerfile template
│   ├── docker-compose.yml.tmpl # docker-compose template
│   ├── Makefile.tmpl       # Makefile template
│   └── go.mod.tmpl         # go.mod template
├── go.mod
├── go.sum
└── main.go                 # main.go (bootstraps Cobra CLI)

```
