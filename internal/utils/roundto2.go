package utils

import (
	// "fmt"
	"math"
)

func RoundTo2(x float64) float64 {
	// y := math.Round(x)
	// fmt.Println("this is y", y)
	// fmt.Println("and this is with the division", math.Round(x*100)/100)
	return math.Round(x*100) / 100
}
