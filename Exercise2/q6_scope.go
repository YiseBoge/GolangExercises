package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}

	//fmt.Printf("%v\n", i)
	// The problem is the fact that it is using the variable 'i' outside of the loop it resides in (it's scope)
}
