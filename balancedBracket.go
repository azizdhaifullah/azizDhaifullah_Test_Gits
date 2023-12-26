package main

import (
	"fmt"
	"slices"
	"strings"
)

var bracket = map[string]string{
	"[": "]",
	"{": "}",
	"(": ")",
}

func validateBracket(param []string, minIdx int, maxIdx int) string {
	decreaseMaxIdx := 0
	increaseMinIdx := 0
	if maxIdx <= minIdx {
		return "Yes"
	} else {
		if bracket[param[minIdx]] == param[maxIdx] {
			decreaseMaxIdx = 1
			increaseMinIdx = 1
		}

		if bracket[param[maxIdx-1]] == param[maxIdx] {
			decreaseMaxIdx = 2
		}

		if bracket[param[minIdx]] == param[minIdx+1] {
			increaseMinIdx = 2
		}

		if decreaseMaxIdx == 0 && increaseMinIdx == 0 {
			return "No"
		}

		return validateBracket(param, minIdx+increaseMinIdx, maxIdx-decreaseMaxIdx)
	}
}

func doBalancedBracket(arrayParams []string) string {
	for _, value := range arrayParams {
		if !slices.Contains([]string{"[", "]", "{", "}", "(", ")"}, value) {
			return "No"
		}
	}
	return validateBracket(arrayParams, 0, len(arrayParams)-1)
}

func main() {
	param := "{ [ ( ] ) }"
	trim := strings.ReplaceAll(param, " ", "")
	arrayParams := strings.Split(trim, "")
	result := doBalancedBracket(arrayParams)
	fmt.Println(result)
}
