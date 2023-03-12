package main

import (
    "fmt"
    "math/rand"
    "time"
    "average"
)

// in '/opt/homebrew/Cellar/go/1.19.1/libexec/src/average/main.go'
// package average 

// func Average(xs []float64) float64 {
//         total := float64(0)
//         for _, x := range xs {
//                 total += x
//         }
//         return total / float64(len(xs))
// }

// in '/opt/homebrew/Cellar/go/1.19.1/libexec/src/average/average_test.go'
// package average

// import "testing"

// func TestAverage(t *testing.T) {
//         var v int
//         v = Average([]int{1, 3})
//         if v != 2 {
//                 t.Error("Expected 2, got", v)
//         }
// }

func main() {
    //Provide seed
    rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	arr := rand.Perm(5) // func Perm(n int) []int
	fmt.Println(arr)

	fmt.Println(average.Average(arr))
}