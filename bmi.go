package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//cant use const because the functino before program start dont have any value now, unlike js
var reader = bufio.NewReader(os.Stdin)
//bufio.NewReacder is for recieve data 
//os.Stdin  is os standard input // our cmd

func main() {
	fmt.Println("Lets calculate BMI")
	fmt.Println("-------------------")

	//take wieight and height 
	
	fmt.Print("Please enter your weight (kg) : ")
	weightInput , _ := reader.ReadString('\n') // read all string until enter key(already included)
 
	
	fmt.Print("Please enter your height (m) : ")
	heightInput , _ := reader.ReadString('\n')

	weightInput = "85\n"
	heightInput = "1.75\n"

	//Save that input in variable (delete the line break)
	weightInput = strings.Replace(weightInput,"\n","",-1)
	heightInput = strings.Replace(heightInput,"\n","",-1)
	fmt.Printf("Your height: %v",heightInput)

	//Print not give line break
	// fmt.Print(weightInput)

	//float64 cant convert string into number, use the 'strconv' package instead!
	weight , _ := strconv.ParseFloat(weightInput, 64)
	height , _ := strconv.ParseFloat(heightInput, 64)
	fmt.Printf("Your height: %.2f",height)

	//Calculate the BMI(weight /(h*h))
	bmi := weight / (height*height)

	fmt.Print("hello")
	fmt.Printf("Your BMI: %.2f",bmi)





	
	//calculate the bmi (w/(h*h))
	//output the BMI

}