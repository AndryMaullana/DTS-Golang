package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%#v\n", numbers)
	fmt.Println(strings.Repeat("#", 25))
	var stringss = [3]string{"Andri", "Dennis", "Kunto"}
	fmt.Printf("%#v", stringss)

	// Memodifikasi array
	fmt.Println()
	var fruits = [4]string{"Apel", "Mangga", "Semangka", "anggut"}
	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[2] = "Watermelon"
	fruits[3] = "grape"
	fmt.Printf("%#v\n", fruits)
	fmt.Println(strings.Repeat("#", 25))
	for i, v := range fruits {
		fmt.Printf("indeks %d, value %s\n", i, v)
	}
	fmt.Println(strings.Repeat("#", 25))
	for i := 0; i < len(fruits); {
		fmt.Printf("Indeks %d, value %s\n", i, fruits[i])
		i++
	}
	fmt.Println(strings.Repeat("#", 25))
	// array multidimensi
	var balance = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for _, arr := range balance {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
