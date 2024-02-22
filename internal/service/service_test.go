package service

import (
	"quiz-app/internal/domain"
	"quiz-app/pkg/storage"
	"testing"
)

func setupMockData() {
    // Reset storage for isolation
    storage.Questions = []domain.Question{
        {ID: 1, Text: "What is 2+2?", CorrectAnswerId: 1, Answers: []domain.Answer{
            {ID: 1, Text: "4"},
            {ID: 2, Text: "22"},
        }},
        {ID: 2, Text: "What is the capital of France?", CorrectAnswerId: 3, Answers: []domain.Answer{
            {ID: 3, Text: "Paris"},
            {ID: 4, Text: "London"},
        }},
    }

    storage.Results = []domain.QuizResult{} // Ensure results are empty
}

func TestEvaluateAnswers(t *testing.T) {
    setupMockData()

    tests := []struct {
        name     string
        answers  []domain.UserAnswer
        expected domain.QuizResult
    }{
        {
            name: "All Correct Answers",
            answers: []domain.UserAnswer{
                {QuestionID: 1, AnswerID: 1},
                {QuestionID: 2, AnswerID: 3},
            },
            expected: domain.QuizResult{
                CorrectAnswers: 2,
                TotalQuestions: 2,
                BetterThanOthers: 0, // Since there are no prior results
            },
        },
        {
            name: "No Correct Answers",
            answers: []domain.UserAnswer{
                {QuestionID: 1, AnswerID: 2},
                {QuestionID: 2, AnswerID: 4},
            },
            expected: domain.QuizResult{
                CorrectAnswers: 0,
                TotalQuestions: 2,
                BetterThanOthers: 0, // Depending on prior results, adjust as necessary
            },
        },
        {
            name: "Some Correct Answers",
            answers: []domain.UserAnswer{
                {QuestionID: 1, AnswerID: 1},
                {QuestionID: 2, AnswerID: 4},
            },
            expected: domain.QuizResult{
                CorrectAnswers: 1,
                TotalQuestions: 2,
                BetterThanOthers: 0, // Adjust based on mock data and prior tests
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := EvaluateAnswers(tt.answers)
            if result.CorrectAnswers != tt.expected.CorrectAnswers || result.TotalQuestions != tt.expected.TotalQuestions {
                t.Errorf("Expected %v, got %v", tt.expected, result)
            }
            // This is a simplistic check. You might want to check BetterThanOthers with more nuanced logic.
        })
    }
}
