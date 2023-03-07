package main

import (
	"log"
	"net/http"
)

func main() {
	setupAPI()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {
	manager := Newmanager()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serverWs)
}
