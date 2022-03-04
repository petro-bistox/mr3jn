package main

import "fmt"

/*
	WaitGroup на симофоре
		- нужно описать симофор который будет ждать выполнения пяти горутин
*/

type sema chan struct{}

func New(n int) sema {
	return make(sema, n)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		go func(n int) {
			// added to our sema
			fmt.Println(n)
		}(num)
	}

	// wait our sema
}
