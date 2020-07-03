package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmp *template.Template

func init() {
	var err error
	tmp, err = template.ParseFiles("./starting-files/templates/index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./starting-files/public"))))
	http.HandleFunc("/", loadmain)
	http.ListenAndServe(":8080", nil)

}

func loadmain(res http.ResponseWriter, req *http.Request) {

	err := tmp.Execute(res, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
