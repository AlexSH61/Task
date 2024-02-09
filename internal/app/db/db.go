package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBase struct {
	config   *Config
	db       *sql.DB
	userRepo *User_repo
	taskRepo *Task_repo
}

func New(config *Config) *DataBase {
	return &DataBase{
		config: config,
	}
}
func (d *DataBase) Open() error {
	db, err := sql.Open("postgres", d.config.DataBaseUrl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	d.db = db
	return nil
}
func (s *DataBase) Close() {
	s.db.Close()

}
func (s *DataBase) User() *User_repo {
	if s.userRepo != nil {
		return s.userRepo
	}
	s.userRepo = &User_repo{
		db: s,
	}
	return s.userRepo
}

func (t *DataBase) Task() *Task_repo {
	if t.taskRepo != nil {
		return t.taskRepo
	}
	t.taskRepo = &Task_repo{
		db: t,
	}
	return t.taskRepo
}
