package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogimg)
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `Foo ran`)
}

func dog(res http.ResponseWriter, req *http.Request) {
	tmp := template.Must(template.ParseFiles("dog.gohtml"))

	tmp.Execute(res, "<h1>This is from dog</h1>")
}

func dogimg(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./res/dog.jpg")
}
