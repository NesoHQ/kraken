# 🐙 Kraken CLI – Local Development & Testing Guide

This guide walks you through **building, testing, and updating the Kraken CLI locally**.

---

## 🔧 Step 1: Prepare and Commit the Code

Make sure your CLI entrypoint (`main.go`) is at the root level so it can be installed directly via `go install`.

```bash
git add .
git rm -r kraken
git commit -m "<commit message>"
git push origin <your branch>

give a pr
```

---

## 🏷️ Step 2: Update Git Tag for Go Install

To make `@latest` point to your updated code:

```bash
# Create and push new tag
git tag v0.1.5
git push origin v0.1.5
```

To dlete a tag:

```bash
# Delete old tag
git tag -d v0.1.5
git push origin :refs/tags/v0.1.5
```

---

## ⚙️ Step 3: Build the CLI Locally

Build the Kraken CLI from source:

```bash
go build -o kraken .
```

> This will generate a `kraken` binary in your current directory.

---

## 🚀 Step 4: Bootstrap a New Service

Use the CLI to scaffold a new service from your config:

```bash
./kraken init test
```

> Replace `test` with your desired service name. This will read from your configured YAML and generate the full project structure.

---

## 🧪 Bonus: Clean and Re-Test

To clean up your test project:

```bash
rm -rf test
```

To rebuild and re-run:

```bash
go build -o kraken .
./kraken init another-service
```

---

## ✅ You’re Ready

Once verified, you can tag a new version (e.g., `v0.1.6`) and share it for production or public installation via:

```bash
go install github.com/NesoHQ/kraken@latest
```