package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	} // Condition loop

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	} // Classic for loop of initialize condition after

	for {
		fmt.Println("Infinite loop until break statement")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // Continue will skip the code that comes after
		}
		fmt.Println(n)
	}
}
