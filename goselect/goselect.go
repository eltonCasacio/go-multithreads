package goselect

import (
	"fmt"
	"math/rand"
	"time"
)

func RunSelectExample() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		c2 <- 2
	}()

	select {
	case msg := <-c1:
		fmt.Printf("Mensagem %d chegou primeiro", msg)
	case msg := <-c2:
		fmt.Printf("Mensagem %d chegou primeiro", msg)

	}
}

func RunSelectExample2() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Printf("Mensagem %d chegou primeiro\n", msg)
		case msg := <-c2:
			fmt.Printf("Mensagem %d chegou primeiro\n", msg)
		}
	}
}

func RunSelectExample3() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 5)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 4)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		fmt.Printf("Mensagem %d chegou primeiro\n", msg1)
	case msg2 := <-c2:
		fmt.Printf("Mensagem %d chegou primeiro\n", msg2)
	case <-time.After(time.Second * 3):
		println("timeout")
	default:
		println("default")
	}
}

func RunSelectExample4() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Nanosecond)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Nanosecond)
		c2 <- 2
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("Mensagem %d chegou primeiro\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("Mensagem %d chegou primeiro\n", msg2)
		case <-time.After(time.Nanosecond):
			println("timeout")
		default:
			println("default")
		}
	}

}

func RunSelectExampleInfinityLoop() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c1 <- 1
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			c2 <- 2
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("Mensagem %d chegou primeiro\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("Mensagem %d chegou primeiro\n", msg2)
		case <-time.After(time.Second * 3):
			println("timeout")
		}
	}

}
