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
	errorserve(res, tmp.ExecuteTemplate(res, "index.gohtml", nil))
}

func abouthandle(res http.ResponseWriter, req *http.Request) {
	errorserve(res, tmp.ExecuteTemplate(res, "about.gohtml", nil))
}

func applyhandle(res http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		errorserve(res, tmp.ExecuteTemplate(res, "applyProcess.gohtml", nil))
	} else if req.Method == http.MethodGet {
		errorserve(res, tmp.ExecuteTemplate(res, "apply.gohtml", nil))
	}

}

func contacthandle(res http.ResponseWriter, req *http.Request) {
	errorserve(res, tmp.ExecuteTemplate(res, "contact.gohtml", nil))
}

func errorserve(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
