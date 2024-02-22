package main

import (
	"fmt"
	"log"
	"net/http"
	customHttp "quiz-app/internal/handler"
)

func main() {
	http.HandleFunc("/questions", customHttp.GetQuestions)
	http.HandleFunc("/submit", customHttp.SubmitAnswers)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
