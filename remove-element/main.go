package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
        return arr[1:]
		} else if position == "right" {
        return arr[:len(arr)-1]
		} else {
        return nil
    }
	//yout code here!

}

func main() {
	arr1 := []interface{}{1, 2, 3, 4, 5}
    result1 := removeLeftRight(arr1, "left")
    fmt.Println(result1) // Output: [2 3 4 5]

    arr2 := []interface{}{"a", "b", "c", "d"}
    result2 := removeLeftRight(arr2, "right")
    fmt.Println(result2) // Output: [a b c]
}
