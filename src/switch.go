package main

import "fmt"

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

	// Switch dengan perbandingan
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
}
