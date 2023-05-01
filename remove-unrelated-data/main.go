package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
	//your code here
}

func main() {
	dataMap := map[string]interface{}{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	key := "age"

	result := removeUnrelated(dataMap, key)
	fmt.Printf("%v", result)
}
