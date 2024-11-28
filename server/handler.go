package main

import (
	"encoding/json"
	"fmt"
	"net/http"

    tp "github.com/siiickok/full-web-proj/types"
)

func (s Server) RegisterHandlers() {
    s.Mux.HandleFunc("GET /item", s.HandleGetItem)
    s.Mux.HandleFunc("GET /items", s.HandleGetItems)

    s.Mux.HandleFunc("POST /post-item", s.PostItem)
}

func (s Server) HandleGetItem(w http.ResponseWriter, r *http.Request) {
    var id struct { Id int `json:"id"` }
    if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    item, err := s.DB.GetItem(id.Id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    if err := WriteJSON(w, item); err != nil {
        fmt.Println("[ERROR]", err)
        return
    }
}

func (s Server) HandleGetItems(w http.ResponseWriter, r *http.Request) {
    var amount struct { Amount int `json:"amount"` }
    if err := json.NewDecoder(r.Body).Decode(&amount); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    item, err := s.DB.GetItems(amount.Amount)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    if err := WriteJSON(w, item); err != nil {
        fmt.Println("[ERROR]", err)
        return
    }
}

func (s Server) PostItem(w http.ResponseWriter, r *http.Request) {
    var it tp.Item 
    json.NewDecoder(r.Body).Decode(&it)
    err := s.DB.PostItem(it)
    if err != nil {
        fmt.Println("[ERROR]", err)
        return
    }
}

func WriteJSON(w http.ResponseWriter, v any) error {
    data, err := json.Marshal(v)
    if err != nil {
        return err
    }
    w.Write(data)
    return nil
}

