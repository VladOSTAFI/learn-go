package main

import "time"

func runGoroutines() {
	go greet()

	println("Test")

	time.Sleep(time.Second * 5)
}

func greet() {
	println("Hello")
}
