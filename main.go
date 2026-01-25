package main

import (
	"os"

	"nhncli/cmd"
	_ "nhncli/cmd/vpc"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
