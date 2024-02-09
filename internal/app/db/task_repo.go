package db

import (
	"database/sql"

	"github.com/AlexSH61/firstRestAPi/internal/app/model"
)

type Task_repo struct {
	db *DataBase
}

func (tr *Task_repo) Create(task *model.Task) (*model.Task, error) {
	_, err := tr.db.db.Exec("INSERT INTO tasks (title, status) VALUES ($1, $2)", task.Title, task.Status)
	return nil, err
}

func (tr *Task_repo) Get(IDTask int) (*model.Task, error) {
	task := &model.Task{}
	err := tr.db.db.QueryRow("SELECT id, title, status FROM tasks WHERE id = $1", IDTask).Scan(&task.IDTask, &task.Title, &task.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *Task_repo) UpdateTask(task *model.Task) error {
	_, err := tr.db.db.Exec("UPDATE tasks SET title = $1, status = $2 WHERE id = $3", task.IDTask, task.Status)
	return err
}

func (tr *Task_repo) Delete(IDTask int) error {
	_, err := tr.db.db.Exec("DELETE FROM tasks WHERE id = $1", IDTask)
	return err
}
