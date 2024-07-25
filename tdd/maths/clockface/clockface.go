package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func secondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30)

	return (math.Pi / (30 / float64(t.Second())))
}

// func zero() float64 {
// 	return 0.0
// }

// func main() {
// 	fmt.Println(10.0 / zero())
// }

func secondHandPoint(t time.Time) Point {
	// return Point{0, -1}

	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
