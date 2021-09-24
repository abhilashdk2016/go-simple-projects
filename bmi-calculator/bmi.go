package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("BMI Calculator")
	fmt.Println("-------------------------------")
	fmt.Print("Please enter your weight (kg): ")
	weightInput, _ := reader.ReadString('\n')
	fmt.Print("Please enter your height (m): ")
	heightInput, _ := reader.ReadString('\n')
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)
	bmi := weight / (height * height)
	fmt.Printf("BMI is: %.2f\n", bmi)
}
