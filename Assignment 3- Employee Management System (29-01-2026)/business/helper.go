package business

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"github.com/go-playground/validator/v10"
)

type EmployeeData struct {
	Id        int    `validate:"required,number"`
	EmpName   string `validate:"required,alphaspace"`
	EmpAge    int    `validate:"required,number"`
	EmpSalary int    `validate:"required,number"`
}

type DepartmentData struct {
	DeptName string `validate:"required,alphaspace"`
	List     []EmployeeData 
}

var Alldepartments map[string]*DepartmentData
var EmployeeIDs map[int]int 

func Initialize() {
	Alldepartments = make(map[string]*DepartmentData)
	EmployeeIDs=make(map[int]int)
}

func (D *DepartmentData) ShowData() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)
	department =strings.ToLower(department)

	data, ok := Alldepartments[department]

	if !ok {
		fmt.Println("Such department doesn't exist!")
	} else {
		if len(data.List) == 0 {
			fmt.Println("No records available!")
			return
		}
		for _, d := range data.List {
			fmt.Printf("ID: %d ,Name: %s , Age: %d , Salary: %d\n", d.Id, d.EmpName, d.EmpAge, d.EmpSalary)
		}
	}

}

func (D *DepartmentData) AverageSalary() {

	totalSalary := 0
	countEmployee := 0
	// exist := false

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)
	department =strings.ToLower(department)

	dept, ok := Alldepartments[department]

	if !ok {
		fmt.Println("No Such Data Found!")
		return
	}

	for _, emp := range dept.List {
		totalSalary += emp.EmpSalary
		countEmployee++
	}

	fmt.Println("Average is ", totalSalary/countEmployee)
}

func (D *DepartmentData) AddEmployee() {
	var e EmployeeData
	validate := validator.New()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)
	department =strings.ToLower(department)

	err1 := validate.Var(department, "required,alphanumspace")

	if err1 != nil {
		fmt.Println("Wrong Entry in department name")
		return
	}

	fmt.Print("Enter the ID of employee:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	e.Id, _ = strconv.Atoi(id)

	fmt.Print("Enter the Name of employee:")
	name, _ := reader.ReadString('\n')
	e.EmpName = strings.TrimSpace(name)

	fmt.Print("Enter the Age of employee:")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	e.EmpAge, _ = strconv.Atoi(age)

	fmt.Print("Enter the salary of employee:")
	salary, _ := reader.ReadString('\n')
	salary = strings.TrimSpace(salary)
	e.EmpSalary, _ = strconv.Atoi(salary)

	err := validate.Struct(e)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Wrong entry in :", err.Field())
		}
		return
	}

	_,checkId:=EmployeeIDs[e.Id]

	if checkId {
		fmt.Println("Employee ID already exist!")
		return
	}

	EmployeeIDs[e.Id]++

	_, ok := Alldepartments[department]

	if ok {
		Alldepartments[department].List = append(Alldepartments[department].List, e)
		return
	}

	Alldepartments[department] = &DepartmentData{
		DeptName: department,
		List:     []EmployeeData{e},
	}
}

func (D *DepartmentData) RemoveEmployee() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)
	department =strings.ToLower(department)

	fmt.Print("Enter the ID of employee:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	empId, _ := strconv.Atoi(id)

	dept, ok := Alldepartments[department]
	if !ok {
		fmt.Println("No Such Data Found!")
		return
	}

	for i, emp := range dept.List {
		if emp.Id == empId {
			dept.List = slices.Delete(dept.List, i, i+1)
			fmt.Println("Employee Data Deleted Successfully!")
			delete(EmployeeIDs,empId)
			return
		}
	}

	fmt.Println("No Such Data Found!")

}

func (D *DepartmentData) GiveRaise() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)
	department =strings.ToLower(department)

	fmt.Print("Enter the ID of employee:")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	empId, _ := strconv.Atoi(id)

	fmt.Print("Enter the raise amount for employee:")
	raise, _ := reader.ReadString('\n')
	raise = strings.TrimSpace(raise)
	increase, _ := strconv.Atoi(raise)

	dept, ok := Alldepartments[department]
	if !ok {
		fmt.Println("No Such Data Found!")
		return
	}

	for i := range dept.List {
		if dept.List[i].Id == empId {
			dept.List[i].EmpSalary += increase
			fmt.Println("Employee Salary changed Successfully!")
			return
		}
	}

	fmt.Println("No Such Data Found!")
}
