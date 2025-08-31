package main

import "fmt"

func main() {
	fmt.Printf("Hello\n")
	TestSomething(42)
}

// TestSomething is a function.
func TestSomething(i int) {
	fmt.Printf("Hello again\n")
	i++

	m := map[string]bool{}
	// What are k, and v?
	for k, v := range m {
		fmt.Printf("k:%v,v:%v", k, v)
	}
}
