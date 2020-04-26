// SPDX-License-Identifier: MIT
package examples

import "fmt"

func snd3(ch chan int) {

	var x int = 42
	fmt.Printf("send %d \n", x)
	ch <- x
}

func rcv3(ch chan int) {

	var x int
	x = 1
	switch x {
	case 1:
		y := ch
		fmt.Printf("case %v", y)
	case 2:
		ch <- -x
	default:
		ch <- -x
		fmt.Printf("default")
	}

	fmt.Printf("received %d \n", x)
}

func ExampleFunc3() {

	var ch chan int = make(chan int)
	go snd3(ch)
	rcv3(ch)

}
