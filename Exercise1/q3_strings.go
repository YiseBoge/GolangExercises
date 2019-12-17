package main

import "fmt"
import "unicode/utf8"

func q3One() {
	for i := 0; i < 100; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%s", "A")
		}
		fmt.Println()
	}
}

func q3Two() {
	count := 0
	bytes := 0
	str := "asSASA ddd dsjkdsjs dk"
	for _, i := range str {
		count++
		bytes += utf8.RuneLen(i)
	}
	fmt.Println("The length is:", count)
	fmt.Println("The bytes length is:", bytes)
}

func q3Three() {
	str := "asSASA ddd dsjkdsjs dk"
	runes := []rune(str)
	runes[4] = 'a'
	runes[5] = 'b'
	runes[6] = 'c'

	fmt.Println("The new string is:", string(runes))
}

func q3Four() {
	str := "foobar"
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Println("The reversed string is:", string(runes))
}

func main() {
	//q3One()
	//q3Two()
	//q3Three()
	q3Four()
}
