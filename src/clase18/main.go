package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)
	go say("World", &wg)
	wg.Wait()
	go func(text string) {
		fmt.Println(text)
	}("Adiós")
	// Basic option but not recommended
	time.Sleep(time.Second * 1)

	// Channels

}
