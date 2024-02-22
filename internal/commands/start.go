/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quiz-app/internal/domain"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("http://localhost:8080/questions")

		if err != nil {
			fmt.Println("Error fetching questions:", err)
			return
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		var questions []domain.Question
		err = json.Unmarshal(body, &questions)
		if err != nil {
			fmt.Errorf("error unmarshaling response: %v", err)
		}

		var userAnswers = []domain.UserAnswer{}

		for i, question := range questions {
			fmt.Printf("%d) %s\n", i+1, question.Text)
			for j, answer := range question.Answers {
				fmt.Printf("%d) %s \n", j+1, answer.Text)
			}
			fmt.Println("Your answer is?")

			var userAnswer int
			_, err := fmt.Scan(&userAnswer)
			if err != nil {
				fmt.Println("Failed to read the answer")
			}

			userAnswers = append(userAnswers, domain.UserAnswer{
				QuestionID: question.ID,
				AnswerID:   userAnswer,
			})
		}

		SubmitResult(userAnswers)

	},
}

func SubmitResult(userAnswers []domain.UserAnswer) {

	userAnswersJSON, err := json.Marshal(userAnswers)
	if err != nil {
		fmt.Println("Error marshaling user answers:", err)
		return
	}

	response, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(userAnswersJSON))
	if err != nil {
		fmt.Println("Error fetching questions:", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var result domain.QuizResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Errorf("error unmarshaling response: %v", err)
	}

	fmt.Printf("Correct Answers: %d of Total Questions: %d\n", result.CorrectAnswers, result.TotalQuestions)
	fmt.Printf("You did better than %d%% of participants😊 \n", result.BetterThanOthers)

}

func init() {
	rootCmd.AddCommand(startCmd)
}
