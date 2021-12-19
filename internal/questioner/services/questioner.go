package services

import (
	"github.com/sinojin/questioner/internal/questioner/domain/question"
	"github.com/sinojin/questioner/internal/questioner/domain/statistic"
)

type Questions struct {
	Q []question.Question
}

type Answers struct {
	A []question.Answer
}

type Results struct {
	TrueAnswerNum int
	Ratio         int
}

//QuestionerService is interface because we can make it wrapper
//for instance performance wrapper
type QuestionerService interface {
	GetAllQuestions() (Questions, error)
	CalculateQuestion(Answers) (Results, error)
}

//questionService is implementing of QuestionerService
type questionService struct {
	qRepo question.QuestionRepository
	sRepo statistic.StatisticRepository
}

func NewQuestionService(qRepo question.QuestionRepository, sRepo statistic.StatisticRepository) QuestionerService {
	return &questionService{qRepo: qRepo, sRepo: sRepo}
}

func (qs *questionService) GetAllQuestions() (Questions, error) {
	questions, _ := qs.qRepo.Get()
	return Questions{questions}, nil
}

func (qs *questionService) CalculateQuestion(answers Answers) (Results, error) {
	//true answer number
	trueN := 0
	//false answer number
	falseN := 0
	//worst player number
	worstPN := 0
	//total player number
	totalPN := 0
	//result of service function
	result := Results{}
	for _, answer := range answers.A {
		question, _ := qs.qRepo.GetByID(answer.ID)
		if question.IsCorrect(answer.Choice) {
			trueN++
			continue
		}
		falseN++
	}
	result.TrueAnswerNum = trueN
	Statistics, _ := qs.sRepo.Get()
	totalPN = len(Statistics)

	if totalPN == 0 {
		result.Ratio = 100
	} else {
		for _, statistic := range Statistics {
			if trueN > statistic.CorrectN {
				worstPN++
			}
		}
		ratio := int((100 * worstPN) / totalPN)
		result.Ratio = ratio
	}
	qs.sRepo.Save(statistic.Statistic{CorrectN: trueN, WrongN: falseN})
	return result, nil
}
