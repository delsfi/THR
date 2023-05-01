package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	userData := strings.Split(data, ",")
	age, _ := strconv.Atoi(userData[1])
	return User{
		Name:    userData[0],
		Age:     age,
		Address: userData[2],
	}
	//your code here!

}

func main() {
	data := "Edo,20,Jakarta"
	user := ConvertData(data)
	fmt.Printf("%+v\n", user) // {Name:Edo Age:20 Address:Jakarta}

	data = "Budi,30,Bandung"
	user = ConvertData(data)
	fmt.Printf("%+v\n", user) // {Name:Budi Age:30 Address:Bandung}
}
