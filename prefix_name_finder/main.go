package main

import (
	"fmt"
	"strings"
)

func prefixNameFinder(definedNames []string, requestedName string) bool {
	for i:=0; i<len(definedNames); i++ {
		splitedDefinedName := strings.Split(definedNames[i], " ")
		splitRequestedName := strings.Split(requestedName, " ")
		j := 0
		for m:=0; m<len(splitedDefinedName); m++ {
			checkStatus := prefixCheck(splitRequestedName[j], splitedDefinedName[m])
			if checkStatus {
				j++
			}
			if j == len(splitRequestedName) {
				return true
			}
		}
	}
	return false
}

func prefixCheck(requestedWord string, definedWord string) bool {
	for k:=0; k<len(definedWord); k++ {
		if len(requestedWord) <= k {
			return true
		}
		if  string(definedWord[k]) != string(requestedWord[k]) {
			return false
		}
	}
	return true
}

func main() {
	definedNames := []string{
		"Chandan Kumar Rai",
		"Mohit Singh Rajput",
		"Ram Kumar Mahto",
	}
	requestedName := "C Kumar Ray"
	res := prefixNameFinder(definedNames, requestedName)
	fmt.Println("Result ::", res)
}
