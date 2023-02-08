package worker

import (
	"fmt"

	"github.com/19chonm/461_1_23/api"
	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/metrics"
)

func runTask(url string, ratingch chan<- fileio.Rating) {
	fmt.Println("My job is", url)

	l := api.GetRepoLicense(url)
	fmt.Println("license: ", *l.License.Name)
	a := api.GetRepoAverageLifespan(url)
	fmt.Println("lifespan: ", a)
	c := api.GetRepoContributors(url)
	fmt.Println("contributors: ", c)

	// rating := fileio.Rating{NetScore: rand.Float64(), Rampup: rampupscore, Url: url} // Placeholder until real data implemented
	// ratingch <- rating

	rampup_score := metrics.ScanRepo(url)
	correctness_score := metrics.ComputeCorrectness(0, 0, 0)
	responsiveness_score := metrics.ComputeResponsiveness(0)
	busfactor_score := metrics.ComputeBusFactor(0, 0, 0, 0)
	license_score := metrics.ComputeLicenseScore("MIT License")

	rampup_factor := metrics.Factor{Weight: 0.15, Value: rampup_score, AllOrNothing: false}
	correctness_factor := metrics.Factor{Weight: 0.15, Value: correctness_score, AllOrNothing: false}
	responsiveness_factor := metrics.Factor{Weight: 0.4, Value: responsiveness_score, AllOrNothing: false}
	busfactor_factor := metrics.Factor{Weight: 0.3, Value: busfactor_score, AllOrNothing: false}
	license_factor := metrics.Factor{Weight: 1.0, Value: float64(license_score), AllOrNothing: true}

	// Assuming only valid urls get passed through
	// url_factor := metrics.Factor{Weight: 1.0, Value: float64(url_score), AllOrNothing: true}

	factors := []metrics.Factor{rampup_factor, correctness_factor, responsiveness_factor, busfactor_factor, license_factor}

	r := fileio.Rating{NetScore: metrics.ComputeNetScore(factors),
		Rampup:         rampup_score,
		Url:            url,
		License:        float64(license_score),
		Busfactor:      busfactor_score,
		Responsiveness: responsiveness_score,
		Correctness:    correctness_score,
	}
	ratingch <- r
}
