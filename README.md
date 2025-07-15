# 🤖 EVA

**EVA** is a CLI tool to automate dev workflows like creating Git branches from Notion tickets and opening GitHub URLs.

## 📦 Features

- `eva branch TASK-123` — Create a Git branch from a Notion task (and move it to "Doing")
- `eva open` — Open the current GitHub branch in your browser
- `eva init` — Set up your local config interactively
- `eva update` — Upgrade to the latest version
- Configurable via `~/.eva/config.yaml`

## Tech Stack

- Go
- Cobra CLI
- Notion API
- GitHub + Make + YAML

## ⚙️ Installation

### 🧩 One-line install (recommended)

```bash
curl -sSfL https://raw.githubusercontent.com/lucmahoux/eva/main/install.sh | sh
```

### With Go

```bash
make install
```

Make sure `$HOME/go/bin` is in your `$PATH`

## 📦 Updating

```bash
eva update
```

## ⚙️ Configuration

The config file is required to use eva with Notion.

### 📥 Option 1 — Automatic (recommended)

```bash
eva init
```

This will:

- Ask for your Notion API Key and Database ID
- Create `~/.eva/config.yaml` for you

### ✍️ Option 2 — Manual

```bash
mkdir -p ~/.eva
nano ~/.eva/config.yaml
```

Paste this:

```yaml
notion_api_key: "your_notion_secret_here"
notion_database_id: "your_notion_database_id"
```

## Development

```bash
make build       # Build binary locally
make run         # Run with args like CMD="branch TASK-123"
make clean       # Clean build artifacts
make snapshot    # Build binaries locally via GoReleaser
make release     # Publish a version (requires VERSION + token)
```
