package adapters

import (
	"errors"
	"github.com/sinojin/questioner/internal/questioner/domain/question"
)

var QuestionCouldntBeFound = errors.New("Question could not be found")

type questionRepository struct {
	////sync.RWMutex is more efficient for m
	//m *sync.Mutex
	list []question.Question
}

func (q *questionRepository) Get() ([]question.Question, error) {
	//q.m.Lock()
	//defer q.m.Unlock()
	return q.list, nil
}
func (q *questionRepository) GetByID(ID string) (question.Question, error) {
	questions, _ := q.Get()
	for _, question := range questions {
		if question.ID == ID {
			return question, nil
		}
	}
	return question.Question{}, QuestionCouldntBeFound
}

func NewQuestionRepository() question.QuestionRepository {
	//added base questions list

	return &questionRepository{list: []question.Question{question1, question2, question3}}
}

var question1 = question.Question{
	ID:          "question1",
	Description: "JSON stands for X ?",
	Choices: []question.Choice{
		question.Choice{
			ID:          0,
			Description: "JavaScript Object Notation",
		}, question.Choice{
			ID:          1,
			Description: "Java Object Notation",
		}, question.Choice{
			ID:          2,
			Description: "JavaScript Object Normalization",
		},
	},
	Answer: 0,
}
var question2 = question.Question{
	ID:          "question2",
	Description: "JSON is a _____ for storing and transporting data.",
	Choices: []question.Choice{
		question.Choice{
			ID:          0,
			Description: "xml format",
		}, question.Choice{
			ID:          1,
			Description: "text format",
		}, question.Choice{
			ID:          2,
			Description: "JavaScript",
		},
	},
	Answer: 1,
}
var question3 = question.Question{
	ID:          "question3",
	Description: "The JSON syntax is a subset of the _____ syntax.",
	Choices: []question.Choice{
		question.Choice{
			ID:          0,
			Description: "Ajax",
		}, question.Choice{
			ID:          1,
			Description: "Php",
		}, question.Choice{
			ID:          2,
			Description: "JavaScript",
		},
	},
	Answer: 2,
}
