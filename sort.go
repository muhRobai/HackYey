package main

import (
	"fmt"
	"strings"
	"strconv"
)

func convertArr(payload string) []int{
	//conver string array to int array
	var result []int
	value := strings.Split(payload, " ")
	for i := 0; i < len(value); i++ {
		val, err := strconv.Atoi(string(value[i]))
		if err != nil {
			fmt.Println(err)
			return result
		}
		result = append(result, val)
	}
	return result
}

func main() {
	//select value
	// value := "1 3 2 4-1 10 8 2"
	// value := "10 20 30-0 0 0"
	// value := "1 2 3 4-1 2 3 4"
	value := "5 5 1-0 1 2"
	
	//prepare data
	arrString := strings.Split(value, "-")
	one := convertArr(arrString[0])
	two := convertArr(arrString[1])

	//bubble sort
	for i := 0; i < len(one); i++ {
		for j := len(one)-1; j >= i + 1; j-- {
			if int(one[j]) < int(one[j-1]){
				one[j],one[j-1] = one[j-1],one[j] // swap array
			}

			if int(two[j]) > int(two[j-1]) {
				two[j], two[j-1] = two[j-1], two[j]
			} 
		}
	}

	var total int
	//get sum difference 
	for l := 0; l < len(one); l++ {
		difference := int(two[l]) - int(one[l])
		if difference < 0 {
			difference = difference * -1
		}
		total += difference
	}

	fmt.Println(total)
}
