package main

import (
	"fmt"
	"strconv"
	"strings"
)

func runDataTypes() {
	str := "ssss"
	intValue := 24

	intValueStr := strconv.Itoa(intValue)

	fmt.Println("value:", str+intValueStr)
}

func runArrays() {
	var a [5]int
	a[0] = 10
	fmt.Println(a)

	b := [4]string{"This", "is", "my", "array"}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)

	for index, value := range b {
		fmt.Println(index, strings.ToUpper(value))
	}
}

func runSlices() {
	// var s []int
	// s = append(s, 1, 2, 3)
	// fmt.Println(s)

	// t := []string{"Hello", "World"}
	// fmt.Println(t)

	// u := make([]int, 5)
	// fmt.Println(u)

	// v := make([]int, 3, 5)
	// fmt.Println(v)

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]

	fmt.Println("Sliced array: ", slice)

	slice[0] = 10
	fmt.Println("Updated array by reference in slice: ", arr)

	var arr2 [3]int

	for i, value := range slice {
		arr2[i] = value
	}

	fmt.Println("Created new array from slice: ", arr2)
}

func runMaps() {
	users := []map[string]string{
		{"name": "John"},
	}

	users = append(users, map[string]string{"name": "Chris"})

	users = append(users, map[string]string{"name": "Ivan", "age": "30"})

	fmt.Println("Users: ", users)

	for _, userData := range users {
		name, nameExists := userData["name"]
		age, ageExists := userData["age"]

		switch {
		case nameExists && ageExists:
			fmt.Println("Name is: ", name, "and age is:", age)
		case nameExists:
			fmt.Println("Name is: ", name)
		case ageExists:
			fmt.Println("Age is: ", age)
		}
	}

	productPrices := map[string]int{"table": 100, "computer": 2431, "keyboard": 302}
	totalPrice := 0

	for _, price := range productPrices {
		totalPrice += price
	}

	fmt.Println("Total price is: ", totalPrice)
}

func runStruct() {
	// var users []User

	// names := []string{"Joe", "Tommy", "Jerremy", "David"}
	// surnames := []string{"Don", "Bon", "Konm", "Bris"}

	// for index, name := range names {
	// 	user := User{
	// 		Name:     name,
	// 		Surname:  surnames[index],
	// 		Age:      rand.Intn(20),
	// 		IsActive: index%2 == 0,
	// 	}

	// 	users = append(users, user)
	// }

	// var activeUsers []User

	// for _, user := range users {
	// 	if user.IsActive {
	// 		activeUsers = append(activeUsers, user)
	// 	}
	// }

	user := User{
		Name:     "Joe",
		Surname:  "Don",
		Age:      22,
		IsActive: true,
	}

	user.PrintFullName()

	user.Rename("Ben", "Who")

	user.PrintFullName()
}

type User struct {
	Name     string
	Surname  string
	Age      int
	IsActive bool
}

func (u User) PrintFullName() {
	fmt.Printf("My name is %s %s\n", u.Name, u.Surname)
}

func (u *User) Rename(name string, surname string) {
	u.Name = name
	u.Surname = surname
}
