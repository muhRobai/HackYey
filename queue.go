package main

import (
	"fmt"
)


func main(){
	i := 1
	var numb []int
	for {
		if i == 1 {
			fmt.Print("insert service: ")
		} else {
			fmt.Print("insert queue time: ")
		}

		var value int
		var end string
		fmt.Scanf("%d", &value)
		numb = append(numb, value)
		fmt.Print("enter to continue insert queue/ DONE to stop: ")
		fmt.Scanf("%s", &end)
		if end == "DONE" {
			break
		}
		i++
	}

	machine := numb[0]
	queue := 0
	time := 0
	var bigger int
	for {
		queue++
		if queue == len(numb) || queue > len(numb) {
			break
		}
		bigger = numb[queue]
		if machine == 1 {
			time+= numb[queue]
		} else {
			for l := 0; l < machine; l++ {
				if queue + 1 == len(numb){
					bigger = numb[queue]
					break
				}

				if bigger < numb[queue+l] {
					bigger = numb[queue+l]
				}
				queue += l
			}
			time += bigger
		}
	}

	fmt.Println(time)
}