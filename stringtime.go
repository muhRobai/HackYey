// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import (
    "fmt"
    "time"
)

func main() {
	date := "1997/01/01"
	fixdate, err := time.Parse("2006/01/02", date)
	if err != nil {
	fmt.Println(err)
	return
	}
	fmt.Println(fixdate.Unix())
}

// sudo chmod -R a+rw go.*
