package domain 

// Question structure
type Question struct {
	ID             int   `json:"id"`
	Text           string   `json:"text"`
	Answers        []Answer `json:"answers"`
	CorrectAnswerId int      `json:"correctAnswerId"`
}

// Answer structure
type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// UserAnswer structure
type UserAnswer struct {
	QuestionID int `json:"questionId"`
	AnswerID  int `json:"answerId"`
}

// QuizResults to return to the user
type QuizResult struct {
	CorrectAnswers int `json:"correctAnswers"`
	TotalQuestions int `json:"totalQuestions"`
	BetterThanOthers int `json:"betterThanOthers"`
}
