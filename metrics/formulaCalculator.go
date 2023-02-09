package metrics

import (
	"math"
)

type Factor struct {
	Weight       float64
	Value        float64
	AllOrNothing bool
}

func ComputeNetScore(fs []Factor) float64 {
	var sum float64

	for _, f := range fs {
		if f.AllOrNothing {
			if f.Value == 0 {
				sum = 0
				break
			}
		}

		sum += f.Value * f.Weight
	}

	return sum
}

func ComputeRampTime(found int, total int) float64 {
	// Compute Ramp-up time based on number of phrases found in README
	if total == 0 {
		return 0
	}
	return float64(found) / float64(total)
}

func ComputeCorrectness(clones int, views int, commits int) float64 {
	// Correctness is determined by a sum three factors: clones, views, and commits
	// Each of these are calculated using a an exponential decay function to ensure that
	// the domain is from -inf to inf, while the range is still between 0 and the weight (0.117, 0.550, or 0.333)
	// As a rough benchmark, I determined the quantity of each metric to reach a certain output value.

	// Example:
	// For clones, the weight is 0.117. To get 80% of that weight, we need the repository to have 2000 clones
	// The result is cs = 0.117 * 0.8 = 0.0936

	var cs, vs, ms float64
	cs = 0.117 * (1 - math.Exp(-0.001*float64(clones)))   // 2k clones for 80%
	vs = 0.550 * (1 - math.Exp(-0.00002*float64(views)))  // 100k views for 86%
	ms = 0.333 * (1 - math.Exp(-0.0005*float64(commits))) // 6000 commits for 90% of this

	return cs + vs + ms
}

func ComputeResponsiveness(days float64) float64 {
	// Compute the responsiveness score based on average
	// number of days to fix bug issues
	if days < 0 {
		return 1
	}

	return math.Exp(-0.05 * float64(days))
}

func ComputeBusFactor(top int, total int) float64 {
	// Compute the Bus factor by measuring the percentage of commits
	// in the past year committed by the top three performers
	if total <= 0 {
		return 0
	}

	return 1 - (float64(top) / float64(total))
}

func ComputeLicenseScore(license string) int {
	if license == "MIT License" {
		return 1
	} else {
		return 0
	}
}
