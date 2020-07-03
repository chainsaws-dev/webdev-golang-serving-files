package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./starting-files/public/"))))
	http.HandleFunc("/", displaytemp)

	http.ListenAndServe(":8080", nil)
}

func displaytemp(rew http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("./starting-files/templates/index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tmp.Execute(rew, nil)

	if err != nil {
		log.Fatalln(err)
	}
}
