package main

import (
  "strconv"
  "strings"
  "log"
  )

func movingBall(value string) {
  payload := strings.Replace(value, " ","", -1)
  biggerBall := 0
  smallerBall := 0
  moving := 0
  changed := false
  for i:= 0; i < len(payload); i++ {
	  biggerBall = 0
	  changed = false
    for j := 0; j < len(payload); j++ {
		item, err := strconv.Atoi(string(payload[j]))
		if err != nil {
			return
		}
		//set value now
		now := item
		//set bigger value
		if biggerBall == 0 {
			biggerBall = now
			continue
		}
		//check smaller value
		if smallerBall < item {
			smallerBall = item
			changed = true // for change left to right
		}

		if smallerBall > biggerBall {
			biggerBall = smallerBall
			changed = true
		}
		//add moving
		if (changed == true) && (smallerBall != now) {
			moving++
		}
	}
		
		strBall := strconv.Itoa(biggerBall)
		payload = strings.Replace(payload, strBall, "", 1)
	}
	log.Println("Moving Ball: ",moving)
	  
}


func main() {
  movingBall("2 1")
  movingBall("5 4 1")
  movingBall("9 1")
  movingBall("7 2 5 9 1")
}
