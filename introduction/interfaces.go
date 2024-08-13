package main

import "fmt"

func runIntercafes() {
	var file ReaderWriter

	file = File{Name: "example.txt"}

	file.Read()
	file.Write()
}

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
	Name string
}

func (f File) Read() {
	fmt.Println("Reading from file:", f.Name)
}

func (f File) Write() {
	fmt.Println("Writing to file:", f.Name)
}
