package db

import (
	usertasks "github.com/AlexSH61/firstRestAPi/internal/app/tasks"
)

type User_repo struct {
	db *DataBase
}

func (r *User_repo) Create(u *usertasks.User) (*usertasks.User, error) {
	if err := r.db.db.QueryRow(
		"INSERT INTO users(email, password) VALUES($1,$2) RETURNING id",
		u.Email,
		u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}
func (r *User_repo) FindByEmail(email string) (*usertasks.User, error) {
	return nil, nil
}
func (r *User_repo) UpdateUser(u *usertasks.User) error {
	_, err := r.db.db.Exec(
		"UPDATE tasks SET name = $1, description = $2 WHERE id = $3",
		u.ID,
		u.Password,
		u.ID,
	)
	return err
}

func (r *User_repo) DeleteUser(ID int) error {
	_, err := r.db.db.Exec("DELETE FROM tasks WHERE id = $1", ID)
	return err
}
