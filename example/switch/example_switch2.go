package example

func ExampleFuncSwitch_Snd_Rcv2() {

	var ch chan bool = make(chan bool)

	// SwitchStmt: Init, Tag, Body
	switch x := <-ch; x && <-x {

	case <-ch, <-ch:
		<-ch
	}
	return
}
