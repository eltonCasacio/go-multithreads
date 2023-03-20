package channel

import "fmt"

func RunChannelExample() {
	channel := make(chan string)

	go func() {
		channel <- "Hello Human!"
	}()

	msg := <-channel

	fmt.Println(msg)
}
