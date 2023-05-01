package main

import "fmt"

func howManyElements(data []any) int {
	return len(data)
	//your code here!
}

func main() {
	fmt.Println(howManyElements([]any{1, 2, 3, 4, 5}))
	fmt.Println(howManyElements([]any{"Edo", "Budi", "Joko", "Tono"}))
}
