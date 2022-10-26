package main

import (
	"fmt"
	"strconv"
)

const Str = "aninija"

func main() {
	var h int32 = 0
	var x int32 = 0
	for _, val := range Str {
		x = 1               // set the first bit
		x = x << (val - 97) // shifting the bit
		fmt.Println("h = ", strconv.FormatInt(int64(h), 2))
		fmt.Println("x = ", strconv.FormatInt(int64(x), 2))

		if (x & h) > 0 { // ANDing the bits of h and x
			fmt.Printf("Duplicate found %c\n", val)
		} else {
			h = h | x
		}
	}
}
