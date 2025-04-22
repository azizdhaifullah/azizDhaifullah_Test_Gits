package main

import "fmt"

func Abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}

func main() {
	x1 := 0
	v1 := 2
	x2 := 5
	v2 := 3

	hasil := "NO"

	if x1 == x2 && v1 == v2 {
		hasil = "YES"
	}
	if v1 == v2 {
		hasil = "NO"
	}

	if (x2-x1)%(v1-v2) == 0 && (x2-x1)/(v1-v2) >= 0 {
		hasil = "YES"
	}

	a := 126
	testHasil := a % 60

	fmt.Println(hasil, testHasil)
}
