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
			} else {
				continue
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
	// Compute the correctness score based on number of repository
	// clones, page views, and number of total commits
	var cs, vs, ms float64
	cs = 0.117 * (1 - math.Exp(-0.001*float64(clones)))   // 2k clones for 80%
	vs = 0.550 * (1 - math.Exp(-0.00002*float64(clones))) // 100k views for 86%
	ms = 0.333 * (1 - math.Exp(-0.0005*float64(clones)))  // 6000 commits for 90% of this

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
