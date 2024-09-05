package main

import "fmt"

func main() { 

	var (
		name string
		surname string
	)
	fmt.Scan(&name)
	fmt.Scan(&surname)
	fmt.Println(name, surname)
}