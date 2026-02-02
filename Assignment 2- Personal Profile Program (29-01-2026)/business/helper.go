package business

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	// Id int 		`validate:"required,number"`
	Name string `validate:"required,alphaspace"`
	Age  int    `validate:"required,number"`
}

func (p *Person) TakeData() {
	validate := validator.New()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the Name:")
	name, _ := reader.ReadString('\n')
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter the Age:")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	p.Age, _ = strconv.Atoi(age)

	err := validate.Struct(p)

	if err == nil {
		fmt.Println("************Entered Successfully!***********")
	} else {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("Wrong entry in :", err.Field())
		}
		fmt.Println()
	}
}

func (p *Person) IntroduceMyself() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the your name:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	// data, ok := d.mp[name]
	// if !ok {
	// 	fmt.Print("\nName not found!\n")
	// 	return
	// }

	fmt.Println("\n***********Your Personal Data**********")
	fmt.Printf("Name is:%s\n", p.Name)
	fmt.Printf("Age is: %d\n", p.Age)

}

// func (p *Person) UpdateAge() {
// 	validate := validator.New()
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter the your name:")
// 	name, _ := reader.ReadString('\n')
// 	name = strings.TrimSpace(name)

// 	_, ok := d.mp[name]
// 	if !ok {
// 		fmt.Print("\nName not found!\n")
// 		return
// 	}

// 	fmt.Print("Enter the Age to update:")
// 	age, _ := reader.ReadString('\n')
// 	age = strings.TrimSpace(age)

// 	map := d.map
// 	mp[name].Age, _ = strconv.Atoi(age)

// 	err := validate.Struct(d.mp[name])

// 	if err == nil {
// 		fmt.Println("************Entered Successfully!***********")
// 	} else {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			fmt.Println("\nWrong entry in :", err.Field())
// 		}
// 		fmt.Println()
// 	}
// }

func (p *Person) CheckForVote() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the your name:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	// data, ok := d.mp[name]
	// if !ok {
	// 	fmt.Print("\nName not found!\n")
	// 	return
	// }

	if p.Age >= 18 {
		fmt.Print("\nYes, you can vote!")
	} else {
		fmt.Printf("\nSorry! you can't vote. You can vote after %d years", (18 - p.Age))
	}
}
