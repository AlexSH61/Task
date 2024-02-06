package db

import usertask "github.com/AlexSH61/firstRestAPi/internal/app/tasks"

type Tasks_repo struct {
	db *DataBase
}


func (tr *Task_repo) Create(*usertask.Task) error {
	_, err := r.db.db.Exec("INSERT INTO tasks (title, status) VALUES ($1, $2)", task.Title, task.Status)
	return err
}

func (r *Task_repo) Get(IDTask int) (*Task, error) {
	task := &Task{}
	err := r.db.db.QueryRow("SELECT id, title, status FROM tasks WHERE id = $1", idTask).Scan(&task.ID, &task.Title, &task.Status)
	if err == sql.ErrNoRows {
		return nil, nil 
	} else if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *Task_repo) Update(IDTask int, *usertask.Task) error {
	_, err := r.db.db.Exec("UPDATE tasks SET title = $1, status = $2 WHERE id = $3", task.Title, task.Status, IDTask)
	return err
}

func (r *Task_repo) Delete(IDTask int) error {
	_, err := r.db.db.Exec("DELETE FROM tasks WHERE id = $1", IDTask)
	return err
}
