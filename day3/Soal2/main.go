package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var returnMap = map[string]int{}
	var isExist bool

	for i := 0; i < len(slice); i++ {
		_, isExist = returnMap[slice[i]]
		if isExist {
			returnMap[slice[i]]++
		} else {
			returnMap[slice[i]] = 1
		}
	}
	return returnMap
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         //map[]
}
