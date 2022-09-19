package channel

import "fmt"

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func recieveOnly(c <-chan int) {
	fmt.Printf("Received: %d\n", <-c)
}

func sendOnly(c chan<- int) {
	c <- 3 // OK
}

func run() {

	myChan := make(chan int)

	// Can receive and send
	go receiveAndSend(myChan)
	myChan <- 1
	fmt.Printf("Value from receiveAndSend: %d\n", <-myChan)

	// Only receive
	go recieveOnly(myChan)

	// Only send
	go sendOnly(myChan)
	fmt.Printf("Value from sendOnly: %d\n", <-myChan)

}
