package channeldirections

import (
	"fmt"
	"sync"
)

var numbers = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func receveOnlyChannel(number int64, c chan<- int64) {
	c <- number
}

func sendOnlyChannel(c <-chan int64, wg *sync.WaitGroup) {
	fmt.Println(<-c)
	wg.Done()
}

func RunChannelDirectionsExample() {
	c := make(chan int64)
	wg := &sync.WaitGroup{}
	wg.Add(10)

	go func() {
		for _, v := range numbers {
			receveOnlyChannel(v, c)
		}
	}()

	go func() {
		for {
			sendOnlyChannel(c, wg)
		}
	}()

	wg.Wait()
}
