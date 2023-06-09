package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
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

	//read our book text data into a byte array
	buffer, commonError := os.ReadFile("pelle.txt") //not this doesn't need to be closed after
	checkForErrorAndFail(commonError)

	var scanners []*bufio.Scanner //the tokenizers for each chapter

	for _, record := range chapters {
		//fmt.Println(record.chapter_start_byte, record.chapter_end_byte)
		var slice = buffer[record.chapter_start_byte:record.chapter_end_byte]
		var chapter_as_string = string(slice)

		//fmt.Println(chapter_as_string[0:100]) //print first 100 words
		string_reader := strings.NewReader(chapter_as_string)
		scanner := bufio.NewScanner(string_reader)
		//Tokenize by words, default is by newline
		scanner.Split(bufio.ScanWords)
		scanners = append(scanners, scanner)

	} //end chapters

	//Loop through all scanners and as
	//long as one-token is found,
	//create a new book

	//This variable will hold our new book
	//as string builder is the best way to concatenate
	//a large number of strings
	var stringBuilder strings.Builder

	var tokenFound = true
	var character_count = 0
	for tokenFound {
		tokenFound = false
		for _, scanner := range scanners {
			if scanner.Scan() {
				//True if there is a word!
				tokenFound = true
				tempWord := scanner.Text()
				character_count = character_count + len(tempWord)

				if stringBuilder.Len() > 0 {
					//Solves problem of spacing
					stringBuilder.WriteString(" ")
				}
				stringBuilder.WriteString(tempWord)

				if character_count > 80 {
					character_count = 0
					//Add new line to make chapters more readable
					stringBuilder.WriteString("\n")
				}
			}
		}
	} //end foreach token
	//Output our great work to console
	//fmt.Println(stringBuilder.String())

	//output our book to a file
	outputStream, commonError := os.Create("output.txt")
	checkForErrorAndFail(commonError)

	outputStream.WriteString(stringBuilder.String())
	outputStream.Sync()
	outputStream.Close()

}
