package report

import (
    "encoding/json"
    "fmt"
    "os"
)

type Report struct {
    Target          string   `json:"target"`
    LiveHosts       []string `json:"live_hosts"`
    Vulnerabilities []string `json:"vulnerabilities"`
}

func SaveJSON(r Report, filename string) {
    if _, err := os.Stat("output"); os.IsNotExist(err) {
        os.Mkdir("output", 0755)
    }

    data, err := json.MarshalIndent(r, "", "  ")
    if err != nil {
        fmt.Println("Error saving report:", err)
        return
    }

    os.WriteFile(fmt.Sprintf("output/%s_report.json", r.Target), data, 0644)
    fmt.Println("[✔] Report saved to output folder.")
}
