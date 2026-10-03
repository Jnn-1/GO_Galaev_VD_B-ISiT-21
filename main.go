package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Добро пожаловать в трекер расходов!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Учебный проект: веб-приложение для учёта расходов на Go.")
}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "метод не разрешён", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "pong")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/ping", ping)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
