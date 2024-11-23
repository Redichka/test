package main

import (
	"html/template"
	"log"
	"net/http"
)

var (
	username = "user"
	password = "123"
)

type WelcomePage struct {
	Username string
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.Handle("/style.css", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		inputUsername := r.FormValue("username")
		inputPassword := r.FormValue("password")

		if inputUsername == username && inputPassword == password {
			http.Redirect(w, r, "/welcome?username="+inputUsername, http.StatusSeeOther)
			return
		} else {
			http.Error(w, "Неправильный логин или пароль", http.StatusUnauthorized)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles("./login.html"))
	tmpl.Execute(w, nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles("./welcome.html"))
	tmpl.Execute(w, WelcomePage{Username: username})
}
