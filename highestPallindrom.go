package main

import (
	"fmt"
	"strings"
)

func generateHighestPallindrom(param []string, minIdx int, maxIdx int) string {
	if maxIdx <= minIdx {
		return strings.Join(param, "")
	} else {
		if param[minIdx] != param[maxIdx] {
			if param[maxIdx] < param[minIdx] {
				param[maxIdx] = param[minIdx]
			} else {
				param[minIdx] = param[maxIdx]
			}
		}
		return generateHighestPallindrom(param, minIdx+1, maxIdx-1)
	}

}

func generateArray(params string, count int, tempArr []string) []string {
	if count >= len(params) {
		return tempArr
	} else {
		tempArr = append(tempArr, string(params[count]))
		return generateArray(params, count+1, tempArr)
	}
}

func main() {
	param := "3943"
	arrayParams := generateArray(param, 0, nil)
	result := generateHighestPallindrom(arrayParams, 0, len(arrayParams)-1)
	fmt.Println("Highest Pallindrom :", result)
}
