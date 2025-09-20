# Entiqon CLI

**Developer & DevOps toolkit for the Entiqon ecosystem**

The Entiqon CLI provides a collection of tools to streamline development, testing, release automation, and operational workflows.
All tools are distributed under the `entiqon/cli` module and live inside the `cmd/` directory.

---

## 📦 Tools

### ✅ Go-native binaries
- **gotestx** — Go Test eXtended tool with coverage support
  - Extends `go test` with coverage, quiet, and clean modes.
  - Supports combined flags (`-cqC`, etc.) and auto package detection.
  - Supersedes legacy Bash helpers `run-tests.sh` and `open-coverage.sh`.

### 📝 Bash-based tools (planned migration to Go)
- **gcpr** — create GitHub Pull Requests
- **gce** — extract commit history
- **gcr** — generate release notes
- **gct** — create and sign tags
- **gsux** — stash/unstash workflow utility
- **gcch** — changelog helper
- **ddc** — deploy Docker container

These remain Bash scripts for now but will gradually be ported to Go under `cmd/`.

---

## 🚀 Installation

From the root of `entiqon`:

```bash
go install ./cli/cmd/gotestx
```

or directly via GitHub (released versions):

```bash
go install github.com/entiqon/cli/cmd/gotestx@latest
```

Check installation:

```bash
gotestx -v
```

---

## 🛠 Development

All CLI tools live under `cmd/`. Shared logic is placed in `internal/`.

```text
cli/
 ├── cmd/
 │    ├── gotestx/      # Go-native binary
 │    ├── gcpr/         # planned Go migration
 │    ├── gcr/
 │    ├── gce/
 │    ├── gct/
 │    ├── gsux/
 │    ├── gcch/
 │    └── ddc/
 ├── internal/          # shared logic for CLI tools
 ├── go.mod
 └── go.sum
```

---

## 🔮 Future

For now, all CLI tools remain inside this repository. In the future, some may be split into standalone modules (e.g., `entiqon/gotestx`) if they grow beyond Entiqon-specific workflows.

This setup ensures:
- Unified release cycle for Entiqon CLI tools.
- Shared infrastructure for CI/CD.
- Simple contributor workflow.

---

## Roadmap
- **Global installer** (`entiqon install cli`) instead of per-repo `bin/`.
- **Unified entrypoint** (`entiqon <command>`) wrapping all scripts.
- **Go-based CLI rewrite** – current scripts are Bash; migration to Go planned for portability, testability, and richer UX.
- **Plugin architecture** – let projects extend CLI with their own subcommands.
- **Improved test harness** – Bats/shunit2 suites for CLI validation.


✅ In short: **Entiqon CLI = developer efficiency + project discipline**.  
It codifies the workflows we already practice (TDD, semantic commits, 100% coverage, structured releases) into **repeatable, versioned, safe automation.**

---

## 📄 License

Part of the [Entiqon Project](https://github.com/entiqon).  
Licensed under the MIT License.

