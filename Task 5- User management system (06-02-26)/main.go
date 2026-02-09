package main

import (
	"fmt"
	"sync"
	"userManagementSystem/business"
	"userManagementSystem/models"
)

func main() {
	var wg sync.WaitGroup
	errChan := make(chan error, 20)

	wg.Add(4)

	usersChan := business.ReadCSVFile("assets/users.csv", business.ParseUser, &wg, errChan)
	deptChan := business.ReadCSVFile("assets/departments.csv", business.ParseDepartment, &wg, errChan)
	permChan := business.ReadCSVFile("assets/permissions.csv", business.ParsePermission, &wg, errChan)
	logChan := business.ReadCSVFile("assets/access_logs.csv", business.ParseAccessLog, &wg, errChan)

	// Common error handler
	go func() {
		for err := range errChan {
			fmt.Println("ERROR:", err)
		}
	}()

	//Data structures which stores the data which is recieved from the channel
	usersByID := make(map[string]models.User)
	deptsByID := make(map[string]models.Department)
	perms := []models.Permission{}
	logs := []models.AccessLog{}

	for u := range usersChan {
		usersByID[u.UserID] = u
	}

	for d := range deptChan {
		deptsByID[d.DepartmentID] = d
	}

	for p := range permChan {
		perms = append(perms, p)
	}

	for l := range logChan {
		logs = append(logs, l)
	}

	wg.Wait()

	//closing error channel because the all other channels are closed now
	close(errChan)

	fmt.Println("Done reading all files!")
}
