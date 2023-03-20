package channelbuffer

func RunChannelWithBufferExample() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "Human!"

	println(<-c)
	println(<-c)
}
