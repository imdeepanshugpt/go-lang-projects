package main

import "fmt"

const JWTTOKEN string = "ABNCDERTA"

func main() {
	fmt.Println("Variables")
	var userNames string = "deepanshu"
	fmt.Println(userNames)

	var number int = 1234
	fmt.Printf("number is %T \n", number)

	var smallFloat float32 = 1234.12312234134
	fmt.Println(smallFloat)
	fmt.Printf("number is %T \n", smallFloat)

	nostyle := "no style declaration"
	fmt.Println(nostyle)
	fmt.Printf("no style is of %T", nostyle)
}
