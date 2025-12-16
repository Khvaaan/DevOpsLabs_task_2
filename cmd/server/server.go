package main

import (
  "log"
  "net/http"
  "os"

  "go-news/pkg/api"
  "go-news/pkg/storage"
  "go-news/pkg/storage/memdb"
  "go-news/pkg/storage/postgres"
)

type server struct {
  db  storage.Interface
  api *api.API
}

func main() {
  var srv server

  storageType := os.Getenv("STORAGE")
  if storageType == "" {
      storageType = "memdb"
  }
  
  if storageType == "postgres" {
      dsn := os.Getenv("DATABASE_URL")
      if dsn == "" {
          log.Fatal("DATABASE_URL is empty")
      }
  
      db, err := postgres.New(dsn)
      if err != nil {
          log.Fatalf("postgres init failed: %v", err)
      }
  
      srv.db = db
  } else {
      srv.db = memdb.New()
  }

  srv.api = api.New(srv.db)


  log.Println("listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", srv.api.Router()))
}
