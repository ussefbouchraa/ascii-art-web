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
	
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/"{
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	
	if r.Method == http.MethodGet {
		tmpl.Execute(w, Data{Result: ""})
		return
	}

	if r.Method == http.MethodPost {
		
		input := r.FormValue("inp1")
		banner := r.FormValue("Files")

		if !A.InitMap(banner){
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return 
		}

		tmpl.Execute(w, Data{Result: A.Storing(input)})
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", HandleRequest)
	
	fmt.Println("Server running on port 8080")
	if err:= http.ListenAndServe("127.0.0.1:8080", nil) ; err != nil{
		log.Fatal(err)
	}
}

