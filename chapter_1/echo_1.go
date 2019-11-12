package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, step string

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
		// s += step + os.Args[i]
		// step = " "
	}
	// fmt.Println(s)
}
