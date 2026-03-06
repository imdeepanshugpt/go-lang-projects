package main

import "fmt"

type User struct {
	Name string
}

func (u *User) changeName() {
	u.Name = "Name changed"
}

func (u User) Greet() string {
	return "Hello " + u.Name
}

func main() {
	usr := User{Name: "Deepanshu"}
	fmt.Println(usr.Greet())
	usr.changeName()
	fmt.Println(usr.Greet())

}
