package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//var templ = template.Must(template.ParseFiles("home.html"))
var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("assets/templates/*"))
}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//templ.Execute(w, nil)

//}
func index(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "index.layout", footer)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	fmt.Println("we got home")
}
func navbar(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "navbar.layout", nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	fmt.Println("we got navbar")
}
func footer(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "footer.layout", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	fmt.Println("we got footer")
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("./assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/navbar", navbar)
	http.HandleFunc("/footer", footer)

	http.ListenAndServe(":8088", nil)
}
