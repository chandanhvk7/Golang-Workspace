package dao

import "myapp/model"

type EmployeeDao interface {
	//id must be a positive int
	//when there is no matching employee, return with error
	FindById(id int) (model.Employee,error)

	//employee must have name and salary
	//after successful operation, new id is returned 
	Save(employee model.Employee)(int, error)
}