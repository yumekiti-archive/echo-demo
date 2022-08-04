package domain

type Tasks []Task

type Task struct {
	ID        int
	Title			string
	Context		string
	Status		bool
}