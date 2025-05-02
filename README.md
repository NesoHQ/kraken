````markdown
# Kraken 🐙 ALPHA - 0.0.1

**Kraken** is a domain-driven, microservice generator for Go projects.

It bootstraps fully structured microservice templates following **Domain Driven Design (DDD)** principles — ready to code, scalable, and production-friendly.

> 🚀 Build microservices faster.  
> 📚 Stay structured with clean architecture.  
> ⚡ Focus on writing business logic, not boilerplate.

---

## 📦 Features

- 📚 Domain-Driven Design (DDD) architecture
- 🗂️ Automatic folder structure generation
- 🛠️ Entity, Repository, Usecase, Handler, Request/Response, Transformer code scaffolding
- ⚡ Simple YAML config-driven
- 🚀 Supports Go Modules
- 🐳 Docker and Docker-Compose files generated
- 🏗️ Production-grade project structure

---

## 🔥 Installation

Make sure you have **Go 1.18+** installed.

```bash
go install github.com/NesoHQ/kraken/cmd/kraken@latest
```
````

This will install the `kraken` CLI globally.

---

## ⚡ Usage

First, create a basic config:

```bash
mkdir my-service
cd my-service
mkdir config
nano config/sample_config.yml
```

Example `sample_config.yml`:

```yaml
service_name: my-service
package_name: myservice
database:
    type: postgres
    schema: public

entities:
    - name: Example
      table_name: examples
      fields:
          - name: id
            type: uuid
            primary: true
          - name: name
            type: string
          - name: description
            type: text
```

Then run:

```bash
kraken init my-service
```

✅ Kraken will generate the full microservice directory and code structure based on the config.

---

## 🏗️ Folder Structure Generated

```
.
├── apm
├── cache
├── cmd
├── config
├── docs
├── healthcheck
├── <your-service>
├── logger
├── metrics
├── repo
├── rest
│   ├── handlers
│   ├── middlewares
│   ├── swagger
│   └── utils
├── structure.txt
├── types
└── util

18 directories, 1 file
```

---

## 🛠️ Roadmap

- [x] Basic `init` command (generate microservice structure)
- [ ] `generate entity` command (add new entities to existing service)
- [ ] Auto-migration SQL generation
- [ ] gRPC and REST API scaffolding
- [ ] GitHub Action template for CI/CD setup

---

## 🤝 Contributing

Pull requests are welcome!  
Feel free to open issues or suggest features.

Steps:

- Fork the repo
- Create a new branch
- Submit a pull request

---

## 📄 License

[MIT](LICENSE)

---

## 🌎 Connect

Built with ❤️ by [NesoHQ](https://github.com/NesoHQ) for the future of open source developers.


