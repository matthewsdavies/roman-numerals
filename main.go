package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// set up the reader for the input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number between 1- 3,999 to convert to Roman Numerals: ")

	// get the input and trim the space
	text, _ := reader.ReadString('\n')
	input := strings.TrimSpace(text)

	// validate the input is a number
	i, err := strconv.Atoi(input)
	if err != nil {
		// not a number force exit
		fmt.Printf("The value is not a number\n")
		return
	}

	// validate the number is in range
	if i == 0 || i > 3999 {
		// not a valid number force exit
		fmt.Printf("The value must be between 1 - 3,999\n")
		return
	}

	// convert to roman numerals
	s := convert(i)

	// output
	fmt.Printf("%d Coverted to Roman Numerals is: %s\n", i, s)
}

func convert(i int) string {
	// convert to string - easier to handle
	// get the length of the string to break it down
	s := strconv.Itoa(i)
	l := len(s)

	switch l {
	case 1:
		u := units(s[l-1:])
		return u
	case 2:
		u := units(s[l-1:])
		t := tens(s[l-2 : l-1])
		return t + u
	case 3:
		u := units(s[l-1:])
		t := tens(s[l-2 : l-1])
		h := hundreds(s[l-3 : l-2])
		return h + t + u
	default:
		u := units(s[l-1:])
		t := tens(s[l-2 : l-1])
		h := hundreds(s[l-3 : l-2])
		th := thousands(s[l-4 : l-3])
		return th + h + t + u
	}
}

// return the units value
func units(u string) string {
	var s string

	switch u {
	case "0":
		s = ""
	case "1":
		s = "I"
	case "2":
		s = "II"
	case "3":
		s = "III"
	case "4":
		s = "IV"
	case "5":
		s = "V"
	case "6":
		s = "VI"
	case "7":
		s = "VII"
	case "8":
		s = "VIII"
	case "9":
		s = "IX"
	}

	return s
}

// return then tens value
func tens(t string) string {
	var s string

	switch t {
	case "0":
		s = ""
	case "1":
		s = "X"
	case "2":
		s = "XX"
	case "3":
		s = "XXX"
	case "4":
		s = "XL"
	case "5":
		s = "L"
	case "6":
		s = "LX"
	case "7":
		s = "LXX"
	case "8":
		s = "LXXX"
	case "9":
		s = "XC"
	}

	return s
}

// return then hundreds value
func hundreds(h string) string {
	var s string

	switch h {
	case "0":
		s = ""
	case "1":
		s = "C"
	case "2":
		s = "CC"
	case "3":
		s = "CCC"
	case "4":
		s = "CD"
	case "5":
		s = "D"
	case "6":
		s = "DC"
	case "7":
		s = "DCC"
	case "8":
		s = "DCCC"
	case "9":
		s = "DM"
	}

	return s
}

// return then thousands value
func thousands(t string) string {
	var s string

	switch t {
	case "0":
		s = ""
	case "1":
		s = "M"
	case "2":
		s = "MM"
	case "3":
		s = "MMM"
	}

	return s
}
