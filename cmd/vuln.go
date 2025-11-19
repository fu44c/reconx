package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var vulnCmd = &cobra.Command{
	Use:   "vuln",
	Short: "Check for vulnerabilities like SQLi, XSS, etc.",
	Long:  "Check the target domain for common vulnerabilities like SQL Injection, Cross-Site Scripting, Local File Inclusion, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		if domain == "" {
			fmt.Println("Please specify a domain using -d")
			return
		}
		// مكان تنفيذ الكود لفحص الثغرات
		fmt.Printf("[*] Checking vulnerabilities for %s...\n", domain)
	},
}

func init() {
	// لا داعي لإضافة rootCmd هنا في هذا الملف
}
