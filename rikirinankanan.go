package main 

import (
	"log"
	"strings"
)

func Moving(value string) {
	payload := strings.Replace(value, " ", ",", 1)
	strSplit := strings.Split(payload, ",")
	word := strSplit[1]
	movingWord := strSplit[0]
	start := 1
	move := false
	pointer := 0
	for i := 0; i < len(movingWord) ; i++ {
		if start < len(word) {
			if string(movingWord[i]) == "l" {
				start++
			}

			if string(movingWord[i]) == "w" {
				for j := start; j < len(word); j++ {
					if (string(word[j]) == " ") && (start < len(word)) {
						start = j + 2
						break
					}
				}
			}

			if string(movingWord[i]) == "h" {
				start--
			}

			if move == true {
				start = start - 3
				move = false
			}
			
			if string(movingWord[i]) == "b" {
				for k := start; k > 0 ; k-- {
					if (string(word[k]) == " "){
						start = k + 2
						move = true
						pointer = start - 3
						break
					}

					if pointer == start {
						start = start - 3
					}
				}
			}
		}
	}
	log.Println(start)

}

func main(){
	Moving("ll aku mau hidup seribu tahun lagi")
	Moving("b robohnya surau kami")
	Moving("wh kaki yang terhormat")
	Moving("llwwllbb bahwa dalam suatu perjuangan kita harus")
}