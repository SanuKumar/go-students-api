package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // this is getting used indirectly
	"github.com/sanukumar/go-students-api/internal/config"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) { //convention use to create instance of the struct
	db, err := sql.Open("sqlite3", cfg.StoragePath) //need to pass the drive, either postgres

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES(?, ?, ?)") // this help to prevent sql injection

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
