package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// This method will panic on errors
// just makes code a bit cleaner ...
// hint, hint, hint
func checkForErrorAndFail(someError error) {
	if someError != nil {
		panic(someError)
	}
}

// Main method
func main() {
	//Open our index file and do absolutely nothing with it!
	var commonError error //use as common error variable
	fileStream, commonError := os.Open("pelle.txt.csv")
	checkForErrorAndFail(commonError)

	//use a CSV encoder to read the values
	//requires import encoding/csv
	lines, commonError := csv.NewReader(fileStream).ReadAll()
	checkForErrorAndFail(commonError)

	//Loop through the index file and read all
	//of the index values and print
	for index, record := range lines {
		if index == 0 {
			//skip the header line
			continue //not required
		} else {
			start_index, commonError := strconv.ParseInt(record[0], 10, 64)
			checkForErrorAndFail(commonError)

			end_index, commonError := strconv.ParseInt(record[1], 10, 64)
			checkForErrorAndFail(commonError)

			//Print our start and end index to the
			//output stream to see that we did
			//actual work!
			fmt.Println(start_index, end_index)
		}
	} //end lines

	fileStream.Close() //be good file stewards
}
