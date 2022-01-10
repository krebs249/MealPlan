package middleware

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"	
	"encoding/json"
	"../models"
)

func init() {
	fmt.Println("Used to init stuff")
}

func GetRecipesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Inside get recipes")

	data, err := ioutil.ReadFile("TestRecipes.json")
	if err != nil {
		log.Fatal(err)
	}

	var recipesObj models.RecipesAPI
	err = json.Unmarshal([]byte(data), &recipesObj)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(recipesObj)
}