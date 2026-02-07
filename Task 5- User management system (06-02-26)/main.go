package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

//CSV file worker that will be called for every file
func ReadCsvFile(fileName string, wg *sync.WaitGroup, channel chan [][]string) {
	defer wg.Done()
	file, err1 := os.Open(fileName)

	//File opening error handled
	if err1 != nil {
		fmt.Printf("\n============Problem in opening the file %s=============\n",fileName)
		channel <- nil
		return
	}

	reader := csv.NewReader(file)
	data, err2 := reader.ReadAll()

	//file reading error handled
	if err2 != nil {
		fmt.Printf("\n============Problem in reading the file %s=============\n",fileName)
		channel <- nil
		return
	}

	//data is feeded to the channel
	channel <- data

	fmt.Printf("\nDone Reading %s file\n", fileName)
}

func main() {

	var wg sync.WaitGroup

	//these are the channels which will get the data from each file
	usersChannel := make(chan [][]string, 1)
	departmentsChannel := make(chan [][]string, 1)
	permissionsChannel := make(chan [][]string, 1)
	access_logsChannel := make(chan [][]string, 1)


	wg.Add(4)
	go ReadCsvFile("assets/users.csv", &wg, usersChannel)
	go ReadCsvFile("assets/departments.csv", &wg, departmentsChannel)
	go ReadCsvFile("assets/permissions.csv", &wg, permissionsChannel)
	go ReadCsvFile("assets/access_logs.csv", &wg, access_logsChannel)

	wg.Wait()

}
