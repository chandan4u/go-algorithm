package main

import "fmt"

func main()  {
	res := factorial(5)
	fmt.Println("Result :", res)


}

func factorial(n int) int {
	fmt.Println("Initialize First :", n)
	if n <= 1 {
		return 1
	}
	val := n* factorial(n-1)
	fmt.Println("Initialize Last :", n)
	return val
}
