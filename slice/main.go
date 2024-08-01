package main

import "fmt"

func main() {
	arr := make([]int64, 5)
	for i := 0; i < len(arr); i++ {
		arr[i] = int64(i)
	}
	//fmt.Println(arr)
	Delete(2, arr)
	fmt.Println(arr)
}
