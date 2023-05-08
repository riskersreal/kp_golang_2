package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var tempSlice, returnSlice []int
	var appearCount int = 1
	var isExist bool
	var countMap = map[int]int{}

	// error handling here

	tempSlice = make([]int, len(angka))
	for i, c := range angka {
		tempSlice[i], _ = strconv.Atoi(string(c))
	}

	for i := 0; i < len(tempSlice); i++ {
		_, isExist = countMap[tempSlice[i]]
		if isExist {
			countMap[tempSlice[i]]++
		} else {
			countMap[tempSlice[i]] = 1
		}
	}

	for key, value := range countMap {
		if value == appearCount {
			returnSlice = append(returnSlice, key)
		}
	}
	return returnSlice
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
