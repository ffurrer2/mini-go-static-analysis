package example

func main() {

  var ch chan bool = make(chan bool)

  for <-ch; <-ch; <-ch {
    <-ch
  }
}
