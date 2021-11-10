package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Loading the env file which contains your unique API_KEY
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}
	api_key := os.Getenv("API_KEY")

	// Making a request to the API
	response, err := http.Get(api_key)
	if err != nil {
		log.Println(err)
	}
	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}