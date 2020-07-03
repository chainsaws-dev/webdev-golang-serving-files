package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmp *template.Template

func init() {
	var err error
	tmp, err = template.ParseGlob("./starting-files/templates/*.gohtml")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", mainhandle)
	http.HandleFunc("/about", abouthandle)
	http.HandleFunc("/apply", applyhandle)
	http.HandleFunc("/contact", contacthandle)

	http.ListenAndServe(":8080", nil)
}

func mainhandle(res http.ResponseWriter, req *http.Request) {
	err := tmp.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func abouthandle(res http.ResponseWriter, req *http.Request) {
	err := tmp.ExecuteTemplate(res, "about.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func applyhandle(res http.ResponseWriter, req *http.Request) {
	var err error
	if req.Method == http.MethodPost {
		err = tmp.ExecuteTemplate(res, "applyProcess.gohtml", nil)
	} else if req.Method == http.MethodGet {
		err = tmp.ExecuteTemplate(res, "apply.gohtml", nil)
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func contacthandle(res http.ResponseWriter, req *http.Request) {
	err := tmp.ExecuteTemplate(res, "contact.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
