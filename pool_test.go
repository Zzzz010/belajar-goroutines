package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	//var pool sync.Pool
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	// pool.Put("Radit")
	// pool.Put("Bagus")
	// pool.Put("Putra")
	pool.Put("Raung")
	pool.Put("Kawijayan")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Done!")
}
