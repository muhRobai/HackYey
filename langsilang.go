package main 

import (
  	"strings"
	"strconv"
	"log"
)

func lineToline(value string) int{
	payload := strings.Split(value, "-")
	vertical := strings.Split(payload[0], " ")
	hosrizontal := strings.Split(payload[1], " ")
	lengthArray := len(hosrizontal)
	more := false
	if len(vertical) > len(hosrizontal) {
  		lengthArray = len(vertical)
  		more = true
  	}
  //for make sure length vertical and horizontal 
  	for i := 0; i < lengthArray; i++ {
		if len(hosrizontal) == len(vertical) {
			break
	 	}

	  	if more == true {
			hosrizontal = append(hosrizontal, "0")
	  	} else {
			vertical = append(vertical, "")
	  	}
  	}

	found := 0
	line := 0
	//getting array vertical
  	for j := 0; j < lengthArray; j++ {
		verticalValue, _ := strconv.Atoi(vertical[j])
		if verticalValue == 0 { 
			continue
		} 
		//getting horizontal value  
		for l := 0; l < 1; l++ {
			horizontalValue, _ := strconv.Atoi(hosrizontal[l])
			if horizontalValue == 0 { 
				continue
			}
			//get x,y coordinate
			for k := 0; k < horizontalValue; k++ {
				if verticalValue > k && horizontalValue > line {
					found++
				}
			}
			line++	
		}
  	}
  	return found
}

func main(){
  log.Println(lineToline("1 1-1 1"))
  log.Println(lineToline("1 1-100"))
  log.Println(lineToline("2 2 2-10 10 10"))
  log.Println(lineToline("2 4 1 2 5 3-3 2 3 1"))
}
