package main

import (
	"log"
	"net/http"

	db "github.com/siiickok/full-web-proj/db"
	tp "github.com/siiickok/full-web-proj/types"

	_ "github.com/lib/pq"
)

const(
    port = ":5050"
    dbStr = "host=localhost user=postgres dbname=postgres password=postgres port=5432 sslmode=disable"
)

func main() { 
    sv, err := NewServer(port, dbStr)
    sv.RegisterHandlers()

    if err != nil {
        log.Fatalln("[PANIC]", err)
    }
    if err := sv.Run(); err != nil {
        log.Fatalln("[PANIC]", err)
    }
}

type Server struct {
    DB   tp.DB
    Mux  *http.ServeMux
    Addr string
}

func NewServer(addr string, dbStr string) (Server, error) {
    mux := http.NewServeMux()
    pq, err := db.PostgresOpen(dbStr); if err != nil { return Server{}, err }
    if err := pq.Ping(); err != nil { return Server{}, err }
    log.Println("[INFO]", "Successfully connected to database")
    if err := pq.Init(); err != nil { return Server{}, err }
    log.Println("[INFO]", "Successfully initialized the database")
    return Server {
        Addr: addr,
        Mux: mux,
        DB: pq,
    }, nil
}

func (s Server) Run() error {
    log.Printf("[INFO] Starting Server at port `%s`\n", s.Addr)
    return http.ListenAndServe(s.Addr, s.Mux)
}

