package main

func removeChar(str string) (res string, numberOfRemovedChars int) {
	for _, char := range str {
		if char != ' ' {
			res = res + string(char)
		} else {
			numberOfRemovedChars++
		}
	}

	return
}

func sumNumbers(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
