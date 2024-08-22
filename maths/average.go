package maths

import "fmt"

func Average(sum float64, length float64) float64 {
	if length <= 0 {
		fmt.Println("Error with length value")
	}
	ave := sum / length
	return ave
}
