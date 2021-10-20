package main

import (
	"fmt"
)

func main() {
	listItems := []int{12, 1, 5, 2, 11, 422, 17, 3, 8, 9}
	keyItem := 12
	fmt.Println(linearSearch(listItems, keyItem))
}

func linearSearch(listItems []int, keyItem int) string {
	for ind, value := range listItems {

		if value == keyItem {
			return "element found at location" + string(rune(ind))
		} else {
			return "element not found"
		}
	}
	return ""
}
