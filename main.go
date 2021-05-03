/*Erica Chou
Kargo Internship Question:
Turn an array of Integers into an array of Strings */
package main

import (
	"fmt"
)

func main() {
	//logic for pulling input from the terminal

	//argsWithoutProg := os.Args[1:]

	//fmt.Println(argsWithoutProg)

	//print out original array
	myarr := []string{"1", "2", "4", "10", "1238"}
	fmt.Printf("Number Array: len=%d %v\n", len(myarr), myarr)
	numarray := numToString(myarr)
	fmt.Printf("String Number Array: len=%d cap=%d %v\n", len(numarray), cap(numarray), numarray)

}

func numToString(array []string) []string {
	var i, j, sizearr int
	sizearr = len(array)
	var numString []string
	var temp string
	for i = 0; i < sizearr; i++ {
		var toAdd string
		//seperate a single term
		temp = array[i]

		//loop through each part of this term and get the phonetic version of each term
		for j = 0; j < len(temp); j++ {
			toAdd += numConversion(temp[j])
		}

		//after we piece together our number we append it to the final array
		numString = append(numString, toAdd)

	}
	return numString
}

//function to convert each number we get into its phonetic version
func numConversion(x byte) string {
	switch x {
	case '0':
		return "Zero"
	case '1':
		return "One"
	case '2':
		return "Two"
	case '3':
		return "Three"
	case '4':
		return "Four"
	case '5':
		return "Five"
	case '6':
		return "Six"
	case '7':
		return "Seven"
	case '8':
		return "Eight"
	case '9':
		return "Nine"
	}
	return "error"
}
