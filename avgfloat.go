package main

import "fmt"

func main() {
	arrFloat := []float64{38.34, 22.32, 23.34}
	fmt.Printf("%f\n", average(arrFloat))
}

func average(xs []float64) (avg float64) { //<1>
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0.0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return // an interesting thing. avg is defined as return value, its type is declared, hence no need to be explicit here
}
