package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Println("Server started")

	http.Handle("/", http.FileServer(http.Dir("./starting-files/")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
