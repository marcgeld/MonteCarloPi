package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Using the Monte Carlo method to estimate the value of π.
//
// The Monte Carlo method is a statistical method that uses random sampling to estimate numerical results.
// In this case, we generate random points within a square and count how many fall within a unit circle.
// The ratio of points in the circle to the total points is used to estimate the value of π.
// The more points generated, the more accurate the estimate will be.
// The formula for the area of a circle is A = πr^2, where r = 1, so the area of the unit circle is π.
// The area of the square is 4, since the side length is 2.
// The ratio of the area of the circle to the area of the square is π/4, which is used to estimate π.
// The more points generated, the closer the estimate will be to the actual value of π.
// The Monte Carlo method is used in many fields, including physics, finance, and computer graphics.
// It is a powerful and versatile technique for estimating complex quantities using random sampling.
// The Monte Carlo method is named after the Monte Carlo Casino in Monaco, where the uncle of one of the inventors played roulette.
// The Monte Carlo method was first used during World War II to simulate the behavior of neutrons in nuclear reactions.
func main() {
	// Total number of points to generate
	const totalPoints = 1000000

	// Initialize the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Number of points that fall within the circle
	circlePoints := 0

	// Generate random points and count how many fall within the circle
	for i := 0; i < totalPoints; i++ {

		// Generate random x and y coordinates in the range 0…1
		x := rand.Float64()
		y := rand.Float64()

		// Check if the point falls within the circle (x^2 + y^2 <= 1)
		if math.Pow(x, 2)+math.Pow(y, 2) <= 1 {
			circlePoints++
		}
	}

	// Estimate the value of π using the ratio of points in the circle to the total points
	piEstimate := 4.0 * float64(circlePoints) / float64(totalPoints)

	fmt.Printf("estimate of π: %f (Using %d points)\n", piEstimate, totalPoints)
	fmt.Printf("actual value of π: %f\n", math.Pi)
}
