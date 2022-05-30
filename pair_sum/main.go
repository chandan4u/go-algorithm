package main

import "fmt"

func main() {
	arr := []int{10, 2, -2, -20, 10}
	sum := -22
	// desiApproach(arr, sum)
// 	hashApproach(arr, sum)

	seqSumNum(arr, sum)
}

func seqSumNum(arr []int, sum int) {
	currentSum := 0
	start := 0
	end := -1
	hashTable := make(map[int]int)
	for i:=0; i<len(arr); i++ {
		currentSum = currentSum + arr[i]
		if currentSum - sum == 0 {
			fmt.Println("<>>>>>>>>>>")
			start = 0
			end = i
			break
		}
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>", hashTable)
		fmt.Println(">>>>>>>>>>", hashTable[currentSum - sum], currentSum)
		_, ok := hashTable[currentSum - sum]
		if ok {
			start = hashTable[currentSum - sum] + 1
			end = i
			break
		}
		hashTable[currentSum] = i
	}
	if end == -1 {
		fmt.Println("Not Found")
	}
	fmt.Println("Start : ", start, "End : ", end)
}


func desiApproach(arr []int, sum int) {
	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			if arr[i] + arr[j] == sum {
				fmt.Printf("Pair with the given sum %d - %d [ %d, %d]\n",arr[i], arr[j], i, j)
			}
		}
	}
}

func hashApproach(arr []int, sum int) {
	hashTable := make(map[int]int)
	for i:=0; i<len(arr); i++ {
		temp := sum - arr[i]
		_, ok := hashTable[temp]
		if ok {
			fmt.Printf("Pair with the given sum %d - %d [%d, %d]\n",temp, arr[i], hashTable[temp], i)
		}
		hashTable[arr[i]] = i
	}
}


