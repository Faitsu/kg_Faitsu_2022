/*Erica Chou
Kargo Internship Question:
Turn an array of Integers into an array of Strings */
package main

import (
	"fmt"
)

func main() {
	//Figuring out if code will work with my own array
	//logic might be slightly different when we pass args into it from the terminal

	//assuming the terminal input will give me an array of INtegers in the form of a string
	myarr := "1 2 10 12 129"
	//print out original array
	fmt.Printf("Number Array: len=%d %v\n", len(myarr), myarr)
	numarray := numToString(myarr)
	fmt.Printf("String Number Array: len=%d cap=%d %v\n", len(numarray), cap(numarray), numarray)

}

func numToString(array string) []string {
	var i, sizearr int
	sizearr = len(array)
	var numString []string
	var temp string
	for i = 0; i < sizearr; i++ {
		if array[i] == ' ' {
			numString = append(numString, temp)
			temp = ""
		} else {
			temp += numConversion(array[i])
		}

		//for the last case, we have to also append the num string
		if i == sizearr-1 {
			numString = append(numString, temp)
		}

		//after we piece together our number we append it to the final array

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
