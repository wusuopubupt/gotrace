package main

import "fmt"

const N = 10

func main() {
	c := make(chan int, 4) // change n=2 to n=1 will cause deallock error , but not n=3
	go func() {
		for i := 0; i < N; i++ {
			c <- i
		}
	}()
	for i := 0; i < N; i++ {
		fmt.Println(<-c)
	}
}
