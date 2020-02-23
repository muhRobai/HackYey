package main

import (
	"fmt"
)

func main(){
	var value, value2, runner string
	result := 0
	var afterRemove string
	haveSameString := false

	for {
		result = 0
		afterRemove = ""
		fmt.Scanf("%s %s", &value, &value2)
	
		for i := 0; i < len(value); i++ {
			haveSameString = false
			for k := 0; k < len(value2); k++ {
				if string(value[i]) == string(value2[k]) {
					haveSameString = true
				}
			}

			if haveSameString == false {
				result++
			} else {
				afterRemove += string(value[i])
			}
		}
		
		if afterRemove != value2 {
			for l := 0; l < len(value2); l++ {
				haveSameString = false
				for j := 0; j < len(afterRemove); j++{
					if string(value2[l]) == string(afterRemove[j]){
						haveSameString = true
					}
				}
				if haveSameString == false {
					result++
				}
			}
			
		}
		fmt.Println(result)
		fmt.Println("end the program : DONE")
		fmt.Scanf("%s", &runner)
		if runner == "DONE" {
			break
		}
	}
}