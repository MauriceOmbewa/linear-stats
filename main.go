package main

import (
	"bufio"
	"fmt"
	"linear-stats/maths"
	"math"
	"os"
	"strconv"
)

func main() {
	var ysum float64
	var xaxis []float64

	if len(os.Args) != 2 {
		fmt.Println("Error: Incorrect number of arguments")
		return
	}

	inputfilename := os.Args[1]
	inputfile, err := os.Open(inputfilename)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer inputfile.Close()

	fileinfo, err := inputfile.Stat()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if fileinfo.Size() == 0 {
		fmt.Println("Error: File is empty")
		return
	}

	scanner := bufio.NewScanner(inputfile)

	var data []float64

	for scanner.Scan() {
		new := scanner.Text()
		if new == ""{
			continue
		}

		atc, err := strconv.ParseFloat(new, 64)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		data = append(data, float64(atc))
	}
	yaxis := data
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := 0; i < len(data); i++ {
		ysum += data[i]
		xaxis = append(xaxis, float64(i))
	}

	var xsum float64
	for i := 0; i < len(xaxis); i++ {
		xsum += xaxis[i]
	}

	length := float64(len(data))
	xmean := maths.Average(xsum, length)
	ymean := maths.Average(ysum, length)

	// Calculate Pearson Correlation Coefficient
	numerator := 0.0
	denominatorX := 0.0
	denominatorY := 0.0

	for i := 0; i < len(xaxis); i++ {
		xDiff := xaxis[i] - xmean
		yDiff := yaxis[i] - ymean

		numerator += xDiff * yDiff
		denominatorX += xDiff * xDiff
		denominatorY += yDiff * yDiff
	}

	denominator := math.Sqrt(denominatorX * denominatorY)
	r := numerator / denominator

	// Calculate Linear Regression Coefficients
	a, b := maths.Linearregression(xaxis, yaxis)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", a, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

