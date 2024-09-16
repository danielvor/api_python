package main

import (
  "fmt"
  "log"
  "net/http"
)

func helloGo(w http.ResponseWriter, r *http.Request) {
  _, err := fmt.Fprintf(w, "Hello Go")
  if err != nil {
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }
}

func main() {
  http.HandleFunc("/go", helloGo)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}