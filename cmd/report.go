package cmd

import (
	"fmt"
	"reconx/internal/output"
	"reconx/internal/utils"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate sample report",
	Run: func(cmd *cobra.Command, args []string) {

		sample := map[string]string{
			"Target": "example.com",
			"Status": "Completed",
			"Notes":  "Report system working successfully.",
		}

		err := output.GenerateHTMLReport("report.html", "Sample Report", sample)
		if err != nil {
			utils.Log(utils.ERROR, "Failed to generate HTML report")
		}

		output.GenerateJSONReport("report.json", sample)
		output.GenerateTextReport("report.txt", sample)

		utils.Log(utils.SUCCESS, "Reports generated successfully!")
		fmt.Println("Files: report.html, report.json, report.txt")
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
