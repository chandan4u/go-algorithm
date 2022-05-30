package main

import "fmt"

func seqSumNumber(arr []int, sum int) []int {
	currentSum := arr[0]
	start := 0
	for i:=1; i<=len(arr); i++ {
		fmt.Println(">>>", currentSum)
		for currentSum > sum && start < i - 1 {
			currentSum = currentSum - arr[start]
			start++
		}
		if currentSum == sum {
			return []int{start, i - 1}
		}
		currentSum = currentSum + arr[i]
	}
	return []int{}
}

func main() {
	arr := []int{2, 10, 9, 6, 3, 10}
	sum := 15
	res := seqSumNumber(arr, sum)
	fmt.Println("Result :", res)
}
