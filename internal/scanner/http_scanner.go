package scanner

import (
    "fmt"
    "net/http"
    "time"
)

func checkURL(url string, ch chan<- string) {
    client := http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(url)
    if err == nil && resp.StatusCode == 200 {
        ch <- fmt.Sprintf("[LIVE] %s", url)
    } else {
        ch <- fmt.Sprintf("[DEAD] %s", url)
    }
}

func Start(urls []string) {
    ch := make(chan string)
    fmt.Println("[*] Starting HTTP scan...")

    for _, u := range urls {
        go checkURL(u, ch)
    }

    for i := 0; i < len(urls); i++ {
        fmt.Println(<-ch)
    }

    fmt.Println("[✔] HTTP scan complete.")
}
