package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Collect subdomains for a given domain",
	Long:  "Collect subdomains by scanning various sources and APIs to find subdomains for the target domain.",
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			fmt.Println("Please specify a domain using -d")
			return
		}
		// مكان تنفيذ الكود لجمع النطاقات
		fmt.Printf("[*] Collecting subdomains for %s...\n", domain)
	},
}

func init() {
	// لا داعي لإضافة rootCmd هنا في هذا الملف
}
