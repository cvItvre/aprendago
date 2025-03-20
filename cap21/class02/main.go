package main

import "fmt"

func main() {

	ch := make(chan int)

	go send(ch, 36)
	receive(ch)

	cs := make(chan<- int)
	cr := make(<-chan int)

	fmt.Println("Channels Types")
	fmt.Printf("ch\t%T\n", ch)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

	fmt.Println("\nAssignement/Convertion")
	fmt.Printf("ch\t%T\n", (chan<- int)(ch))
	fmt.Printf("ch\t%T\n", (<-chan int)(ch))

}

func send(sender chan<- int, num int) {
	sender <- num
}

func receive(receiver <-chan int) {
	fmt.Println("Valor recebido do canal:", <-receiver)
}
