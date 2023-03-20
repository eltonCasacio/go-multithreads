package loadbalancer

import (
	"fmt"
)

func worker(workID int, data chan int) {
	for v := range data {
		fmt.Printf("Worker %d received %d\n", workID, v)
	}
}

func RunLoadBalanceExample() {
	data := make(chan int)

	for y := 0; y < 10000; y++ {
		go worker(y, data)
	}

	for i := 0; i < 100000; i++ {
		data <- i
	}

}
