package main

import (
	"fmt"
)

func HeksYes(payload string) {
	move := false
	start := ""
	seconds := 0
	for i := 0; i < len(payload); i++ {
		//fot frist click and set start value
		if i == 0 {
			start = string(payload[i])
			continue
		}

		//check same value or not
		if string(payload[i]) == start {
			move = true
		}

		// if have some value add seconds
		if move == true {
			seconds++
		} else {
			//if don't have same value and seconds greater than 0 reduce seconds 
			if seconds > 0 {
				seconds--	
			}
		}

		//set start and move to early
		start = string(payload[i])
		move = false
	}

	//covert seconds to 60
	if seconds == 0 {
		seconds = 60
	}

	//if seconds have value greater than 60
	for j := 0; j < seconds; j++ {
		if j >= 60 {
			seconds -= 60
			j = 0	
		}
	}

	fmt.Println(seconds)
}

func main() {
	HeksYes("aiaiaiai")
	HeksYes("aaiiaa")
	HeksYes("aaaaaaaaaaaaaaaaaaaaaa")
}