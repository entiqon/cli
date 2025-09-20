package main

import (
	"os"

	"github.com/entiqon/cli/internal/gotestx"
)

func main() {
	code := gotestx.Run(os.Args[1:], os.Stdout, os.Stderr)
	os.Exit(code)
}
