package main

import (
	"fmt"
	"time"
)

func count(ch chan int) {
	for num := range ch {
		result := num * num
		fmt.Printf("Квадрат числа %d: %d\n", num, result)
	}
}

func main() {
	ch := make(chan int)

	go count(ch)

	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)

	time.Sleep(1 * time.Second)
}
