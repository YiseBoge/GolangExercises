package main

import "fmt"

func q1One() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func q1Two() {
	j := 0
loop:
	fmt.Println(j)
	j++
	if j < 10 {
		goto loop
	}
}

func q1Three() {
	arr := [10]int{}
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}

func main() {

	q1One()
	q1Two()
	q1Three()

}
