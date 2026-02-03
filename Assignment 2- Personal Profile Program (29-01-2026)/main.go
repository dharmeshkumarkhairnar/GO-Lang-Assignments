package main

import (
	"assignment2/business"
	"fmt"

	//"github.com/go-playground/validator/v10"
)

func main() {
	// var d business.Data
	// d.Initialize()
	//validate := validator.New()
	var p business.Person
	for true {
		n := 0
		fmt.Println("What you want to do?")
		fmt.Println("1.Add Personal Info\n2.Introduce Myself\n3.Update Age\n4.Check whether I can Vote\n5.Exit\nChoose the number:->")
		fmt.Scanln(&n)
		if n == 5 {
			fmt.Println("Exited!")
			break
		}
		switch n {
		case 1:
			p.TakeData()
		case 2:
			p.IntroduceMyself()
		case 3:
			{
				if len(p.Name) == 0 {
					fmt.Println("Enter the Name and age first then you can update age")
					continue
				}
				fmt.Println("Enter the age")
				newAge := 0
				_,err:=fmt.Scanln(&newAge)

				if err!=nil {
					fmt.Println("Wrong Age entered1")
					var remove string
					fmt.Scanln(&remove)
				}
				if p.Age > newAge {
					fmt.Println("Wrong Entry in Age")
					continue
				}
				p.Age=newAge	
			}
		case 4:
			p.CheckForVote()
		default:
			fmt.Println("No such option Available")
		}
	}
}
