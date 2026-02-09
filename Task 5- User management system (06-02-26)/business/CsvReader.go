package business

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

//This function takes CSV file to read and return the channel of file data
func ReadCSVFile[T comparable](filePath string, parseFn func([]string) (T, error), wg *sync.WaitGroup, errChan chan<- error) <-chan T {

	resultChannel := make(chan T)

	go func() {
		defer wg.Done()
		defer close(resultChannel)

		file, fileOpenError := os.Open(filePath)
		if fileOpenError != nil {
			errChan <- fmt.Errorf("open %s: %w", filePath, fileOpenError)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)

		//Reading Header line and skiping it
		if _, headerError := reader.Read(); headerError != nil {
			errChan <- fmt.Errorf("read header %s: %w", filePath, headerError)
			return
		}

		for {
			record, readLineError := reader.Read()
			if readLineError != nil {
				break
			}

			parsed, parsingError := parseFn(record)
			if parsingError != nil {
				errChan <- fmt.Errorf("parse %s: %w", filePath, parsingError)
				continue
			}

			resultChannel <- parsed
		}
	}()

	return resultChannel
}
