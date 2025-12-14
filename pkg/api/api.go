package api

import (
  "encoding/json"
  "go-news/pkg/storage"
  "net/http"

  "github.com/gorilla/mux"
)

type API struct {
  db     storage.Interface
  router *mux.Router
}

func New(db storage.Interface) *API {
  api := API{db: db}
  api.router = mux.NewRouter()
  api.endpoints()
  return &api
}

func (api *API) endpoints() {
  api.router.HandleFunc("/tasks", api.tasksHandler).Methods(http.MethodGet, http.MethodOptions)
  api.router.HandleFunc("/tasks", api.addTaskHandler).Methods(http.MethodPost, http.MethodOptions)
  api.router.HandleFunc("/tasks", api.updateTaskHandler).Methods(http.MethodPut, http.MethodOptions)
  api.router.HandleFunc("/tasks", api.deleteTaskHandler).Methods(http.MethodDelete, http.MethodOptions)
}

func (api *API) Router() *mux.Router {
  return api.router
}

func (api *API) tasksHandler(w http.ResponseWriter, r *http.Request) {
  tasks, err := api.db.Tasks()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  bytes, err := json.Marshal(tasks)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Write(bytes)
}

func (api *API) addTaskHandler(w http.ResponseWriter, r *http.Request) {
  var t storage.Task
  if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  if err := api.db.AddTask(t); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
}

func (api *API) updateTaskHandler(w http.ResponseWriter, r *http.Request) {
  var t storage.Task
  if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  if err := api.db.UpdateTask(t); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
}

func (api *API) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
  var t storage.Task
  if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  if err := api.db.DeleteTask(t); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
}
