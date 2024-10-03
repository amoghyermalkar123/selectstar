package main

import (
	"fmt"
	"log"
	"net/http"
	"xerus/internal/db"
	"xerus/internal/handler"

	"github.com/joho/godotenv"
	"github.com/sevlyar/go-daemon"
)

const PORT = "8091"

func main() {
	cntxt := &daemon.Context{
		PidFileName: "selectstar.pid",
		PidFilePerm: 0644,
		LogFileName: "selectstar.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")

	runServer()
}

func runServer() {
	_ = godotenv.Load()
	mux := http.NewServeMux()

	db := db.DB()
	defer db.Close()
	h := handler.NewHandler(db)

	mux.HandleFunc("GET /static/", h.ServeStaticFiles)
	mux.HandleFunc("GET /", h.Home)
	mux.HandleFunc("GET /search", h.SearchQuery)
	mux.HandleFunc("POST /query", h.AddQuery)
	mux.HandleFunc("GET /query", h.GetQuery)
	mux.HandleFunc("DELETE /query", h.DeleteQuery)
	fmt.Printf("server is running on port %s\n", PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		fmt.Println(err)
	}
}
