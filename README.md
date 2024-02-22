# Code Test Quiz

As discussed, here is the test:

## Task Description

Your task is to build a simple quiz application with a few questions, each having multiple choice answers with only one correct answer.

### Preferred Stack

- **Backend:** Golang
- **Database:** In-memory (no database required)

### Preferred Components

- REST API or gRPC
- CLI that communicates with the API, preferably using [Cobra](https://github.com/spf13/cobra) as the CLI framework

## User Stories/Use Cases

1. User should be able to retrieve questions with multiple answers.
2. User should be able to select only one answer per question.
3. User should be able to answer all the questions, submit their answers, and receive feedback on the number of correct answers.
4. User should be able to see how their performance compares to others who have taken the quiz (e.g., "You performed better than 60% of all participants").

Please let us know when you would be able to complete this task.

If you have any questions based on our call or the test above, don't hesitate to contact me.

Wishing you a great day!

Kind regards,


### Implementaion
go install github.com/spf13/cobra-cli@latest
cobra-cli init quiz-cli
cobra-cli add start
