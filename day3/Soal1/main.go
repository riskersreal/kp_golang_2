package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var arraySame []string
	var isSame bool = false
	arraySame = arrayA

	for i := 0; i < len(arrayB); i++ {
		for j := 0; j < len(arrayA); j++ {
			if arrayB[i] == arrayA[j] {
				isSame = true
			}
		}
		if isSame == false {
			arraySame = append(arraySame, arrayB[i])
		}
		isSame = false
	}
	return arraySame

	// solusi yg lebih benar ada di catatan
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil "jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []

}
