package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", fmt.Errorf("TrimSpace empt: %w", errorEmptyInput)
	}

	tmp1 := strings.Split(input, "-")
	tmp2 := strings.Split(input, "+")
	tmp1L := countLen(tmp1)
	tmp2L := countLen(tmp2)
	if tmp1L != 2 && tmp2L != 2 {
		return "", fmt.Errorf("Not right len: %w", errorNotTwoOperands)
	}
	var digits []string
	var symb string
	if input[0] == 45 {
		digits, symb = getNums(input[1:])
		digits[0] = "-" + digits[0]

	} else {
		digits, symb = getNums(input)
	}

	num1, err := strconv.Atoi(digits[0])
	if err != nil {
		return "", fmt.Errorf("First arr cant conv: %w", err)
	}

	num2, err := strconv.Atoi(digits[1])
	if err != nil {
		return "", fmt.Errorf("First arr cant conv: %w", err)
	}
	if symb == "-" {
		return strconv.Itoa(num1 - num2), nil
	}
	return strconv.Itoa(num1 + num2), nil
}

// func main() {
// 	fmt.Println(StringSum(" -24 - 55 "))
// }

func countLen(arr []string) (count int) {
	for _, v := range arr {
		if v != "" {
			count++
		}
	}
	return
}

func getNums(str string) ([]string, string) {
	var tmp string
	var symb string
	for _, v := range str {
		if string(v) == "-" {
			symb = "-"
			tmp += "OPERATION"
		} else if string(v) == "+" {
			symb = "+"
			tmp += "OPERATION"
		} else if string(v) != " " {
			tmp += string(v)
		}
	}
	return strings.Split(tmp, "OPERATION"), symb
}
