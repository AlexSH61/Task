package apiserver

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/AlexSH61/firstRestAPi/internal/app/db"
	"github.com/AlexSH61/firstRestAPi/internal/app/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var ()

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	db     *db.DataBase
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configTask(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}

func (s *APIserver) configTask() error {
	st := db.New(s.config.Task)
	if err := st.Open(); err != nil {
		return err
	}
	s.db = st
	return nil
}
func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/tasks", s.GetAllTasksHandler).Methods("GET")
	s.router.HandleFunc("/tasks", s.CreateTaskHandler).Methods("POST")
	s.router.HandleFunc("/tasks/{id:[0-9]+}", s.UpdateTaskStatusHandler).Methods("PUT")
	s.router.HandleFunc("/tasks/{id:[0-9]+}", s.DeleteTaskHandler).Methods("DELETE")
}

func (s *APIserver) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.db.Task().GetAllTasks(1)
	if err != nil {
		s.logger.Errorf("Failed to get tasks: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		s.logger.Errorf("Failed to encode tasks to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *APIserver) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var newTask model.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		s.logger.Errorf("Failed to decode JSON request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	createdTask, err := s.db.Task().Create(&newTask)
	if err != nil {
		s.logger.Errorf("Failed to create task: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createdTask)
	if err != nil {
		s.logger.Errorf("Failed to encode created task to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *APIserver) UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	var updateData struct {
		IDTask int    `json:"idtask"`
		Status string `json:"status"`
	}

	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		s.logger.Errorf("Failed to decode JSON request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = s.db.Task().UpdateTaskStatus(&model.Task{})

	if err != nil {
		s.logger.Errorf("Failed to update task status: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *APIserver) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.logger.Errorf("Invalid task ID: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = s.db.Task().Delete(id)
	if err != nil {
		s.logger.Errorf("Failed to delete task: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
