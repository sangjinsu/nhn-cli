package main

import (
	"fmt"
	"os"

	"nhncli/cmd"
	_ "nhncli/cmd/compute"
	_ "nhncli/cmd/vpc"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
