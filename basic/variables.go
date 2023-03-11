package main

import "fmt"

var x, y, z int = 3, 4, 5
var c, python, java = "hi", "golang", true

func main() {
	c, python, java := true, false, "no!" // := can only be used in function
	fmt.Println(x, y, z, c, python, java)
}