package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i ++
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string {
	result := strings[:0]

	for _, s := range strings {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}
