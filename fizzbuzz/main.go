package main

import (
	"fmt"
	"strconv"
)

func main() {
	//variables
	for i := 0; i <= 100; i++ {
		out := ""

		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if out == "" {
			out = strconv.Itoa(i)
		}
		fmt.Println(out)
	}
}
