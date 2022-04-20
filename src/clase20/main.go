package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	//First channel that handles 2 data at the time
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	// Range and close
	//It's always best practice to close the channel if you know that it wont receive more data
	close(c)

	// c <- "Mensaje3"

	for message := range c {
		fmt.Println(message)
	}

	// Select (Dynamic channels)
	// To monitor the channels
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje1", email1)
	go message("Mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
