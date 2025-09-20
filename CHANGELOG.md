# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/)
and this project adheres to [Semantic Versioning](https://semver.org/).

---

## [v1.0.0] - 2025-09-19

### Added
- Initial extraction of **Entiqon CLI** as a standalone module (`github.com/entiqon/cli`).
- Provides a **developer & DevOps toolkit** to streamline workflows in the Entiqon ecosystem.
- Features include:
    - **Git & Release automation** (`gcpr`, `gce`, `gcr`, `gct`, `gsux`, `gcch`, `ddc`).
    - **Testing & Coverage** (`gotestx`, `run-tests.sh`, `open-coverage.sh`, CI integration with Codecov).
    - **Documentation helpers** (changelog regeneration, release note generation).
- Workflow support for **semantic commits, signed tags, releases, and stash safety**.
- CI/CD integration with enforced minimum 80% coverage.

### Purpose
The CLI bridges the gap between **library development** (e.g., `db`, `common`) and **release management**, providing runtime automation for Git, tests, coverage, tagging, and lifecycle management.

### Notes
- This release consolidates the CLI history from the `entiqon` monorepo.
- Future roadmap includes:
    - **Global installer** (`entiqon install cli`).
    - **Unified entrypoint** (`entiqon <command>`).
    - Migration from **Bash scripts to Go** for portability and richer UX.
    - **Plugin architecture** for extensibility.
    - Improved CLI test harness.
