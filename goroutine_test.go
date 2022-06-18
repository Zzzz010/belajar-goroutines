package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloRaung() {
	fmt.Println("Hello Raung!")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloRaung()
	fmt.Println("Hello")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
