/*
2. Write a program to calculate the area of the circle, First line has a value of the radius of the circle
constraint
1. Use const PI from the package math package
2. Use the Pow function from the math package
3. Round area float value to 2 decimal places

*/

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Program to calcualte the area of circle ")
	var radius, area float64
	const pi_value = math.Pi

	fmt.Println("Enter the value of the radius")
	fmt.Scanln(&radius)

	area = pi_value * *&radius * radius

	// upto two decimal points
	area_of_circle_ := fmt.Sprintf("%.2f", area)

	fmt.Println("The area of the circle for the given radius is : ", area_of_circle_)

}
