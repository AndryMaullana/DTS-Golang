package main

import (
	"fmt"
)

func main() {
	var number = 7
	switch number {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("awesome")
	case 6:
		fmt.Println("worry")
	default:
		fmt.Println("not bad")
	}

	// Switch with relational operators
	var score int32 = 6
	switch {
	case score == 6:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	case score < 5:
		fmt.Println("It is oke, but study harder")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You don't have a good score yet")
		}
	}
	// Switch ( Fallthrough Keyword)
	fmt.Println("=====================")
	switch {
	case score == 6:
		fmt.Println("perfect")

	case (score < 8) && (score > 3):
		fmt.Println("not bad")
		fallthrough // kata kunci ini akan melanjutkan kode di bawah walaupun kondisinya bernilai true
	case score < 5:
		fmt.Println("It is oke, but study harder")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("You don't have a good score yet")
		}
	}
	// Switch dan if digabung menjadi kondisi bersarang
	fmt.Println("=================================")
	var score1 int = 0
	if score1 > 7 {
		switch score1 {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score1 == 5 {
			fmt.Println("Not Bad")
		} else if score1 == 3 {
			fmt.Println("Keep Trying")
		} else {
			fmt.Println("You can do it")
			if score1 == 0 {
				fmt.Println("try harder")
			}
		}
	}
}
