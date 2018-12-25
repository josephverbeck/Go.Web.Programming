package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("temp.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func list(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("listTemp.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thur", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func settingDot(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("setDot.html")
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/days", list)
	http.HandleFunc("/set_dot", settingDot)
	server.ListenAndServe()
}
