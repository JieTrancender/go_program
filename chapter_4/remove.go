package main

import "fmt"

func remove(ints []int, i int) []int {
	ints[i] = ints[len(ints)-1]
	return ints[:len(ints)-1]
}

func remove2(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8}
	fmt.Println(remove(s, 2))
}
