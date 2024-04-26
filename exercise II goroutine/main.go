package main

import (
	"fmt"
)

func sayHello(c chan string) {
	fmt.Println("sayHello started", <-c)
}
func sayHi() {
	fmt.Println("sayHi started")
}

func loopGo(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	fmt.Println("main started")

	// c := make(chan string)

	// go sayHello(c)
	// go sayHi()

	// time.Sleep(1 * time.Second)

	c := make(chan int, 5)

	go loopGo(c)

	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("main ended")

}
