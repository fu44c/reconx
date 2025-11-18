// Root command for ReconX-Go CLI.
// All subcommands attach to this command.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"reconx/internal/utils"
)

var rootCmd = &cobra.Command{
	Use:   "reconx",
	Short: "ReconX — Modular Recon Framework",
	Long:  "ReconX-Go: clean, extensible recon framework by Ahmed Ibrahim.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Log(utils.INFO, "No subcommand provided.")
		fmt.Println("Try: reconx --help")
	},
}

// Execute triggers Cobra's root command.
func Execute() error {
	return rootCmd.Execute()
}
