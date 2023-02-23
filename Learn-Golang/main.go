package main

import "fmt"

func main() {

	x := foo(5)
	n, s := bar(15, "Hai")
	fmt.Println(x)
	fmt.Println(n)
	fmt.Println(s)
}

func foo(n int) int {
	return n
}

func bar(n int, s string) (int, string) {
	return n, s
}
