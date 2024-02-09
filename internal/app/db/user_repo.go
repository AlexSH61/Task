package db

import "github.com/AlexSH61/firstRestAPi/internal/app/model"

type User_repo struct {
	db *DataBase
}

func (r *User_repo) Create(u *model.User) (*model.User, error) {
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	if err := r.db.db.QueryRow(
		"INSERT INTO users(email,encrypto_password) VALUES($1,$2) RETURNING id",
		u.Email,
		u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}
func (r *User_repo) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.db.db.QueryRow(
		"SELECT id, email,encrypto_password FROM users WHERE email = $1",
		email).Scan(
		&u.ID, &u.Email, &u.Password); err != nil {
		return nil, err
	}
	return &model.User{}, nil
}
func (r *User_repo) UpdateUser(u *model.User) error {
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
