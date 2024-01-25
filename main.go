package main

import "fmt"

func main() {
	var mySlice []string
	mySlice = append(mySlice, "Hello", "World", "from", "Golang", "slice", "!")
	fmt.Println(mySlice)
}
