package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type employeeData struct {
	empName   string
	empAge    int
	empSalary int
}

type departmentData struct {
	deptName string
	list     []employeeData
}

var alldepartments map[string]departmentData

func initialize() {
	alldepartments = make(map[string]departmentData)
}

func (D departmentData) showData() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	data, ok := alldepartments[department]

	if !ok {
		fmt.Println("Such department doesn't exist!")
	} else {
		if len(data.list) == 0 {
			fmt.Println("No records available!")
			return
		}
		for _, d := range data.list {
			fmt.Println("")
			fmt.Printf("Name: %s , Age: %d , Salary: %d\n", d.empName, d.empAge, d.empSalary)
			fmt.Println("")
		}
	}

}

func (D departmentData) averageSalary() int {

	totalSalary := 0
	countEmployee := 0
	exist := false

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	for key, value := range alldepartments {
		if key == department {
			exist = true
			for _, d := range value.list {
				totalSalary += d.empSalary
				countEmployee++
			}
		}
	}
	if !exist {
		return -1
	}
	return totalSalary / countEmployee
}

func (D departmentData) addEmployee() {
	var e employeeData

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	fmt.Print("Enter the Name of employee:")
	name, _ := reader.ReadString('\n')
	e.empName = strings.TrimSpace(name)

	fmt.Print("Enter the Age of employee:")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	e.empAge, _ = strconv.Atoi(age)

	fmt.Print("Enter the salary of employee:")
	salary, _ := reader.ReadString('\n')
	salary = strings.TrimSpace(salary)
	e.empSalary, _ = strconv.Atoi(salary)

	data, ok := alldepartments[department]
	var d departmentData

	if ok {
		d.deptName = department
		d.list = data.list
		d.list = append(d.list, e)
		alldepartments[department] = d
	} else {
		d.deptName = department
		d.list = append(d.list, e)
		alldepartments[department] = d
	}
}

func (D departmentData) removeEmployee() {
	reader := bufio.NewReader(os.Stdin)
	deleted := false

	fmt.Print("Enter the Department Name of employee:")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	fmt.Print("Enter the Name of employee:")
	empName, _ := reader.ReadString('\n')
	empName = strings.TrimSpace(empName)

	for key, value := range alldepartments {
		if key == department {
			for i := 0; i < len(value.list); i++ {
				if value.list[i].empName == empName {
					value.list = slices.Delete(value.list, i, i+1)
					alldepartments[department] = value
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

func (D departmentData) giveRaise() {
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

	for key, value := range alldepartments {
		if key == department {
			for i := 0; i < len(value.list); i++ {
				if value.list[i].empName == empName {
					value.list[i].empSalary = value.list[i].empSalary + increase
					alldepartments[department] = value
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
