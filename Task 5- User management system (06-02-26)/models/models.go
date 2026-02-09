package models

import "time"

type User struct {
	UserID       string    `csv:"user_id" json:"user_id"`
	Name         string    `csv:"name" json:"name"`
	Email        string    `csv:"email" json:"email"`
	Role         string    `csv:"role" json:"role"`
	DepartmentID string    `csv:"department_id" json:"department_id"`
	Status       string    `csv:"status" json:"status"`
	CreatedAt    time.Time `csv:"created_at" json:"created_at"`
	LastLogin    time.Time `csv:"last_login" json:"last_login"`
}

type Department struct {
	DepartmentID   string  `csv:"department_id" json:"department_id"`
	DepartmentName string  `csv:"department_name" json:"department_name"`
	ManagerID      string  `csv:"manager_id" json:"manager_id"`
	Budget         float64 `csv:"budget" json:"budget"`
	Location       string  `csv:"location" json:"location"`
	EmployeeCount  int     `csv:"employee_count" json:"employee_count"`
}

type Permission struct {
	PermissionID string    `csv:"permission_id" json:"permission_id"`
	UserID       string    `csv:"user_id" json:"user_id"`
	Resource     string    `csv:"resource" json:"resource"`
	CanRead      bool      `csv:"can_read" json:"can_read"`
	CanWrite     bool      `csv:"can_write" json:"can_write"`
	CanDelete    bool      `csv:"can_delete" json:"can_delete"`
	CanAdmin     bool      `csv:"can_admin" json:"can_admin"`
	GrantedDate  time.Time `csv:"granted_date" json:"granted_date"`
	ExpiresDate  time.Time `csv:"expires_date" json:"expires_date"`
}

type AccessLog struct {
	LogID          string    `csv:"log_id" json:"log_id"`
	UserID         string    `csv:"user_id" json:"user_id"`
	Resource       string    `csv:"resource" json:"resource"`
	Action         string    `csv:"action" json:"action"`
	Timestamp      time.Time `csv:"timestamp" json:"timestamp"`
	IPAddress      string    `csv:"ip_address" json:"ip_address"`
	StatusCode     int       `csv:"status_code" json:"status_code"`
	ResponseTimeMS int       `csv:"response_time_ms" json:"response_time_ms"`
}
