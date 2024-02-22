package handler 

import (
	"encoding/json"
	"log"
	"net/http"
	"quiz-app/internal/domain"
	"quiz-app/internal/service"
)

// GetQuestions handler
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetQuestions())
}

// SubmitAnswers handler
func SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var userAnswers []domain.UserAnswer
	err := json.NewDecoder(r.Body).Decode(&userAnswers)
	if err != nil {
		log.Println("Error decoding user answers:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Evaluate answers - this is simplified for brevity
	result := service.EvaluateAnswers(userAnswers)

	log.Println("Quiz result:", result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}
