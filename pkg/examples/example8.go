package examples

func snd8A() chan int {
	x := make(chan int)
	go snd8B(x)
	return x
}

func snd8B(ch chan int) {
	ch <- 42
}

func snd8C() chan int {
	x := make(chan int)
	go rcv8D(x)
	return x
}

func rcv8D(ch chan int) {
	<-ch
}

func ExampleFunc8() {

	<-snd8A()
	snd8C() <- 43
}
