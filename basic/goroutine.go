package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go f(0)
	time.Sleep(time.Second * 1)
}
