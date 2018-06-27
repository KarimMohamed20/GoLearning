package main

import "fmt"

func main() {
	a := 15
	b := &a // Var (a) Memory Read Address .

	fmt.Println(b)
	fmt.Println(*b)
}
