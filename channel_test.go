package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
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
	channel <- "Kereta Api Argo Wilis"
}

func GiveMeResponse2(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Kereta Api Argo Wilis"
	channel <- "Kereta Api Turangga"
}

func TestParameterChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Kereta Api Bangunkarta"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}

func TestBuffer(t *testing.T) {
	channel := make(chan string, 4)
	defer close(channel)

	go func() {
		channel <- "Bima"
		channel <- "Arjuna"
		channel <- "Nakula"
		channel <- "Sadhewa"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("Done!")

	time.Sleep(3 * time.Second)
	// data := <-channel
	// fmt.Println(data)
	// fmt.Println(data)
	// fmt.Println(data)
	// fmt.Println(data)

	// go func(m *chan string) {
	// 	fmt.Println("Entering the goroutine...")
	// 	for {
	// 		fmt.Println(<-*m)
	// 	}
	// }(&channel)
	// time.Sleep(5 * time.Second)

}

func TestRangeChannel(t *testing.T) {
	channels := make(chan string)

	go func() {
		for i := 0; i < 15; i++ {
			channels <- "Perulangan Ke " + strconv.Itoa(i)
		}
		close(channels)
	}()

	for data := range channels {
		fmt.Println("Data diterima", data)
	}
	fmt.Println("Done")
}
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data pertama dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data pertama dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string, 2)
	channel2 := make(chan string, 2)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse2(channel1)
	go GiveMeResponse2(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data pertama dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data pertama dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 4 {
			break
		}
	}
}
