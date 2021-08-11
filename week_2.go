package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Converting a String to a Float64
func conv(s []string) []float64{
	var converted []float64
	for _, va := range s {
		convert, _ := strconv.ParseFloat(va, 64)
		converted = append(converted, convert)
	}
	return converted
}
// Variable Number of String
func varNumber(strs ...string) (varSum []float64) {
	for _, vals := range strs {
		if strings.Contains(vals, "*") {
			newStr := conv(strings.Split(vals, "*"))
			newfloat := Multiply(newStr)
			varSum = append(varSum, newfloat)
		}
		if strings.Contains(vals, "/") {
			newStr := conv(strings.Split(vals, "/"))
			newfloat := Divs(newStr)
			varSum = append(varSum, newfloat)
		}
		if strings.Contains(vals, "+") {
			newStr := conv(strings.Split(vals, "+"))
			newfloat := Addition(newStr)
			varSum = append(varSum, newfloat)
		}
		if strings.Contains(vals, "-") {
			newStr := conv(strings.Split(vals, "-"))
			newfloat := Subtraction(newStr)
			varSum = append(varSum, newfloat)
		}
	}
	return
}

// Divs Function
func Divs(slice []float64) float64{
	var sum float64 = 1
	for i, va := range slice {
		if i == 0 {
			sum = va
		} else {
			sum /= va
		}
	}
	return sum
}

// Addition Function
func Addition(slice []float64) float64{
	var total float64 = 1
	for i, va := range slice {
		if i == 0 {
			total = va
		} else {
			total += va
		}
	}
	return total
}

// Multiply FUnction
func Multiply(slice []float64) float64{
	var  multiplied float64 = 1
	for i, va := range slice {
		if i == 0 {
			multiplied *= va
		} else {
			multiplied *= va
		}
	}
	return multiplied
}

// Subtraction Function
func Subtraction(slice []float64) float64{
	var subtracted float64 = 1
	for i, va := range slice {
		if i == 0 {
			subtracted = va
		} else {
			subtracted -= va
		}
	}
	return subtracted
}

func main() {
	ar:= varNumber("2*3*4*5", "25/5/2", "2+3+4+5.9+6+7", "20.54-7.65-2.897")
	fmt.Println(ar)
}




