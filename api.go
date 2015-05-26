package main

import (
  "net/http"
  "log"
)

func main() {
  log.Print("Serving on 3001")
  http.ListenAndServe(":3001", nil)
}
