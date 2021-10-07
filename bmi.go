package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chimpaji/bmi-golang/info"
)

//bufio.NewReacder is for recieve data
//os.Stdin  is os standard input // our cmd




func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Seperator)

	//take wieight and height 
	
	fmt.Print(info.WeightPrompt)
	weightInput , _ := reader.ReadString('\n') // read all string until enter key(already included)
 
	
	fmt.Print(info.HeightPrompt)
	heightInput , _ := reader.ReadString('\n')

	weightInput = "85\n"
	heightInput = "1.75\n"

	//Save that input in variable (delete the line break)
	weightInput = strings.Replace(weightInput,"\n","",-1)
	heightInput = strings.Replace(heightInput,"\n","",-1)

	//Print not give line break
	// fmt.Print(weightInput)

	//float64 cant convert string into number, use the 'strconv' package instead!
	weight , _ := strconv.ParseFloat(weightInput, 64)
	height , _ := strconv.ParseFloat(heightInput, 64)

	//Calculate the BMI(weight /(h*h))
	bmi := weight / (height*height)

	fmt.Printf("Your BMI: %.2f",bmi)





	
	//calculate the bmi (w/(h*h))
	//output the BMI

}