package main

import (
	"fmt"
)

func bubble(s []int) {
	for i := len(s) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	arr := []int{2, 10, 1, 9, 5, 6, 8, 3, 7, 4}
	bubble(arr)
	fmt.Println(arr)
}
