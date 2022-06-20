package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronously(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hallo Raung")
	time.Sleep(3 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronously(group)
	}

	group.Wait()
	fmt.Println("Done!")
}
