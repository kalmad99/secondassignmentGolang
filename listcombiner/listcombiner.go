package main

import "fmt"

var pow = []string{"a", "b", "c"}

func main() {
	for i, v := range pow {
		fmt.Printf("%s, %d, ", v, i+1)
	}
}
