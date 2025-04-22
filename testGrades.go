package main

import "fmt"

func main() {
	grade := 37
	diff := 5 - (grade % 5)
	gradeDiff := grade + diff

	if diff < 3 && gradeDiff >= 40 {
		grade += diff
	}

	fmt.Println(diff, grade)
}
