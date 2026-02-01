// Copyright (c) 2026 Jinsu. All rights reserved.
// Use of this source code is governed by a PolyForm Noncommercial License 1.0.0
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"nhncli/cmd"
	_ "nhncli/cmd/blockstorage"
	_ "nhncli/cmd/compute"
	_ "nhncli/cmd/loadbalancer"
	_ "nhncli/cmd/objectstorage"
	_ "nhncli/cmd/vpc"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
