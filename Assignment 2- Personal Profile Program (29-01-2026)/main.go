package main

import "fmt"


func main() {
	initialize()
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
			takeData()
		case 2:
			introduceMyself()
		case 3:
			updateAge()
		case 4:
			checkForVote()
		default:
			fmt.Print("\n")
			fmt.Println("No such option Available")
			fmt.Print("\n")
		}
	}

}
