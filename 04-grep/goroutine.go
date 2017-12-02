package main

import (
	"fmt"
	"time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
    fmt.Println("Wilkommen!")
    fmt.Println("Time for some threading fun!\n")

    fmt.Println("### Without threading:\n")
    time.Sleep(500*time.Millisecond)
	say("world")
	say("hello")

    time.Sleep(1*time.Second)

    fmt.Println("\n### With threading:\n")
    time.Sleep(500*time.Millisecond)
	go say("world")
	say("hello")
}
