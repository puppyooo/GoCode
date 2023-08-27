package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less []int
	var greater []int
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
func main() {
	fmt.Println("Hello, World!")
	arr := []int{10, 5, 2, 3}
	fmt.Println(arr)
	fmt.Println(quickSort(arr))
	fmt.Println("hello world")
}
