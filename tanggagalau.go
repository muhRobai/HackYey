package main

import (
  "strings"
  "strconv"
  "fmt"
)

// for set month in string value
var month_string = []string {
  "jan",
  "feb",
  "mar",
  "apr",
  "mei",
  "jun",
  "jul",
  "agu",
  "sep",
  "okt",
  "nov",
  "des",
}


func dateValidator(value string) string{
	//replace "/" and "-" for preparetion data
	date := strings.Replace(value, "/"," ", -1)
	fixdate := strings.Replace(date, "-"," ", -1)
	strSplit := strings.Split(fixdate, " ")
    haveStringMonth := false
    validDay := 0
	validMonth := 0
	validyear := false

    for i := 0; i < len(strSplit); i++ {
		//check have year value or not
		if len(strSplit[i]) == 4 {
			validyear = true
		}
		
		//check data have a string month or not
      	for j := 0; j < len(month_string); j++ {
        	if string(strSplit[i]) == month_string[j] {
				haveStringMonth = true
        	}
		}
		
		/*
		if data don't have string 
		month data must check 
		have 2 month or not
		*/
		if haveStringMonth == false {
			for k := 1; k < 13; k++ {
			  mount, _ := strconv.Atoi(strSplit[i])
			  if mount == k {
				  validMonth++
			  }
			}	
		}

		//validate day, in my opinion day in 1 month is 31, not specific
      	for l := 1; l < 32; l++ {
			day, _ := strconv.Atoi(strSplit[i])
        	if day == l {
          		validDay++
        	}
    	}
	}

	//checkvalid or not date value
	if validyear == false  || validMonth >= 2 || validDay < 1{
		return "date-not-valid"
	}
	
	return "date-valid"
}

func main(){
  fmt.Println(dateValidator("2016/10/10"))
  fmt.Println(dateValidator("2016-10-20"))
  fmt.Println(dateValidator("15/2/1998"))
  fmt.Println(dateValidator("04/05/2010"))
  fmt.Println(dateValidator("jan 3 2004"))
  fmt.Println(dateValidator("2011-08-22"))
  fmt.Println(dateValidator("23 mar 2020"))
  fmt.Println(dateValidator("04/07/2010"))
  fmt.Println(dateValidator("1965-3-agu"))
}
