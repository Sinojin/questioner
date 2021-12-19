package question

type Question struct {
	ID          string
	Description string

	Choices []Choice

	Answer int
}

func (q *Question) IsCorrect(answer int) bool {
	return q.Answer == answer
}
