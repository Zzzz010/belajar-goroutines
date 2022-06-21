package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func AddtoMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}
	// addToMap := func(value int) {
	// 	data.Store(value, value)
	// }

	for i := 0; i < 100; i++ {
		go AddtoMap(data, i, group)
	}

	group.Wait()

	time.Sleep(3 * time.Second)
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
