package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/fu44c/reconx/internal/collector"
)

var domain string

var collectCmd = &cobra.Command{
    Use:   "collect",
    Short: "Collect subdomains for a target",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Collecting domains for:", domain)
        collector.Start(domain)
    },
}

func init() {
    collectCmd.Flags().StringVarP(&domain, "domain", "d", "", "Target domain")
    rootCmd.AddCommand(collectCmd)
}
