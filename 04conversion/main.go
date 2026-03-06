package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out pizza app")

	fmt.Println("Please rate out pizza")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("Thanks for rating : ", input)
	}

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("error is ", err)
	} else {
		fmt.Println("Added 1 to rating : ", numRating+1)
	}

}
