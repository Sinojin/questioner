package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/sinojin/questioner/internal/common/genproto/questioner"
	"google.golang.org/grpc"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
)

func Cli() {
	for {
		prompt := DoYouWantNewGame()
		result, _ := prompt.Run()
		if strings.ToLower(result) == "n" {
			break
		}
		answerQuestions()
	}

}

func DoYouWantNewGame() promptui.Prompt {
	validate := func(input string) error {
		if strings.ToLower(input) != "y" && strings.ToLower(input) != "n" {
			return errors.New("You can use Y or N")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Do you want to play ?",
		Validate: validate,
	}
	return prompt
}

func answerQuestions() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := questioner.NewQuestionerServiceClient(conn)
	questionResponse, err := client.GetAllQuestions(context.Background(), &questioner.Empty{})
	if err != nil {
		panic(err)
	}
	answers := askQuestions(questionResponse)

	response, err := client.CalculateAllQuestions(context.Background(), answers)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Your Correct answers number is %v .\n", response.GetCorrectAnswerNumber())
	fmt.Printf("You scored higher than %v%v of all quizzers .\n", response.GetRatio(), "%")

}

func askQuestions(questions *questioner.Questions) *questioner.Answers {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Description | cyan }}",
		Inactive: "  {{ .Description | cyan }} ",
		Selected: "\U0001F336 {{ .Description | red | cyan }}",
	}
	answers := make([]*questioner.Answer, 0)
	for _, q := range questions.GetQuestions() {
		prompt := promptui.Select{
			Label:     q.Description,
			Items:     q.GetChoices(),
			Templates: templates,
			Size:      3,
		}
		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return nil
		}
		answers = append(answers, &questioner.Answer{
			ID:     q.ID,
			Choice: int32(i),
		})
	}
	return &questioner.Answers{Answers: answers}
}
