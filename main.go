package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB()

	log.Println("Listening...")

	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
}
