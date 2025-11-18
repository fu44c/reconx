package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"reconx/internal/utils"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS module (placeholder)",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Log(utils.INFO, "DNS module loaded.")
		fmt.Println("DNS module is active. Add your logic here.")
	},
}

func init() {
	rootCmd.AddCommand(dnsCmd)
}
