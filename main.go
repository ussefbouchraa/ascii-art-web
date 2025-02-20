package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	A "asciiartweb/asciiart"

)

type Data struct {
	Result string
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {  
	
	if r.URL.Path != "/"{
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil{
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	
	if r.Method == http.MethodGet {
		if err := tmpl.Execute(w, Data{Result: ""}); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("inp1")
		banner := r.FormValue("Files")

		if !A.InitMap(banner){
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return 
		}

		if err := tmpl.Execute(w, Data{Result: A.Storing(input)}); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Bad Request", http.StatusBadRequest)	
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", HandleRequest)
	
	fmt.Println("Server running on port 8080")
	if err:= http.ListenAndServe(":8080", nil) ; err != nil{
		log.Fatal(err)
	}
}

