package singleton

import "fmt"

func Execute() {
	for i := 0; i < 10; i++ {
		// using normal goroutines lock unlock
		go getInstace()

		// using sync do Once
		go getInstaceOnce()
	}

	fmt.Scanln()
}
