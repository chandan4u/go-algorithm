package main

import (
	"fmt"
	"strings"
)

func main() {
	arrString := [][]string{
		{"A", "X", "M", "C"},
		{"P", "E", "S", "Q"},
		{"L", "A", "L", "B"},
		{"D", "Y", "T", "H"},
	}
	searchStr := "MXAPEY"
	isFoundStatus := searchString(arrString, searchStr)
	fmt.Println("Result :", isFoundStatus)
}

func searchString(arrString [][]string, searchStr string) bool {
	splitSearchStr := strings.Split(searchStr, "")
	for index, value := range arrString {
		for innerIndex, innerValue := range value {
			if splitSearchStr[0] == innerValue {
				status := checkAllDirection(arrString, splitSearchStr, index, innerIndex, 1)
				if status {
					return true
				}
			}
		}
	}
	return false
}

func checkAllDirection(arrString [][]string, splitSearchStr []string, upperIndex, innerIndex, searchStrIndex int) bool {
	if indexValidate(arrString, upperIndex+1, innerIndex) && arrString[upperIndex+1][innerIndex] == splitSearchStr[searchStrIndex] {
		fmt.Println(" UpperIndex Check Pass + :: ", upperIndex+1, "-", innerIndex)
		checkAllDirection(arrString, splitSearchStr, upperIndex+1, innerIndex, searchStrIndex+1)
	}
	if indexValidate(arrString, upperIndex-1, innerIndex) && arrString[upperIndex-1][innerIndex] == splitSearchStr[searchStrIndex] {
		fmt.Println(" UpperIndex Check Pass - :: ", upperIndex-1, "-", innerIndex)
		checkAllDirection(arrString, splitSearchStr, upperIndex-1, innerIndex, searchStrIndex+1)
	}
	if indexValidate(arrString, upperIndex, innerIndex+1) && arrString[upperIndex][innerIndex+1] == splitSearchStr[searchStrIndex] {
		fmt.Println(" InnerIndex Check Pass  + :: ", upperIndex, "-", innerIndex+1)
		checkAllDirection(arrString, splitSearchStr, upperIndex, innerIndex+1, searchStrIndex+1)
	}
	if indexValidate(arrString, upperIndex, innerIndex-1) && arrString[upperIndex][innerIndex-1] == splitSearchStr[searchStrIndex] {
		fmt.Println(" InnerIndex Check Pass  - :: ", upperIndex, "-", innerIndex-1)
		checkAllDirection(arrString, splitSearchStr, upperIndex, innerIndex-1, searchStrIndex+1)
	}
	return searchStrIndex == len(splitSearchStr)
}

func indexValidate(arrString [][]string, upperIndex, innerIndex int) bool {
	if upperIndex > 0 && innerIndex > 0 && len(arrString) > upperIndex && len(arrString[upperIndex]) > innerIndex {
		return true
	} else {
		return false
	}
}
