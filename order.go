package main

import "fmt"

func main() {
	aArr1 := [...]int{3, 1, 5, 4, 7, 10, 9, 8}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if aArr1[j] > aArr1[i] {
				num := aArr1[j]
				aArr1[j] = aArr1[i]
				aArr1[i] = num
			}
		}
	}
	fmt.Println(aArr1)

}
