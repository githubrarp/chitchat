package main

import (
	"html/template"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {

	threads, err := data.Threads()
	if err == nil {
		_, err := session(w, r)
		public_tmp_files := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		private_tmp_files := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}
		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(public_tmp_files...))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}
	/*
		files := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html",
		}
		templates := template.Must(template.ParseFiles(files...))
		threads, err := data.Threads()
		if err == nil{
			templates.ExecuteTemplate(w, "layout", threads)
		}*/
}
