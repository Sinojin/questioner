package question

type QuestionRepository interface {
	//Get returns question array
	//maybe ctx is needed, but it's basic repo
	Get() ([]Question, error)
	GetByID(string) (Question, error)
	//I removed save because predefined json we can use
	//Save(question Question) error
}
