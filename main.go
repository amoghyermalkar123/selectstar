package main

import (
	"fmt"
	"net/http"
	"xerus/internal/db"
	"xerus/internal/handler"

	"github.com/joho/godotenv"
)

const PORT = "8091"

func main() {
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
