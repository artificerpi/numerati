package main

import (
	"fmt"
	"math"
)

type vector []float64

var (
	vec1 = vector{1, 2, 3, 4}
	vec2 = vector{3, 2, 4, 1.9}
)

func calcManhattanDist(v1, v2 vector) float64 {
	var length int = len(v1)
	if len(v2) > length {
		length = len(v2)
	}

	var sum float64
	for i := 0; i < length; i++ {
		sum += math.Abs(v1[i] - v2[i])
	}
	return sum
}

func calcEuclidianDist(v1, v2 vector) float64 {
	var length int = len(v1)
	if len(v2) > length {
		length = len(v2)
	}

	var sum float64
	for i := 0; i < length; i++ {
		sum += math.Abs(math.Pow(v1[i], 2) - math.Pow(v2[i], 2))
	}
	sum = math.Sqrt(sum / float64(length))
	return sum
}

func calcCosineDist(v1, v2 vector) float64 {
	var length int = len(v1)
	if len(v2) > length {
		length = len(v2)
	}

	var x, y, z float64
	for _, e := range v1 {
		x += math.Pow(e, 2)
	}
	for _, e := range v2 {
		y += math.Pow(e, 2)
	}
	for i := 0; i < length; i++ {
		z += v1[i] * v2[i]
	}

	return math.Sqrt(math.Pow(z, 2) / (x * y))

}

func main() {
	fmt.Println("Mahattan", calcManhattanDist(vec1, vec2))
	fmt.Println("Euclidian", calcEuclidianDist(vec1, vec2))
	fmt.Println("Cosine", calcCosineDist(vec1, vec2))
}
