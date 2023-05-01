package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
		} else if position == "down" {
		return append(data, newData)
		} else {
		return data
	}
}
	//your code here!

func main() {
	arr := []int{1, 2, 3, 4, 5}
	newArr := AddElement(arr, 6, "up")
	fmt.Println(newArr) // [6 1 2 3 4 5]

	newArr = AddElement(arr, 6, "down")
	fmt.Println(newArr) // [1 2 3 4 5 6]

	newArr = AddElement(arr, 6, "invalid")
	fmt.Println(newArr)
}
