package db

import (
	"fmt"
	"strings"
	"testing"
)

// testing db
func TestDB(t *testing.T, databaseURL string) (*DataBase, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.DataBaseUrl = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
