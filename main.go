package main

import (
  "fmt"
  "net/http"
  "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
  out, err := exec.Command("sh","-c","/path/to/my/script").Output()
  if err != nil {
    http.Error(w, err.Error(), 500)
  }
  fmt.Fprintf(w, "%s", out)
}

func main() {
  http.HandleFunc("/secret_hash", handler)
  http.ListenAndServe("172.17.0.1:8080", nil)
}
