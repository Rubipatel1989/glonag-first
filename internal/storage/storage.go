package storage

import "goLangFirst/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudentList() ([]types.Student, error)
}
