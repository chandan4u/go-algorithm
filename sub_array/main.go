package main

import "fmt"

var resCount [5]int

func countSubarrays(arr []int) {
	for i:=0; i<len(arr); i++ {
		resCount[i] = 1
		countArr(arr, i, i)
	}
}

func countArr(arr []int, item int, itemIndex int) {
	if  itemIndex == len(arr) - 1 || itemIndex == 0 {
		return
	}
	if itemIndex + 1 < len(arr) - 1 && arr[item] > arr[itemIndex + 1] {
		fmt.Println(">>>>>>>>> + ", arr[item], arr[itemIndex + 1])
		countArr(arr, item, itemIndex + 1)
		resCount[item] = resCount[item] + 1
	}
	if itemIndex - 1 > 0 && arr[item] > arr[itemIndex - 1] {
		fmt.Println(">>>>>>>>> - ", arr[item], arr[itemIndex - 1])
		countArr(arr, item, itemIndex - 1)
		resCount[item] = resCount[item] + 1
	}

}

func main() {
	arr := []int{3, 4, 1, 6, 2}
	countSubarrays(arr)
	fmt.Println("Result :", resCount)
}