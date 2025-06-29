package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "✅ ラムちゃんからの応答だっちゃ〜〜〜！！")
    })

    fmt.Printf("🌐 Listening on port %s...\n", port)
    http.ListenAndServe(":"+port, nil)
}
