package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	fmt.Println(sum(2, 4))

	fmt.Println(calculation(penjumlahan, 2, 3, 5))

	fmt.Println(calculation(pengurangan, 2, 3, 5))

}

type calc func(x ...int) int

func penjumlahan(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func pengurangan(x ...int) int {
	total := 0

	for _, v := range x {
		total -= v
	}

	return total
}

func calculation(f calc, x ...int) float64 {

	return float64(f(x...))

}

func sum(x ...int) int {

	total := 0

	for _, v := range x {
		total += v
	}

	return total
}
