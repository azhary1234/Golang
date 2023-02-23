package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)

	foo()

	y := incremetor()()

	fmt.Println(y)
	x := foo(5)
	n, s := bar(15, "Hai")
	fmt.Println(x)
	fmt.Println(n)
	fmt.Println(s)
}

func foo(n int) int {
	fmt.Println(x)
	return n
}

func bar(n int, s string) (int, string) {
	return n, s
}

func incremetor() func() int {
	return func() int {
		return 1
	}
}
