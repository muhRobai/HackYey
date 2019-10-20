package main

import (
	"fmt"
	"strconv"
)

func sumRolling(item string) {
	
	roling := 0
	moving := "kiri"
	defaultMoving := "kiri"
	start := 0

	for i := 0; i < len(item); i++ {
		value,_ := strconv.Atoi(string(item[i])) 
		if i == 0 {
			start = value
			continue
		}
		if start - value > 0 {
			if moving == "kiri" {
				moving = "kanan"
			} else {
				moving = "kiri"
			}
		}

		if start - value < 0 {
			if moving == "kanan" {
				moving = "kiri"
			}
		}

		if moving != defaultMoving {
			roling++
			defaultMoving = moving
		}
		start = value
		
	}
	fmt.Println(roling)
}

func main()  {
	sumRolling("12345")
	sumRolling("2121")
	sumRolling("981")
}