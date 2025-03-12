package main

import (
	A "asciiartweb/asciiart"
	F "fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Status string
	Result string
}

var(
	tmpl, _ = template.ParseFiles("index.html")
	tmplStatus, _ = template.ParseFiles("index1.html")
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	
	if tmplStatus == nil  {http.Error(w, "500 Internal Server Error", http.StatusInternalServerError); return
	}
	if tmpl == nil {
		w.WriteHeader(500); tmplStatus.Execute(w, Data{Status: "500 Internal Server Error"});return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(404); tmplStatus.Execute(w, Data{Status: "404 Not Found"}); return
	}

	if r.Method == http.MethodGet {
		if err := tmpl.Execute(w, Data{Result: ""}); err != nil {
			w.WriteHeader(500); tmplStatus.Execute(w, Data{Status: "500 Internal Server Error"})
		}
		return
	}

	if r.Method == http.MethodPost {
		input, banner := r.FormValue("inp1"), r.FormValue("Files")

		if len(input) > 300 || input == "" || !A.InitMap(banner) {
			w.WriteHeader(400) ;tmplStatus.Execute(w, Data{Status: "400 Bad Request"}); return
		}

		if err := tmpl.Execute(w, Data{Result: A.Storing(input)}); err != nil {
			w.WriteHeader(500); tmplStatus.Execute(w, Data{Status: "500 Internal Server Error"})
		}
		return
	}
	w.WriteHeader(405); tmplStatus.Execute(w, Data{Status: "405 Method Not Allowed"})
}


func staticFileServer(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if tmplStatus == nil  {http.Error(w, "500 Internal Server Error", http.StatusInternalServerError); return }
	
	switch url {
	case "/static/style.css", "/static/icon.png":
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	case "/static/":
		w.WriteHeader(http.StatusForbidden); tmplStatus.Execute(w, Data{Status: "403 Forbidden"})
	default:
		w.WriteHeader(http.StatusNotFound); tmplStatus.Execute(w, Data{Status: "404 Not Found"})
	}
}


func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		os.Stderr.WriteString("Err: Invalid Usage [go run .]\n")
		return
	}

	http.HandleFunc("/static/", staticFileServer)
	http.HandleFunc("/", HandleRequest)

	F.Println("Server running on port http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
