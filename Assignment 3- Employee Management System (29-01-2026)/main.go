package main

import (
	"employeeManagementSystem/business"
	"fmt"
)

func main() {
	var E business.Employee
	E.NewEmployeeData()
	//var d models.DepartmentData
	
	for true {
		n := 0
		fmt.Println("What you want to do?")
		fmt.Println("1.Add Employee\n2.Show average salary of department\n3.Delete employee data\n4.Give raise to employee\n5.Show employee for Department\n6.Exit\nChoose the number:->")
		fmt.Scanln(&n)
		if n == 6 {
			fmt.Println("Exited!")
			break
		}
		switch n {
		case 1:
			E.AddEmployee()
		case 2:
			val := E.AverageSalary()
			if val == -1 {
				fmt.Println("No Such Department Exists!")
			} else {
				fmt.Println("Average is ", val)
			}
		case 3:
			E.RemoveEmployee()
		case 4:
			E.GiveRaise()
		case 5:
			E.ShowData()
		default:
			fmt.Print("\n")
			fmt.Println("No such option Available")
			fmt.Print("\n")

		}
	}

	fmt.Println(E.Alldepartments)
}
