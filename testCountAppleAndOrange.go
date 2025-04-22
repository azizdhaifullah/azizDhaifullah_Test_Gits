package main

import "fmt"

func main() {
	s := 2
	t := 3
	a := 1
	b := 5
	apples := []int{-2}
	oranges := []int{-1}
	countApples := 0
	countOranges := 0

	for _, value := range apples {
		rangeApples := a + value
		if s <= rangeApples && rangeApples <= t {
			countApples++
		}
	}

	for _, value := range oranges {
		rangeOranges := b + value
		if s <= rangeOranges && rangeOranges <= t {
			countOranges++
		}
	}

	fmt.Println(countApples, countOranges)
}
