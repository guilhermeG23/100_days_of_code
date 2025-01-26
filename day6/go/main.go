package main

import (
	"fmt"
	"time"
)

func test_action() {
	fmt.Println("Hello World")
}

func test_channel(value int, channel_connection chan int) {
	select {
	case channel_connection <- value:
	default:
	}
}

func main() {
	test_action()
	go test_action()

	time.Sleep(time.Second)
	fmt.Println("---------------------------------")

	// Unbuffered channel
	channel_connection := make(chan int)

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range values {
		go test_channel(number, channel_connection) // Feed the channel
		go test_channel(number, channel_connection) // Trying to feed the channel, but it already has value
		show_value := <-channel_connection          // Empty the channel
		fmt.Println(show_value)
	}

	time.Sleep(time.Second)
	fmt.Println("---------------------------------")

	// Buffered channel
	channel_connection1 := make(chan int, 2)

	channel_connection1 <- 1
	channel_connection1 <- 2

	select {
	case channel_connection1 <- 3: // Attempt to populate an already full channel
	default:
	}

	fmt.Println(<-channel_connection1)
	fmt.Println(<-channel_connection1)

	// As the channel accepts multiple inputs,
	// for will present the problem of feeding repetitions,
	// because, it will feed both channels with the same value and prevent the third,
	// but even so the for continues, giving the disparity
	for _, number := range values {
		go test_channel(number, channel_connection1)
		go test_channel(number, channel_connection1)
		go test_channel(number, channel_connection1)
		fmt.Println(<-channel_connection1)
	}

}
