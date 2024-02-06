package usertask

type User struct {
	ID       int
	Email    string
	Password string
}

type Task struct {
	IDTask int
	Name   string
	Status string
}
