package storage

import "github.com/sanukumar/go-students-api/internal/types"

// interface --> can be used to create this as plugin

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
