package service

import (
	"quiz-app/internal/domain"
	"quiz-app/pkg/storage"
)

func GetQuestions() []domain.Question {
	return storage.Questions
}

func EvaluateAnswers(userAnswers []domain.UserAnswer) domain.QuizResult {
	correctCount := 0
	questionMap := make(map[int]domain.Question)

	for _, q := range storage.Questions {
		questionMap[q.ID] = q
	}

	for _, ua := range userAnswers {
		if question, ok := questionMap[ua.QuestionID]; ok {
			if ua.AnswerID == question.CorrectAnswerId {
				correctCount++
			}
		}
	}

	result := domain.QuizResult{
		CorrectAnswers: correctCount,
		TotalQuestions: len(storage.Questions),
	}

	score := float32(correctCount) / float32(len(storage.Questions))

	numberBelowOurScore := 0
	for _, r := range storage.Results {
		if float32(r.CorrectAnswers)/float32(r.TotalQuestions) < score {
			numberBelowOurScore++
		}
	}

	if len(storage.Results) > 0 {
		result.BetterThanOthers = int(float32(numberBelowOurScore) / float32(len(storage.Results)) * 100)
	}

	storage.Results = append(storage.Results, result)
	return result
}
