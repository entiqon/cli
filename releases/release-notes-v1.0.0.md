# Release Notes — cli v1.0.0 (2025-09-19)

This is the **first stable release** of the Entiqon CLI as a standalone module.  
It provides a consolidated toolkit for developers and operators working with the Entiqon ecosystem.

---

## ✨ Added

### Core Features
- **CLI Toolkit Extraction**
    - `cli/` folder has been split from the `entiqon` monorepo into its own repository: `github.com/entiqon/cli`.
    - Establishes a dedicated home for Entiqon’s developer and DevOps automation utilities.

- **Git & Release Automation**
    - `gcpr`: create GitHub PRs quickly.
    - `gce`: extract commits for changelogs.
    - `gcr`: generate releases with notes.
    - `gct`: manage signed tags.
    - `gsux`: safe stash utilities.
    - `gcch`: changelog regeneration.
    - `ddc`: deploy Docker containers.

- **Testing & Coverage**
    - `gotestx`: custom test runner with consolidated coverage output.
    - `run-tests.sh` and `open-coverage.sh`: helpers to streamline local and CI testing.
    - CI/CD integration with enforced **minimum 80% coverage** and Codecov reporting.

- **Documentation Helpers**
    - Scripts for changelog generation and release note preparation.
    - Support for semantic versioning and structured commit history.

### Purpose
The Entiqon CLI bridges the gap between **library development** (`db`, `common`) and **release management**, providing automation for testing, Git workflows, tagging, and lifecycle management.

---

## 🚀 Roadmap
Planned for upcoming releases:
- **Unified Entrypoint**: `entiqon <command>` to simplify usage.
- **Installer**: Global installation script (`entiqon install cli`).
- **Go-based Commands**: Migrate from Bash scripts to Go for portability and richer UX.
- **Plugin Architecture**: Extend CLI capabilities modularly.
- **CLI Test Harness**: Improve test coverage and usability of CLI itself.

---

## 📦 Availability

```bash
go install github.com/entiqon/cli@v1.0.0
