// ReconX-Go
// Lightweight modular recon framework.
// Entry point: loads banner + CLI dispatcher.

package main

import (
	"fmt"
	"os"

	"reconx/cmd"
	"reconx/internal/ascii"
	"reconx/internal/utils"
)

func main() {
	// Initialize logging system before anything else.
	utils.InitLogger()

	// Display the ASCII banner once at startup.
	ascii.ShowBanner()

	// Delegate execution to Cobra's command handler.
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
