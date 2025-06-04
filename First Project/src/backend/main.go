package main

import (
	// "fmt"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

type API struct {
	Key string `json:"key"`
}

func init() {
	err := godotenv.Load("../../env.env") // go up from src/backend to root
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}
func getapi() string {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Fatal("API key not found in environment")
	}
	return apiKey
}

func main() {

	// r := chi.NewRouter()

	// r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
	// 	res, err := callOpenRouter("whats your name?")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	w.Write([]byte(fmt.Sprintf("apikey: %s  response: %s ", getapi(), res)))
	// })
	// r.Mount("/user", userroutes())
	// r.Mount("/ai", airoutes())

	// http.ListenAndServe(":10", r)
	gemi()

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

func gemi(){

	ctx := context.Background()
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey:  getapi(),
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.0-flash",
        genai.Text("Explain how AI works in a few words"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result.Text())
}

func callOpenRouter(userMessage string) (string, error) {
	apiKey := getapi()

	url := "https://api.openrouter.ai/v1/chat/completions"

	body := map[string]interface{}{
		"model": "openai/gpt-4o",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": userMessage,
			},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "http://localhost")
	req.Header.Set("X-Title", "ChiReactBot") // name it whatever

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}
