package main

import (
	"./db"
	"fmt"
)



func main() {
	//r := router.Router()
	fmt.Println("Starting Meal Plan")
	//log.Fatal(http.ListenAndServe(":8080", r))

	db.Connect()


	var email string
	var password string
	fmt.Println("Please enter email: ")
	fmt.Scanln(&email)
	fmt.Println("Please enter password: ")
	fmt.Scanln(&password)

	db.Create_User(email, password)

}
