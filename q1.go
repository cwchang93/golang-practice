package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hi")
	// fmt.Print(q1())
	// num := 123
	// var numofzero = 5 - len(strconv.Itoa(num))
	// fmt.Printf("0"*numofzero + %d\n", 123)
	// fmt.Printf("0" * numofzero)
	// fmt.Printf("%05d", 12)
}

func q1() string {
	total := 100000
	num := 12

	// var lenofNum = len(strconv.Itoa(num))
	var final = strconv.Itoa(total + num)

	// println(len(strconv.Itoa(num)))
	// strconv.Itoa(final)
	return final[1:]

}
