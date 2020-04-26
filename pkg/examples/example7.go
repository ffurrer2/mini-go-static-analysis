// SPDX-License-Identifier: MIT
package examples

import "fmt"

func snd7(ch chan int) {
	ch <- 42
}

func ExampleFunc7() {

	x := make(chan int)
	go snd7(x)
	y := <-x + 1

	fmt.Println(y)
}
