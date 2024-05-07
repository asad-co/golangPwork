package main

import "fmt"

func main() {
	a := 5
	fmt.Println(&a)
	fmt.Println(a)

	b := &a
	fmt.Println(*b)
	fmt.Println(b)

	*b = 1
	fmt.Println(a)
	changeValue(&a)

	fmt.Println(a)
}

func changeValue(a *int) {
	*a = 100
	return
}
