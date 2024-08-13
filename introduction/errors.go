package main

import (
	"errors"
	"fmt"
)

// func runErrors() {
// 	content, err := readFile("errors_2.go")

// 	if errors.Is(err, os.ErrNotExist) {
// 		println("File does not exist")
// 		return
// 	} else if err != nil {
// 		println("Error: ", err.Error())
// 		return
// 	}

// 	println("File content: ", content)
// }

// func readFile(path string) (*os.File, error) {
// 	value, error := os.Open(path)

// 	if error != nil {
// 		return nil, fmt.Errorf("Failed to open file %s: %w", path, error)
// 	}

// 	return value, nil
// }

var ErrDivisionByZero = errors.New("division by zero")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func runErrors() {
	result, err := divide(10, 0)
	if errors.Is(err, ErrDivisionByZero) {
		fmt.Println("Cannot divide by zero.")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
