package main

import (
	"fmt"
	"math"
)

type factor struct {
	weight       float64
	value        float64
	allOrNothing bool
}

func computeNetScore(fs []factor) float64 {
	var sum float64

	for _, f := range fs {
		if f.allOrNothing {
			if f.value == 0 {
				sum = 0
				break
			} else {
				continue
			}
		}

		sum += f.value * f.weight
	}

	return sum
}

func computeRampTime(found int, total int) float64 {
	// Compute Ramp-up time based on number of phrases found in README
	return float64(found) / float64(total)
}

func computeCorrectness(clones int, views int, commits int) float64 {
	// Compute the correctness score based on number of repository
	// clones, page views, and number of total commits
	var cs, vs, ms float64
	cs = 0.117 * (1 - math.Exp(-0.001*float64(clones)))   // 2k clones for 80%
	vs = 0.550 * (1 - math.Exp(-0.00002*float64(clones))) // 100k views for 86%
	ms = 0.333 * (1 - math.Exp(-0.0005*float64(clones)))  // 6000 commits for 90% of this

	return cs + vs + ms
}

func computeResponsiveness(days float64) float64 {
	// Compute the responsiveness score based on average
	// number of days to fix bug issues
	if days < 0 {
		return 1
	}

	return math.Exp(-0.05 * float64(days))
}

func computeBusFactor(total int, p1 int, p2 int, p3 int) float64 {
	// Compute the Bus factor by measuring the percentage of commits
	// in the past year committed by the top three performers
	return 1 - (float64(p1+p2+p3) / float64(total))
}

func testCompute() {
	// Test function to see if all the other functions work
	var temp = []factor{
		factor{
			// URL
			weight:       0,
			value:        0,
			allOrNothing: true,
		},
		factor{
			// License
			weight:       0,
			value:        0,
			allOrNothing: true,
		},
		factor{
			// RampTime
			weight:       0.15,
			value:        0,
			allOrNothing: false,
		},
		factor{
			// StandardOfCorrectness
			weight:       0.15,
			value:        0,
			allOrNothing: false,
		},
		factor{
			// Responsiveness
			weight:       0.4,
			value:        0,
			allOrNothing: false,
		},
		factor{
			// BusFactor
			weight:       0.3,
			value:        0,
			allOrNothing: false,
		},
	}

	temp[0].value = 1
	temp[1].value = 1
	temp[2].value = computeRampTime(4, 5)
	temp[3].value = computeCorrectness(10000, 100000, 1000)
	temp[4].value = computeResponsiveness(14)
	temp[5].value = computeBusFactor(100, 9, 7, 3)

	res := computeNetScore(temp)
	fmt.Println(res)

}
