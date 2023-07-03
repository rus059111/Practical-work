package main

import (
	"fmt"
	"math/rand"
	"time"
)

type point struct {
	x int
	y int
}

func generatePoints() [3]point {
	rand.Seed(time.Now().UnixNano())

	var points [3]point

	for i := 0; i < 3; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		p := point{x, y}
		points[i] = p
	}

	return points
}

func transformPoints(points [3]point) [3]point {
	var transformedPoints [3]point

	for i := 0; i < 3; i++ {
		x := 2*points[i].x + 10
		y := -3*points[i].y - 5
		p := point{x, y}
		transformedPoints[i] = p
	}

	return transformedPoints
}

func main() {
	points := generatePoints()

	fmt.Println("Сгенерированные точки:")
	for _, p := range points {
		fmt.Println(p)
	}

	transformedPoints := transformPoints(points)

	fmt.Println("Преобразованные точки:")
	for _, p := range transformedPoints {
		fmt.Println(p)
	}
}
