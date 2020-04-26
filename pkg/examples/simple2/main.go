package main

func main() {

	var t chan bool = make(chan bool)
	var f chan bool = make(chan bool)

	var x bool = true
	var y bool = false

	go func() {
		for {
			t <- x
		}
	}()

	go func() {
		for x && !y {
			t <- true
		}
	}()

	if x {
		x = y
	} else {
		y = <-t
	}

	select {

	case x = <-t:
		x = y
	case y = <-f:
		y = x
	}
}
