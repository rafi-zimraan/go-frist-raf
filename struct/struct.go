package main

import "fmt"

type User struct {
	name  string
	email string
	phone int
	age   int
}

func main() {
	user := User{
		name:  "rafi zimraan arjuna wijaya",
		email: "zimraan27@gmail.com",
		phone: 620987654321,
		age:   19,
	}
	fmt.Println(user)
}
