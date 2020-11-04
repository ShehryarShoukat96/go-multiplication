package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var startTime int64 = time.Now().UnixNano() / int64(time.Millisecond)
	var start int = 1
	var end int = 12

	var tableString string = ""

	for tableDigit := start; tableDigit <= end; tableDigit++ {
		var lastValue int = (tableDigit * 12)

		for nthValue := tableDigit; nthValue <= lastValue; nthValue = (nthValue + tableDigit) {
			tableString = tableString + " " + strconv.Itoa(nthValue)
		}

		tableString = tableString + "\n"
	}

	fmt.Printf("%s\n", tableString)
	var endTime int64 = (time.Now().UnixNano() / int64(time.Millisecond)) - startTime

	fmt.Printf("%d \n", endTime)
}
