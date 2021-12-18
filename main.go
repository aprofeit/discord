package main

import (
  "net/http"
  "os"
  "fmt"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://discord.gg/qKRzRVK5kz", 307)
  })

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  fmt.Println("Starting the server")
  http.ListenAndServe(":" + port, nil)
}
