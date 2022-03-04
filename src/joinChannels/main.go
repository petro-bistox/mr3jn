package main

import "fmt"

/*
Слить N каналов в один
	- нужно описать функцию, которая смерджит все данные из n каналов в один и вернет его.
*/
func joinChannels(chans ...<-chan int) <-chan int {
	mergedChan := make(chan int)
	// implement joining incomming channels to the merged channel
	return mergedChan
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}

}
