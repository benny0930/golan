package main

import "strconv"

func main() {
	aNum1 := []int{2, 3, 4, 5, 6, 7, 8, 9}
	aNum2 := []int{2, 3, 4, 5, 6, 7, 8, 9}
	//aStr := []string{"","","","","","","","",}
	for _, v1 := range aNum1 {
		sStr := ""
		for _, v2 := range aNum2 {
			sStr = sStr + (strconv.Itoa(v2) + "X" + strconv.Itoa(v1) + " = ")
			num := v2 * v1
			if num > 9 {
				sStr = sStr + strconv.Itoa(num) + "  "
			} else {
				sStr = sStr + " " + strconv.Itoa(num) + "  "
			}
		}
		println(sStr)
	}
}
