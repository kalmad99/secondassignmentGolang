package main

import "fmt"

var first = []string{"a", "b", "c"}
var second = []string{"1", "2", "3"}

func main() {
	for i := range first {
		fmt.Printf("%v, %v, ", first[i], second[i])
	}
}
