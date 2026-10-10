package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"project/internal/models"
)

var (
	expenses []models.Expense
	mu       sync.Mutex
)

type PageData struct {
	Title    string
	Expenses []models.Expense
	Today    string
}

func render(w http.ResponseWriter, page string, data PageData) {
	tmpl, err := template.ParseFiles("web/templates/layout.html", "web/templates/"+page)
	if err != nil {
		log.Println(err)
		http.Error(w, "ошибка шаблона", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	render(w, "home.html", PageData{Title: "Главная"})
}

func About(w http.ResponseWriter, r *http.Request) {
	render(w, "about.html", PageData{Title: "О проекте"})
}

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "метод не разрешён", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "pong")
}

func ExpenseList(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	list := expenses
	mu.Unlock()

	render(w, "expenses.html", PageData{Title: "Траты", Expenses: list})
}

func ExpenseForm(w http.ResponseWriter, r *http.Request) {
	today := time.Now().Format("2006-01-02")
	render(w, "add.html", PageData{Title: "Добавить трату", Today: today})
}

func ExpenseAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "не удалось разобрать форму", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil || amount <= 0 {
		http.Error(w, "сумма должна быть положительным числом", http.StatusBadRequest)
		return
	}

	description := strings.TrimSpace(r.FormValue("description"))
	if description == "" {
		http.Error(w, "описание не может быть пустым", http.StatusBadRequest)
		return
	}

	date := r.FormValue("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	mu.Lock()
	expenses = append(expenses, models.Expense{
		Amount:      amount,
		Description: description,
		Date:        date,
	})
	mu.Unlock()

	http.Redirect(w, r, "/expenses", http.StatusSeeOther)
}
