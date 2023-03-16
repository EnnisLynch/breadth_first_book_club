package main

import (
	"os"
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
	fileStream.Close() //be good file stewards
}
