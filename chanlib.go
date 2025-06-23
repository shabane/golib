package golib

import (
	"fmt"
)

func Mychn() {
	channel := make(chan int)
	go func() { channel <- 1 }()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	f := <-channel
	s := <-channel
	t := <-channel

	fmt.Println(f, s, t)
}
