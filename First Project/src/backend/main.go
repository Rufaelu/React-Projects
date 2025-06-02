package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type API struct {
	Key string `json:"key"`
}

func init()  {
	err := godotenv.Load("../../env.env") // go up from src/backend to root
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
}
func getapi() string{
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Fatal("API key not found in environment")
	}
	return apiKey
}

func main() {

	r := chi.NewRouter()

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {


		w.Write([]byte(getapi()))
	})
	r.Mount("/user", userroutes())
	r.Mount("/ai", airoutes())

	http.ListenAndServe(":10", r)

}

func userroutes() http.Handler {

	r := chi.NewRouter()

	r.Get("/hey", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hey from the user"))
	})
	return r
}

func airoutes() http.Handler {

	r := chi.NewRouter()

	r.Get("/Bitch", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hey from the Ai"))
	})

	return r

}
