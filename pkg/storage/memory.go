package storage

import (
	. "quiz-app/internal/domain"
)

var Questions = []Question{
	{
		ID:   1,
		Text: "What is 2+2?",
		CorrectAnswerId: 2,
		Answers: []Answer{
			{ID: 1, Text: "3"},
			{ID: 2, Text: "4"},
			{ID: 3, Text: "5"},
		},
	},
	{
		ID:   2,
		Text: "What is 4+4?",
		CorrectAnswerId: 3,
		Answers: []Answer{
			{ID: 1, Text: "3"},
			{ID: 2, Text: "4"},
			{ID: 3, Text: "8"},
		},
	},
}

// you were better than 10% 

var Results=[]QuizResult{};
