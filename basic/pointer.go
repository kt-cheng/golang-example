package main

import "fmt"

// 1
// func zero(x int) {
// 	x = 0
// }
// func main() {
// 	x := 5
// 	zero(x)
// 	fmt.Println(x) // x is still 5
// }

// 2
// func zero(x int) int {
// 	x = 0
// 	return x
// }
// func main() {
// 	x := 5
// 	x = zero(x)
// 	fmt.Println(x) // x is 0
// }

// 3
func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}
