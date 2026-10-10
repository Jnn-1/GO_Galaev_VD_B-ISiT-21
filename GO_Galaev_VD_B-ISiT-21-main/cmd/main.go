package main

import (
	"log"
	"net/http"

	"project/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/ping", handlers.Ping)

	mux.HandleFunc("GET /expenses", handlers.ExpenseList)
	mux.HandleFunc("GET /expenses/add", handlers.ExpenseForm)
	mux.HandleFunc("POST /expenses/add", handlers.ExpenseAdd)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
