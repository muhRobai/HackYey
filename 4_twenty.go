package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minisweeper(payload string) int {
	value := strings.Split(payload, " ")
	fristValue := 0
	lastValue := 0
	returnValue := 0
	trap := false

	for i := 0; i < len(value); i++ {
		numb, _ := strconv.Atoi(value[i])
		fristValue = numb

		//for add trap value where boxs value = 1 and trap is false
		if fristValue == 1 && lastValue != fristValue && trap == false{
			returnValue++
			trap = true
		}

		//for add trap value where boxs value = 2 and trap value false
		if fristValue == 2 && trap == false {
			returnValue += 2
			trap = true
		}
		
		//reset trap to false 
		if lastValue == fristValue {
			trap = false
		}

		//for set last value if having 2 data or more
		if i > 0 {
			lastValue = fristValue
		}
	}

	return returnValue
}

func main(){
	fmt.Println(minisweeper("0 1 1 0 0 0"))
	fmt.Println(minisweeper("1 0 0 0 1 1 0 2 1"))
	fmt.Println(minisweeper("0 0 0 0 0 0 0"))
	fmt.Println(minisweeper("0 1 1 0 1 2 1 0 1"))
}