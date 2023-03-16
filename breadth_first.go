package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Go doesn't have a tuple
// so we can create our own
type Chapter struct {
	chapter_start_byte int64
	chapter_end_byte   int64
}

// Helper method to create a new chapter
func NewChapter(start_byte int64, end_byte int64) *Chapter {
	var chapter = new(Chapter)
	chapter.chapter_start_byte = start_byte
	chapter.chapter_end_byte = end_byte
	return chapter
}

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
	var commonError error   //use as common error variable
	var chapters []*Chapter //contain all index values for chapters
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
			//fmt.Println(start_index, end_index)

			//Store the chapter index information
			var chapter = NewChapter(start_index, end_index)
			chapters = append(chapters, chapter) //appends to splice
		}
	} //end lines

	fileStream.Close() //be good file stewards
}
