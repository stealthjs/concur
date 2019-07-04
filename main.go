package main

import (
	"fmt"
	"runtime"
)

func main() {
	jobs := make(chan int, 50)
	results := make(chan int, 50)

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(jobs, results)
	}

	for i := 0; i < 45; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 45; j++ {
		fmt.Println(<-results)
	}
	fmt.Println("Finished")
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
