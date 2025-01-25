package main

import "fmt"

func test_action() {
	fmt.Println("Hello World")
}

func test_action1() int {
	return 1
}

func test_action2(value int) int {
	return value
}

func main() {
	test_action()
	var value int = test_action1()
	fmt.Println(value)
	var value1 int = test_action2(1)
	fmt.Println(value1)
}
