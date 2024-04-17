package main

import "fmt"

// global variable -- Pascal case: first letter capital
var Val int = 55

// package variable -- Camel case: first letter small
var val int = 34

func main() {
	/*
		//fist method
		var a int
		a = 55

		//second method
		var b int = 50
		var c = 45

		//third method
		d := 40

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	*/

	/*
		types of variables:
		1: Local Variables: inside declare of function
		2: Global Variables: declare outside of functions but only accessable to within package
		3: Package Variables: also outside of functions, accessible on all other packages
	*/
	// val := 55
	// fmt.Printf("%v, %T", val, val)

	//type conversion

	// var a int32 = 100

	// var b int64 = int64(a)

	// c := 30

	// fmt.Printf("%v, %T", b, b)
	// fmt.Printf("%v, %T", c, c)

	// c := 2.5
	// var d float64 = c
	// fmt.Printf("%v,%T", d, d)

	// var str int = 111
	// var d string = strconv.Itoa(str)

	// fmt.Printf("%v , %T", d, d)
	/*
		primitive data types:
		int: data is int
		int8: data is int with 8 bits ie -128 to 127
		uint: all positive number, negative not allowed
		complex64/complex128: used to store complex var like 2+5i
	*/
	// var a int8 = -128
	// fmt.Println(a)

	// c := complex(2, 5)
	// fmt.Printf("%v,%T", imag(c), c)

	//const type is for values that don't change
	// const i int = iota
	// const a int = 0
	// fmt.Printf("%v,%T \n", i, i)
	// fmt.Printf("%v,%T", a, a)

	//iota leads to increment of value as compare to previous one and it must be defined in group to make it work:

	// const (
	// 	x = iota
	// 	_
	// 	y
	// 	z
	// )
	// fmt.Println(x, y, z)
	/*
		//array
		arr := [...]int{2, 3, 4, 5, 6, 4, 3}
		fmt.Println(arr[1:])

		//double dimension array
		var matrix [3][3]int = [3][3]int{{2, 3, 4}, {5, 6, 4}, {3, 5, 7}}
		fmt.Println(matrix[:2][1])
		fmt.Println(matrix[1][:2])
		fmt.Println(matrix[:][0:])
	*/
	/*
		//map
		empSal := make(map[string]int) //declaration
		empSal = map[string]int{ //initialization
			"Amna":  200,
			"Saima": 100,
			"Wajid": 500,
		}

		empSal2 := map[string]int{ //declaration + initialization
			"Saima": 100,
			"Wajid": 500,
			"Amna":  200,
		}
		emp := empSal2
		empSal2["Asad"] = 400
		delete(empSal2, "Saima")
		fmt.Println(emp)

		zeeshi, ok := empSal2["Zeeshi"]
		fmt.Println(zeeshi, ok)

		_, flag := empSal2["Zeeshi"]
		fmt.Println(flag)
	*/

	//slicing it refers to certain memory point and donot copy the array elements just like map
	var arr []int = []int{1, 2, 3, 4}
	var arr2 = append(arr, 10000)
	arr[2] = 5
	fmt.Println(arr2)
	// var arr2 []int = make([]int, 2, 4)
	// fmt.Println(cap(arr2))
	// fmt.Println(len(arr2))

}
