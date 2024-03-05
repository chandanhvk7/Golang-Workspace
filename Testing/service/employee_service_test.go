package service

import (
	"errors"
	"myapp/model"
	"testing"
)

type MockEmployeeDao struct {
}

func (dao MockEmployeeDao) FindById(id int) (model.Employee, error) {
	panic("Not ready yet!")

}

func (dao MockEmployeeDao) Save(employee model.Employee) (int, error) {
	if employee.Name == "John" {
		return 1, nil
	}
	if employee.Name == "00" {
		return 0, errors.New("db connect failed")
	}
	return 0, nil
}

//end of mock Implementation of Employee Interface

func TestAddNewEmployee(t *testing.T) {
	var mockDao MockEmployeeDao
	service := EmployeeService{dao: mockDao}
	t.Run("valid employee added", func(t *testing.T) {
		emp := model.Employee{Name: "John", Salary: 44593}
		savedEmp := service.AddNewEmployee(emp)
		if savedEmp.Id != 1 {
			t.Error("saved id is incorrect")
		}
	})
	t.Run("employee with blank name",func(t *testing.T) {
		defer func ()  {
			if r:=recover();r==nil{
				t.Error("was expecting a panic;did not get one")
			}
			
		}()
		emp:=model.Employee{Name: "    ",Salary: 44000}
		service.AddNewEmployee(emp)
		
	})
	t.Run("employee with salary<25000",func(t *testing.T) {
		defer func(){
			if r:=recover();r==nil{
				t.Error("was expecting a panic;did not get one")
			}
			}()
			emp:=model.Employee{Name: "John",Salary: 00}
			service.AddNewEmployee(emp)
		})
	t.Run("employee with salary<25000",func(t *testing.T) {
		defer func(){
			if r:=recover();r==nil{
				t.Error("was expecting a panic;did not get one")
			}
			}()
			emp:=model.Employee{Name: "John",Salary: 500001}
			service.AddNewEmployee(emp)
		})
}

//go test -v -cover ./service
