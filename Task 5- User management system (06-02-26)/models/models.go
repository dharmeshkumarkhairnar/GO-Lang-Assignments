package models

import "time"

type Users struct {
	Id         int       `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Dept_id    string    `json:"department_id"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"created_at"`
	Last_login time.Time `json:"last_login"`
}

type Departments struct {
	Dept_id        int     `json:"department_id"`
	Dept_name      string  `json:"department_name"`
	Manager_id     int     `json:"manager_id"`
	Budget         float64 `json:"budget"`
	Location       string  `json:"location"`
	Employee_count int     `json:"employee_count"`
}

type Permissions struct {
	Permission_id int `json:"permission_id"`
	User_id int `json:"user_id"`
	Resource string `json:"resource"`
	Can_read bool `json:""`
	Can_write bool `json:""`
	Can_delete bool `json:""`
	Can_admin bool `json:""`
	Granted_date time.Time `json:""`
	Expire_date time.Time `json:""`
}
