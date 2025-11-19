package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan HTTP/HTTPS protocols",
	Long:  "Scan the HTTP/HTTPS protocols of the provided domain to find live endpoints and potential vulnerabilities.",
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			fmt.Println("Please specify a domain using -d")
			return
		}
		// مكان تنفيذ الكود لفحص HTTP/HTTPS
		fmt.Printf("[*] Scanning HTTP/HTTPS for %s...\n", domain)
	},
}

func init() {
	// لا داعي لإضافة rootCmd هنا في هذا الملف
}
