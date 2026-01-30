package business

import (
	"bufio"
	"employeeManagementSystem/models"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Employee struct {
	Alldepartments map[string]models.DepartmentData
}

func (E *Employee) NewEmployeeData() {
	(*E).Alldepartments=make(map[string]models.DepartmentData)
}

func (E *Employee) ShowData() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	data, ok := E.Alldepartments[department]

	if !ok {
		fmt.Println("")
		fmt.Println("Such department doesn't exist!")
		fmt.Println("")
	} else {
		if len(data.List) == 0 {
			fmt.Println("No records available!")
			return
		}
		for _, d := range data.List {
			fmt.Println("")
			fmt.Printf("Name: %s , Age: %d , Salary: %d\n", d.EmpName, d.EmpAge, d.EmpSalary)
		}
		fmt.Println("")
	}

}

func (E *Employee) AverageSalary() int {

	totalSalary := 0
	countEmployee := 0
	exist := false

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	for key, value := range E.Alldepartments {
		if key == department {
			exist = true
			for _, d := range value.List {
				totalSalary += d.EmpSalary
				countEmployee++
			}
		}
	}
	if !exist {
		return -1
	}
	return totalSalary / countEmployee
}

func (E *Employee) AddEmployee() {
	var e models.EmployeeData

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _:= reader.ReadString('\n')
	department = strings.TrimSpace(department)

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

	data, ok := E.Alldepartments[department]
	var d models.DepartmentData

	if ok {
		d.DeptName = department
		d.List = data.List
		d.List = append(d.List, e)
		E.Alldepartments[department] = d
	} else {
		d.DeptName = department
		d.List = append(d.List, e)
		E.Alldepartments[department] = d
	}
}

func (E *Employee) RemoveEmployee() {
	reader := bufio.NewReader(os.Stdin)
	deleted := false

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	fmt.Print("Enter the Name of employee:")
	empName, _ := reader.ReadString('\n')
	empName = strings.TrimSpace(empName)

	for key, value := range E.Alldepartments {
		if key == department {
			for i := 0; i < len(value.List); i++ {
				if value.List[i].EmpName == empName {
					value.List = slices.Delete(value.List, i, i+1)
					E.Alldepartments[department] = value
					deleted = true
				}
			}
		}
	}

	if deleted {
		fmt.Println("Employee Data Deleted Sucessfully!")
	} else {
		fmt.Println("No Such Data Found!")
	}

}

func (E *Employee) GiveRaise() {
	reader := bufio.NewReader(os.Stdin)
	updated := false

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	fmt.Print("Enter the Name of employee:")
	empName, _ := reader.ReadString('\n')
	empName = strings.TrimSpace(empName)

	fmt.Print("Enter the raise amount for employee:")
	raise, _ := reader.ReadString('\n')
	raise = strings.TrimSpace(raise)
	increase, _ := strconv.Atoi(raise)

	for key, value := range E.Alldepartments {
		if key == department {
			for i := 0; i < len(value.List); i++ {
				if value.List[i].EmpName == empName {
					value.List[i].EmpSalary = value.List[i].EmpSalary + increase
					E.Alldepartments[department] = value
					updated = true
				}
			}
		}
	}

	if updated {
		fmt.Println("Employee Salary changed Sucessfully!")
	} else {
		fmt.Println("No Such Data Found!")
	}
}
