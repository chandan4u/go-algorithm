package main

import (
	"bytes"
	"fmt"
	"strings"
)

var count = 0
var indexListArray bytes.Buffer
var visited [6][6]bool

func main() {
	arrString := [][]string{
		{"B", "B", "A", "B", "B", "N"},
		{"B", "B", "M", "B", "B", "O"},
		{"B", "B", "A", "B", "B", "Z"},
		{"B", "O", "Z", "O", "N", "A"},
		{"B", "B", "O", "Z", "B", "M"},
		{"B", "B", "N", "A", "M", "A"},
	}
	searchStr := "AMAZON"
	isFoundStatus := searchString(arrString, searchStr)
	fmt.Println("Result :", isFoundStatus)
	fmt.Println("Find Count ::", count)
	fmt.Println("Found Index", indexListArray.String())
}

func searchString(arrString [][]string, searchStr string) bool {
	splitSearchStr := strings.Split(searchStr, "")
	for index, value := range arrString {
		for innerIndex, innerValue := range value {
			// Check first character of search string
			if splitSearchStr[0] == innerValue {
				indexStr := fmt.Sprintf("{ [%d,%d]", index, innerIndex)
				fmt.Println(" UpperIndex Check Pass + :: ", index, "-", innerIndex, splitSearchStr[0])
				checkAllDirection(arrString, splitSearchStr, index, innerIndex, 0, indexStr)
				fmt.Println("Visited", visited)
			}
		}
	}
	if count > 0 {
		return true
	}
	return false
}

func checkAllDirection(arrString [][]string, splitSearchStr []string, upperIndex, innerIndex, searchStrIndex int, indexList string) {
	if searchStrIndex == len(splitSearchStr) - 1 && searchStrIndex != len(splitSearchStr) {
		indexListArray.WriteString(fmt.Sprintf("%s },", indexList))
		count = count + 1
		return
	}

	visited[upperIndex][innerIndex] = true

	if indexValidate(arrString, upperIndex+1, innerIndex) && !visited[upperIndex+1][innerIndex] && arrString[upperIndex+1][innerIndex] == splitSearchStr[searchStrIndex+1] {
		fmt.Println(" UpperIndex Check Pass + :: ", upperIndex+1, "-", innerIndex, splitSearchStr[searchStrIndex])
		indexList = fmt.Sprintf("%s, [%d,%d]", indexList, upperIndex+1, innerIndex)
		checkAllDirection(arrString, splitSearchStr, upperIndex+1, innerIndex, searchStrIndex+1, indexList)
		visited[upperIndex+1][innerIndex] = true
	}
	if indexValidate(arrString, upperIndex-1, innerIndex) && !visited[upperIndex-1][innerIndex] && arrString[upperIndex-1][innerIndex] == splitSearchStr[searchStrIndex+1] {
		fmt.Println(" UpperIndex Check Pass - :: ", upperIndex-1, "-", innerIndex, splitSearchStr[searchStrIndex])
		indexList = fmt.Sprintf("%s, [%d,%d]", indexList, upperIndex-1, innerIndex)
		checkAllDirection(arrString, splitSearchStr, upperIndex-1, innerIndex, searchStrIndex+1, indexList)
		visited[upperIndex-1][innerIndex] = true
	}
	if indexValidate(arrString, upperIndex, innerIndex+1) && !visited[upperIndex][innerIndex+1] && arrString[upperIndex][innerIndex+1] == splitSearchStr[searchStrIndex+1] {
		fmt.Println(" InnerIndex Check Pass  + :: ", upperIndex, "-", innerIndex+1, splitSearchStr[searchStrIndex])
		indexList = fmt.Sprintf("%s, [%d,%d]", indexList, upperIndex, innerIndex+1)
		checkAllDirection(arrString, splitSearchStr, upperIndex, innerIndex+1, searchStrIndex+1, indexList)
		visited[upperIndex][innerIndex+1] = true
	}
	if indexValidate(arrString, upperIndex, innerIndex-1) && !visited[upperIndex][innerIndex-1] && arrString[upperIndex][innerIndex-1] == splitSearchStr[searchStrIndex+1] {
		fmt.Println(" InnerIndex Check Pass  - :: ", upperIndex, "-", innerIndex-1, splitSearchStr[searchStrIndex])
		indexList = fmt.Sprintf("%s, [%d,%d]", indexList, upperIndex, innerIndex-1)
		checkAllDirection(arrString, splitSearchStr, upperIndex, innerIndex-1, searchStrIndex+1, indexList)
		visited[upperIndex][innerIndex-1] = true
	}
}

func indexValidate(arrString [][]string, upperIndex, innerIndex int) bool {
	if upperIndex >= 0 && innerIndex >= 0 && len(arrString) > upperIndex && len(arrString[upperIndex]) > innerIndex {
		return true
	} else {
		return false
	}
}
