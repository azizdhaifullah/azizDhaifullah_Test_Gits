package main

import "fmt"

type ReturnMapped map[string]int

var alphabeticArr = func() ReturnMapped {
	charArray := make(ReturnMapped)
	count := 1
	for i := 'a'; i <= 'z'; i++ {
		charArray[string(i)] = count
		count++
	}

	return charArray
}

func stringToArray(s string) (charArray []string) {
	for _, char := range s {
		charArray = append(charArray, string(char))
	}

	return charArray
}

func mappedArray(charArray []string) ReturnMapped {
	mappedArray := make(ReturnMapped)
	for _, value := range charArray {
		_, ok := mappedArray[value]

		if ok {
			mappedArray[value] = mappedArray[value] + alphabeticArr()[value]
		} else {
			mappedArray[value] = alphabeticArr()[value]
		}

	}
	return mappedArray
}

func weightedStrings(alpha string, queries []int) (returnArray []string) {
	charArray := stringToArray(alpha)
	returnData := mappedArray(charArray)

	count := 0
	for _, value := range returnData {
		if value == queries[count] {
			returnArray = append(returnArray, "Yes")
		} else {
			returnArray = append(returnArray, "No")
		}
		count++
	}
	return
}

func main() {
	stringData := "aaaabbbzzz"
	queries := []int{4, 6, 78}
	fmt.Println("Dictionary =", alphabeticArr())
	fmt.Println("String =", stringData)
	fmt.Println("Queries =", queries)
	fmt.Println("Result =", weightedStrings(stringData, queries))
}
