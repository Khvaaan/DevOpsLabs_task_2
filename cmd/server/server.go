package main

import (
  "log"
  "net/http"

  "go-news/pkg/api"
  "go-news/pkg/storage"
  "go-news/pkg/storage/memdb"
)

type server struct {
  db  storage.Interface
  api *api.API
}

func main() {
  var srv server

  srv.db = memdb.New()
  srv.api = api.New(srv.db)

  log.Println("listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", srv.api.Router()))
}
