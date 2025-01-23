package main

import "fmt"

// Custom type
type TestValue struct {
	t1 string
	t2 string
}

func main() {
	var value1 int = 10
	var value2 string = "t1"
	var value3 bool = true
	var value4 float32 = 1.1
	var value5 float64 = 1.1
	var value6 *int = &value1 // Memory reference
	value7 := 1               //Inference -> The language engine evaluates the data and has the type at the time of assembly

	var v1, v2, v3 int = 1, 2, 3

	// Array
	var value8 = [2]int{1, 2}
	value9 := [2]int{1, 2}

	value10 := TestValue{"t1", "t2"}

	// Constants
	const value11 int = 10

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3)
	fmt.Println(value4)
	fmt.Println(value5)
	fmt.Println(value6)
	fmt.Println(value7)
	fmt.Println(v1, v2, v3)
	fmt.Println(value8)
	fmt.Println(value9)
	fmt.Println(value10)
	fmt.Println(value11)
}
