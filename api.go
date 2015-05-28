package main

import (
  "net/http"
  "log"
)

type Todo struct {
  id int64
  task string
  completed bool
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  log.Print("Went to " + r.URL.Path[len("/"):])
  http.ServeFile(w, r, "./index.html")
}

func main() {
  log.Print("Serving on 3001")
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":3001", nil)
}
