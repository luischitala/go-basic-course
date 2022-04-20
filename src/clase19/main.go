//Channels in go
package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)
	fmt.Println("Hello")

	go say("Bye", c)

	//Extract the data from the channel

	fmt.Println(<-c)
}
