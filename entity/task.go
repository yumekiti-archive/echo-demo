package entity

type Tasks []Task

type Task struct {
	ID      int
	Title   string
	Context string
	Status  bool
}
