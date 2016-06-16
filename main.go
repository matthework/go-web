package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	router := NewRouter()
	port := os.Getenv("PORT")
	fmt.Println("Listening on " + port + "...")

	log.Fatal(http.ListenAndServe(":" + port, router))
}
