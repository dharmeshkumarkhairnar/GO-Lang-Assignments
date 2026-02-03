package main

import (
	"employeeManagementSystem/business"
	"fmt"
)

func main() {
	var D business.DepartmentData
	business.Initialize()
	
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
			D.AddEmployee()
		case 2:
			D.AverageSalary()
		case 3:
			D.RemoveEmployee()
		case 4:
			D.GiveRaise()
		case 5:
			D.ShowData()
		default:
			fmt.Println("No such option Available")
		}
	}
}
