package service

import (
	"myapp/dao"
	"myapp/model"
	"strings"
)

type EmployeeService struct {
	dao dao.EmployeeDao
}

func(service EmployeeService) GetEmployee(id int) *model.Employee{
	if id<=0{
		panic("invalid id supplied")
	}
	emp,err:=service.dao.FindById(id)
	if err!=nil{
		return nil
	}
	return &emp
}

func(service EmployeeService) AddNewEmployee(emp model.Employee)model.Employee{
	emp.Name=strings.TrimSpace(emp.Name)
	if emp.Name==""{
		panic("employee name cannot be blank")
	}
	if len(emp.Name)>50{
		panic("employee name cannot exceed 50 chars")
	}
	if emp.Salary<25000{
		panic("Salary must be >=25000")
	}
	if emp.Salary>500000{
		panic("Salary must be <=500000")
	}
	newId,err:=service.dao.Save(emp)
	if err!=nil{
		panic(err.Error())
	}
	emp.Id=newId
	return emp
}