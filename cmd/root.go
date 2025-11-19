package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var domain string
var auto bool
var output string

// rootCmd هو الأمر الرئيسي الذي يحتوي على الأوامر
var rootCmd = &cobra.Command{
	Use:   "reconx",
	Short: "ReconX - Advanced Bug Bounty & Recon Tool",
	Long:  "ReconX automates reconnaissance, scanning, and vulnerability detection for bug bounty hunters.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // طباعة التعليمات إذا لم يتم تحديد أي أمر
	},
}

func init() {
	// إضافة الأوامر (التي سيتم تعريفها في ملفات أخرى) إلى rootCmd هنا فقط
	rootCmd.AddCommand(collectCmd)  // تأكد من أن collectCmd تم تعريفه فقط في collect.go
	rootCmd.AddCommand(scanCmd)     // تأكد من أن scanCmd تم تعريفه فقط في scan.go
	rootCmd.AddCommand(vulnCmd)     // تأكد من أن vulnCmd تم تعريفه فقط في vuln.go

	// إضافة الأعلام (flags) للأوامر
	rootCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Target domain (required)")
	rootCmd.PersistentFlags().BoolVar(&auto, "auto", false, "Run full automated workflow")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Path to save the output report")
}

// هذا دالة لتنفيذ rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
