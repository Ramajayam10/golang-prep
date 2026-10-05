# Golang Workspace

This workspace has been configured with Go 1.27.0 and tooling.

---

## 🛠 Installed Components

- **Go SDK**: `go1.27.0 windows/amd64` (installed in `C:\Program Files\Go`)
- **Language Server**: `gopls` (v0.23.0) for code intelligence, autocomplete, and diagnostics
- **Editor Integration**: Official Go extension (`golang.go`) configured with auto-formatting on save

---

## 🚀 Quick Start Commands

Run these commands in PowerShell or Command Prompt from this folder:

```powershell
# Run the application directly
go run main.go

# Build an executable binary (.exe)
go build -o golang-app.exe main.go

# Run unit tests
go test ./...

# Format all code files
go fmt ./...

# Add a third-party package/dependency
go get <package-name>
# Example: go get github.com/gin-gonic/gin

# Clean up unused dependencies in go.mod
go mod tidy
```

---

## 📂 Project Structure

```
Golang/
├── .vscode/
│   └── settings.json    # Editor settings (auto-format, gopls configuration)
├── .gitignore           # Git ignore rules for Go binaries and caches
├── go.mod               # Module definitions and dependencies
├── main.go              # Starter application
└── README.md            # Setup and command documentation
```
