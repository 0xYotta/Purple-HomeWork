package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//asking for operation

	//asking for numbers
	//printing res

	usrChoise := getUserInput("Chose operation and enter it ('SUM', 'MED', 'AVG')")
	usrInput := getUserInput("Please enter numbers without spaces")

	res := calc(usrChoise, usrInput)

	fmt.Printf("%s; \n%#v; \n%.2f", usrChoise, convertStringToSliceOfInts(usrInput), res)

	fmt.Println()
	fmt.Println()
}

// Scanning for users input
func getUserInput(msg string) string {
	var input string
	fmt.Printf("%s ->  ", msg)
	fmt.Scan(&input)
	return input
}

// Converting []string to a []int and sorting it
func convertStringToSliceOfInts(input string) []int {
	tempSlice := strings.Split(input, ",")
	res := make([]int, len(tempSlice))

	for i, num := range tempSlice {
		res[i], _ = strconv.Atoi(num)
	}

	slices.Sort(res)
	return res
}

func avg(sli []int) float64 {

	if len(sli) == 0 {
		return 0
	}

	var sum float64
	for _, val := range sli {
		sum += float64(val)
	}

	sum /= float64(len(sli))
	return sum
}

func sum(sli []int) int {

	if len(sli) == 0 {
		return 0
	}

	var sum int
	for _, num := range sli {
		sum += num
	}
	return sum
}

//TODO
//Proccessing users input (commands: AVG, SUM, MED)

func med(sli []int) float64 {

	if len(sli) == 0 {
		return 0
	}

	if len(sli)%2 == 0 {
		index := len(sli)/2 - 1
		res := (float64(sli[index]) + float64(sli[index+1])) / 2
		return res
	}

	index := len(sli) / 2
	fmt.Println(index)
	return float64(sli[index])
}

func calc(usrChoise string, usrInput string) float64 {

	usrChoise = strings.ToLower(usrChoise)

	data := convertStringToSliceOfInts(usrInput)

	switch usrChoise {

	case "med":
		return med(data)

	case "avg":
		return avg(data)

	case "sum":
		return float64(sum(data))

	default:
		fmt.Println("Something went wrong. Exit...")
		return 0.0
	}

}
