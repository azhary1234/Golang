package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(x...))
}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		fmt.Print("Amount :", v, "+", total, " =")
		total += v
		fmt.Println(total)
	}

	return total
}
