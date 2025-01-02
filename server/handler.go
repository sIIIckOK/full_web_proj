package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tp "github.com/siiickok/full-web-proj/types"
)

func (s Server) RegisterHandlers() {
    // Frontend
    s.Mux.Handle("/", http.FileServer(http.Dir("./dist")))

    //API
    s.Mux.HandleFunc("GET /api/get-item", s.HandleGetItem)
    s.Mux.HandleFunc("GET /api/get-items", s.HandleGetItems)

    s.Mux.HandleFunc("POST /api/post-item", s.HandlePostItem)

    //Testing
    s.Mux.HandleFunc("GET /api/ping", s.HandlePing)
}

func (s Server) HandleGetItem(w http.ResponseWriter, r *http.Request) {
    var id struct { Id int `json:"id"` }
    if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    item, err := s.DB.GetItem(id.Id)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    if err := writeJSON(w, item); err != nil {
        fmt.Println("[ERROR]", err)
        return
    }
    log.Println("GET GetItem", item.Title, item.Price)
}

func (s Server) HandleGetItems(w http.ResponseWriter, r *http.Request) {
    items, err := s.DB.GetItems()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    if err := writeJSON(w, items); err != nil {
        fmt.Println("[ERROR]", err)
        return
    }
}

func (s Server) HandlePostItem(w http.ResponseWriter, r *http.Request) {
    var it tp.Item 
    json.NewDecoder(r.Body).Decode(&it)
    err := s.DB.PostItem(it)
    if err != nil {
        fmt.Println("[ERROR]", err)
        return
    }
    log.Println("POST PostItem", it.Title, it.Price)
}

func (s Server) HandlePing(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Ping")
    ping := struct { 
        Ping string `json:"ping"` 
    }{ 
        Ping: "pong", 
    }
    if err := writeJSON(w, ping); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
    }
}

func writeJSON(w http.ResponseWriter, v any) error {
    w.Header().Add("Content-Type", "application/json")
    data, err := json.Marshal(v)
    if err != nil {
        return err
    }
    w.Write(data)
    return nil
}

