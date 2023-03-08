package main

import "fmt"

func main() {
	var i = 21
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println("%")
	var j = true
	fmt.Println(j)
	fmt.Println("")
	fmt.Println(10101)
	fmt.Println("?")
	fmt.Println(i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", 15)
	fmt.Println("F")
	var c rune = '\u042F'
	fmt.Printf("%c\n", c)
	// fmt.Println("U+042F ")
	fmt.Println("")
	var k float64 = 123.456
	fmt.Printf("%.3f\n", k)
	fmt.Printf("%.4e\n", k)

}
