package domain

type QuestionCriterion struct {
	Criterion string
	Score     int64
}

type Question struct {
	Question        string
	PossibleAnswers []string
}
