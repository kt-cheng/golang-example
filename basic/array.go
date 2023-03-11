package main

import "fmt"

func main() {
    var x [4]float64
    x[0] = 1
    x[1] = 2
    x[2] = 3
    x[3] = 4
	// x := [4]float64{1, 2, 3, 4}

    var total float64 = 0
	for _, value := range x {
		total += value
	}
    fmt.Println(total / float64(len(x)))
}

// slice
// map