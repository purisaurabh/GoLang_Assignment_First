/*

Write a program to calculate the simple interest
First-line has the comma-separated values of the Principal, rate and time (in years) respective
*constraints: Round simple interest float value to 2 decimal places

*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Program to find the Simple Interest")

	var principal, rate, time, simple_interest float64

	fmt.Println("Enter the value of the principal")
	fmt.Scanln(&principal)

	fmt.Println("Enter the value of the rate")
	fmt.Scanln(&rate)

	fmt.Println("Enter the value of time")
	fmt.Scanln(&time)

	simple_interest = (principal * rate * time) / 100

	simple_interest_value := fmt.Sprintf("%.2f", simple_interest)
	fmt.Println("The value of the simple interest is ", simple_interest_value)

}
