package main

import (
	"html/template"
	"net/http"
)

//Not complete.  Work book requires download fro github

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"template/layout.html",
		"template/navbar.html",
		"template/index.html"}
	template := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		template.ExecuteTemplate(w, "layout", threads)
	}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
