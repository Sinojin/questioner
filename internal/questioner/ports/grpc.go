package ports

import (
	"context"
	pb "github.com/sinojin/questioner/internal/common/genproto/questioner"
	"github.com/sinojin/questioner/internal/questioner/domain/question"
	"github.com/sinojin/questioner/internal/questioner/services"
)

type GrpcServer struct {
	s services.QuestionerService
}

func NewGrpcServer(s services.QuestionerService) *GrpcServer {
	return &GrpcServer{s}
}

func (g GrpcServer) GetAllQuestions(ctx context.Context, in *pb.Empty) (*pb.Questions, error) {
	qsts, _ := g.s.GetAllQuestions()

	result := new(pb.Questions)
	//translate domain questions to pb questions
	qs := make([]*pb.Question, 0)

	for _, question := range qsts.Q {
		qs = append(qs, domainQuestionToPbQuestion(question))
	}
	result.Questions = qs
	return result, nil
}

func (g GrpcServer) CalculateAllQuestions(ctx context.Context, in *pb.Answers) (*pb.Staticstic, error) {
	a := pbAnswerToServiceAnswers(in)
	r, _ := g.s.CalculateQuestion(a)
	return &pb.Staticstic{
		CorrectAnswerNumber: int32(r.TrueAnswerNum),
		Ratio:               int32(r.Ratio),
	}, nil
}
func pbAnswerToServiceAnswers(in *pb.Answers) services.Answers {
	answers := make([]question.Answer, 0)
	for _, a := range in.Answers {
		answers = append(answers, question.Answer{
			ID:     a.GetID(),
			Choice: int(a.GetChoice()),
		})
	}
	return services.Answers{answers}
}

func domainChoiceToPbChoice(choice []question.Choice) []*pb.Choice {
	choices := make([]*pb.Choice, 0)
	for _, chs := range choice {
		choices = append(choices, &pb.Choice{
			ID:          int32(chs.ID),
			Description: chs.Description,
		})
	}
	return choices
}
func domainQuestionToPbQuestion(q question.Question) *pb.Question {
	chs := domainChoiceToPbChoice(q.Choices)
	return &pb.Question{
		ID:          q.ID,
		Description: q.Description,
		Choices:     chs,
	}
}
