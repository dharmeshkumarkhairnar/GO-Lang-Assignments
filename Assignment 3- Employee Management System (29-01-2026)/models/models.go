package models

type EmployeeData struct {
	EmpName   string
	EmpAge    int
	EmpSalary int
}

type DepartmentData struct {
	DeptName string
	List     []EmployeeData
}
