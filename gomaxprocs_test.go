package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	groups := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		groups.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			groups.Done()
		}()
	}

	TotalCPU := runtime.NumCPU()
	fmt.Println("CPU ", TotalCPU)

	TotalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread ", TotalThread)

	TotalGourotine := runtime.NumGoroutine()
	fmt.Println("Goroutine ", TotalGourotine)

	groups.Wait()
}

func TestChangeGomaxprocs(t *testing.T) {
	groups := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		groups.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			groups.Done()
		}()
	}

	TotalCPU := runtime.NumCPU()
	fmt.Println("CPU ", TotalCPU)

	runtime.GOMAXPROCS(17)
	TotalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread ", TotalThread)

	TotalGourotine := runtime.NumGoroutine()
	fmt.Println("Goroutine ", TotalGourotine)

	groups.Wait()
}
