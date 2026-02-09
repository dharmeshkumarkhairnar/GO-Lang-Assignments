package business

import (
	"strconv"
	"time"
	"userManagementSystem/models"
)

const (
	dateTimeLayout = "2006-01-02 15:04:05"
	dateLayout     = "2006-01-02"
)

//Funtion which parse the user file data
func ParseUser(row []string) (models.User, error) {
	createdAt, createdAtErr := time.Parse(dateLayout, row[6])
	if createdAtErr != nil {
		return models.User{}, createdAtErr
	}

	lastLogin, lastLoginErr := time.Parse(dateTimeLayout, row[7])
	if lastLoginErr != nil {
		return models.User{}, lastLoginErr
	}

	return models.User{
		UserID:       row[0],
		Name:         row[1],
		Email:        row[2],
		Role:         row[3],
		DepartmentID: row[4],
		Status:       row[5],
		CreatedAt:    createdAt,
		LastLogin:    lastLogin,
	}, nil
}

//Funtion which parse the department file data
func ParseDepartment(row []string) (models.Department, error) {
	budget, budgetErr := strconv.ParseFloat(row[3], 64)
	if budgetErr != nil {
		return models.Department{}, budgetErr
	}

	empCount, empCountErr := strconv.Atoi(row[5])
	if empCountErr != nil {
		return models.Department{}, empCountErr
	}

	return models.Department{
		DepartmentID:   row[0],
		DepartmentName: row[1],
		ManagerID:      row[2],
		Budget:         budget,
		Location:       row[4],
		EmployeeCount:  empCount,
	}, nil
}


//Funtion which parse the permission file data
func ParsePermission(row []string) (models.Permission, error) {
	canRead, canReadErr := strconv.ParseBool(row[3])
	if canReadErr != nil {
		return models.Permission{}, canReadErr
	}

	canWrite, canWriteErr := strconv.ParseBool(row[4])
	if canWriteErr != nil {
		return models.Permission{}, canWriteErr
	}

	canDelete, canDeleteErr := strconv.ParseBool(row[5])
	if canDeleteErr != nil {
		return models.Permission{}, canDeleteErr
	}

	canAdmin, canAdminErr := strconv.ParseBool(row[6])
	if canAdminErr != nil {
		return models.Permission{}, canAdminErr
	}

	grantedDate, grantedDateErr := time.Parse(dateLayout, row[7])
	if grantedDateErr != nil {
		return models.Permission{}, grantedDateErr
	}

	expiresDate, expiresDateErr := time.Parse(dateLayout, row[8])
	if expiresDateErr != nil {
		return models.Permission{}, expiresDateErr
	}

	return models.Permission{
		PermissionID: row[0],
		UserID:       row[1],
		Resource:     row[2],
		CanRead:      canRead,
		CanWrite:     canWrite,
		CanDelete:    canDelete,
		CanAdmin:     canAdmin,
		GrantedDate:  grantedDate,
		ExpiresDate:  expiresDate,
	}, nil
}

//Funtion which parse the access logs file data
func ParseAccessLog(row []string) (models.AccessLog, error) {
	timestamp, timestampErr := time.Parse(dateTimeLayout, row[4])
	if timestampErr != nil {
		return models.AccessLog{}, timestampErr
	}

	statusCode, statusCodeErr := strconv.Atoi(row[6])
	if statusCodeErr != nil {
		return models.AccessLog{}, statusCodeErr
	}

	respTime, respTimeErr := strconv.Atoi(row[7])
	if respTimeErr != nil {
		return models.AccessLog{}, respTimeErr
	}

	return models.AccessLog{
		LogID:          row[0],
		UserID:         row[1],
		Resource:       row[2],
		Action:         row[3],
		Timestamp:      timestamp,
		IPAddress:      row[5],
		StatusCode:     statusCode,
		ResponseTimeMS: respTime,
	}, nil
}
