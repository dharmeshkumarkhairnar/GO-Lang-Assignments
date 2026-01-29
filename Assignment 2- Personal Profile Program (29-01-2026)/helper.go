package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/go-playground/validator/v10"
)

type Person struct {
	name string `validate:"required,alpha"`
	age  int    `validate:"required,number"`
}

var mp map[string]Person

func initialize() {
	mp = make(map[string]Person)
}

func takeData() {
	validate := validator.New()
	reader := bufio.NewReader(os.Stdin)
	var p Person

	fmt.Print("Enter the Name:")
	name, _ := reader.ReadString('\n')
	p.name = strings.TrimSpace(name)

	fmt.Print("Enter the Age:")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	p.age, _ = strconv.Atoi(age)

	err := validate.Struct(p)

	if err == nil {
		mp[p.name] = p
		fmt.Println("************Entered Successfully!***********")
		fmt.Print("\n")
	} else {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("")
			fmt.Println("Wrong entry in :", err.Field())
		}
		fmt.Println("")
	}
}

func introduceMyself() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	data, ok := mp[name]
	if !ok {
		fmt.Println("")
		fmt.Print("Name not found!")
		fmt.Println("")
		return
	}

	fmt.Println("")
	fmt.Println("***********Your Personal Data**********")
	fmt.Println("Name is:",data.name)
	fmt.Println("Age is:",data.age)
	fmt.Println("")

}

func updateAge() {
	validate := validator.New()
	reader := bufio.NewReader(os.Stdin)
	var p Person

	fmt.Print("Enter the your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	_, ok := mp[name]
	if !ok {
		fmt.Println("")
		fmt.Print("Name not found!")
		fmt.Println("")
		return
	}

	p.name = name

	fmt.Print("Enter the Age to update:")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	p.age, _ = strconv.Atoi(age)

	err := validate.Struct(p)

	if err == nil {
		mp[p.name] = p
		fmt.Println("************Entered Successfully!***********")
		fmt.Print("\n")
	} else {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("")
			fmt.Println("Wrong entry in :", err.Field())
		}
		fmt.Println("")
	}
}

func checkForVote() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the your name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	data, ok := mp[name]
	if !ok {
		fmt.Println("")
		fmt.Print("Name not found!")
		fmt.Println("")
		return
	}

	if data.age>=18 {
		fmt.Println("")
		fmt.Println("Yes, you can vote!")
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Printf("Sorry! you can't vote. You can vote after %d years\n",(18-data.age))
	}
}
