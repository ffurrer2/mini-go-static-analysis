// SPDX-License-Identifier: MIT
package examples

import "fmt"
import "time"

func snd2(ch chan int) {

	var x int = 0
	for i := 0; i < 10; i++ {
		x++
		ch <- x
		time.Sleep(1 * 1e9)
	}
}

func rcv2(ch chan int) {

	var x int
	for i := 0; i < 10; i++ {
		x = <-ch
		fmt.Printf("received %d \n", x)
	}
}

func ExampleFunc2() {

	var ch chan int = make(chan int)
	go snd2(ch)
	rcv2(ch)

	for i := 0; i < 10; i++ {
		fmt.Printf("Hello World!\n")
	}

	array := []int{1, 2, 3}
	for range array {
		fmt.Printf("1 Hello World!\n")
	}

	for _, value := range array {
		fmt.Printf("2 Hello World %d!\n", value)
	}

	for i, _ := range array {
		fmt.Printf("3 Hello World %d!\n", i)
	}

	for i, value := range array {
		fmt.Printf("4 Hello World %d, %d!\n", i, value)
	}

	chanArray := []chan int{make(chan int), make(chan int), make(chan int)}
	for range chanArray {
		fmt.Printf("1 Hello World!\n")
	}

	for _, value := range chanArray {
		fmt.Printf("2 Hello World %v!\n", value)
		value <- 42
	}

	for i, _ := range chanArray {
		fmt.Printf("3 Hello World %d!\n", i)
	}

	for i, value := range chanArray {
		fmt.Printf("4 Hello World %d, %v!\n", i, value)
		<-value
	}

	if true {
		<-chanArray[0]
	} else if 1 == 2 {
		fmt.Printf("Hello World!\n")
	} else {
		fmt.Printf("Hello World!\n")
	}

	if true {
		fmt.Printf("Hello World!\n")
	} else if 1 == 2 {
		<-chanArray[0]
	} else {
		fmt.Printf("Hello World!\n")
	}
}

func ExampleFunc2A() int {
	return exampleFunc2B()
}

func exampleFunc2B() int {
	a := 41 + 1
	return a
}
