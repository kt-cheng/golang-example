package main

import ("fmt" ; "container/list") // 利用 ; 讓兩個 library 寫在同一行

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

// heap