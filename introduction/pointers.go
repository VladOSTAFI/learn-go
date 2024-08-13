package main

func updateName(n *string) {
	*n = "Joe"
}

func runPointers() {
	// name := "Kate"

	// m := &name // m is stored reference to the value of the variable name

	// fmt.Println("M:", m)

	// fmt.Println("Memory address m:", &m)
	// fmt.Println("Value of m:", *m) // to dereference value on variabel m use asterisk

	// fmt.Println("Value of name:", name)

	// name := "Kate"

	// m := &name

	// fmt.Println("Value of name:", name)
	// fmt.Println("Value of m:", *m)

	// name = "John"

	// fmt.Println("Value of name(2):", name)
	// fmt.Println("Value of m(2):", *m)

	// *m = "Bob"

	// fmt.Println("Value of name(3):", name)
	// fmt.Println("Value of m(3):", *m)

	// name := "Kate"

	// fmt.Println("Value of name:", name)

	// updateName(&name)

	// fmt.Println("Value of name(2):", name)

}
