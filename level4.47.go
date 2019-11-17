package main

import (
	"fmt"
	"strings"
	"strconv"
)

func elevator(floor string) int {
	bigger := 1 // min floor 1
	lower := 100 // max floor 100
	var stop int
	move := true // true for up and false for down
	value := strings.Replace(floor, " ","", -1)

	for i := 0; i < len(value); i++ {
		stop,_ = strconv.Atoi(string(value[i]))
		
		if stop == 0 {
			continue
		}
		//check elevator up or down
		if stop < bigger {
			move = false
		} else {
			move = true
		}
		
		//set the bigger floor
		if stop > bigger {
			bigger = stop
		}

		//check the lower floor if elevator down
		if stop < lower && move == false{
			lower = stop
		}



	}

	if move == true {
		return bigger
	}

	return lower
}

func main() {
	fmt.Println(elevator("0 1 0 0 0"))
	fmt.Println(elevator("2 5 0 0 4"))
	fmt.Println(elevator("0 0 0 0 0"))
	fmt.Println(elevator("2 7 2 0 2 0 3"))
}