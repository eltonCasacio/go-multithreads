package channelrange

import (
	"fmt"
	"sync"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func RunChannelExample() {
	c := make(chan int64)

	go func() {
		for _, v := range numbers {
			c <- int64(v)
		}
		close(c)
	}()

	func() {
		for x := range c {
			fmt.Println(x)
		}
	}()
}

func RunChannelExampleWithWaitGroup() {
	c := make(chan int64)
	wg := &sync.WaitGroup{}
	wg.Add(9)

	go func() {
		for _, v := range numbers {
			c <- int64(v)
		}
	}()

	go func() {
		for x := range c {
			fmt.Println(x)
			wg.Done()
		}
	}()

	wg.Wait()
}
