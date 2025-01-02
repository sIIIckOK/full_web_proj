package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	db "github.com/siiickok/full-web-proj/db"
	tp "github.com/siiickok/full-web-proj/types"
)

const (
	port = ":6969"
)

type Server struct {
	Mux  *http.ServeMux
	DB   tp.DB
	Addr string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("[PANIC]", err)
	}
	var (
		dbAddr     = getEnvOrDefault("DB_ADDRESS", "5432")
		hostName   = getEnvOrDefault("DB_HOSTNAME", "localhost")
		userName   = getEnvOrDefault("USERNAME", "postgres")
		dbName     = getEnvOrDefault("DB_NAME", "postgres")
		dbPassword = getEnvOrDefault("DB_PASSWORD", "postgres")
	)
	dbStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		hostName,
		userName,
		dbName,
		dbPassword,
		dbAddr)
	sv, err := NewServer(port, dbStr)
	if err != nil {
		log.Fatalln("[PANIC]", err)
	}

	sv.RegisterHandlers()
	if err != nil {
		log.Fatalln("[PANIC]", err)
	}
	if err := sv.Run(); err != nil {
		log.Fatalln("[PANIC]", err)
	}
}

func NewServer(addr string, dbStr string) (Server, error) {
	mux := http.NewServeMux()
	pq, err := db.PostgresOpen(dbStr)
	if err != nil {
		return Server{}, err
	}
	if err := pq.Ping(); err != nil {
		return Server{}, err
	}
	log.Println("[INFO]", "Successfully connected to database")
	if err := pq.Init(); err != nil {
		return Server{}, err
	}
	log.Println("[INFO]", "Successfully initialized the database")
	return Server{
		Addr: addr,
		Mux:  mux,
		DB:   pq,
	}, nil
}

func (s Server) Run() error {
	log.Printf("[INFO] Link: http://localhost%s\n", port)
	log.Printf("[INFO] Starting Server at port `%s`\n", s.Addr)
	return http.ListenAndServe(s.Addr, s.Mux)
}

func getEnvOrDefault(key string, default_str string) string {
	res := os.Getenv(key)
	if res == "" {
		return default_str
	}
	return res
}
