package main

import (
	"log"
	"fmt"
	"net/http"
	"./router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting Meal Plan")
	log.Fatal(http.ListenAndServe(":8080", r))
}
