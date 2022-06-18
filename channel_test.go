package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channels := make(chan string)
	defer close(channels)

	go func() {
		time.Sleep(2 * time.Second)
		channels <- "Raung Kawijayan"
	}()

	data := <-channels
	fmt.Println(data)
	time.Sleep(1 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Raditya Bagus Putra"
}

func TestParameterChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}
