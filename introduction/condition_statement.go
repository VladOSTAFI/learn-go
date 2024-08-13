package main

import (
	"fmt"
	"strconv"
)

func runConditionStatement(x int) {
	if x > 20 {
		fmt.Println("bigger then 20")
	} else {
		fmt.Println("less then 20")
	}

	y := 12

	if y > 10 {
		fmt.Println("y is bigger then 10")
	}

	y = 8

	if y < 10 {
		fmt.Println("y is less then 10")
	}

	switch {
	case x == 20, x > 20:
		fmt.Println("x is equal or bigger then 20")
	case x < 10:
		fmt.Println("x is less then 10")
	}

	i := 1
	for i < x {
		iStr := strconv.Itoa(i)
		fmt.Println("i is equeal to " + iStr)

		i = i * 2
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	word := "test me"

	for index, char := range word {
		if char != ' ' {
			fmt.Printf("Index: %d, Value: %c\n", index, char)
		}
	}
}
