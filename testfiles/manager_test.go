package testfiles

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/worker"
)

var url_ch_size = 100
var rating_ch_size = 5

/*
testing for rating channel and workers
expected behavior for StartWorkers():
rating channel should be closed, tasks should be run filling up rating channel
*/
func TestManagerGoodInput(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	url_ch <- "https://github.com/expressjs/express"
	url_ch <- "https://github.com/facebook/react"
	url_ch <- "https://github.com/request/request"

	worker.StartWorkers(url_ch, rating_ch)

	// if can still be read from, rating channel is not closed
	_, ok := (<-rating_ch)
	if ok {
		t.Errorf("rating channel was not closed")
	}
}

func TestManagerCreatesRatings(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	url_ch <- "https://github.com/lodash/lodash"
	url_ch <- "https://github.com/nullivex/nodist"
	url_ch <- "https://github.com/cloudinary/cloudinary_npm"

	worker.StartWorkers(url_ch, rating_ch)

	// comparison to default values of rating type to check
	// if they were untouched
	for rating := range rating_ch {
		if rating.Busfactor == 0 && rating.Correctness == 0 && rating.License == false && rating.NetScore == 0 && rating.Rampup == 0 && rating.Responsiveness == 0 {
			t.Errorf("ratings were not created correctly")
		}
	}
}

func TestManagerBadInput(t *testing.T) {
	url_ch := make(chan string, url_ch_size)
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	url_ch <- "https://badurl.com/blabla/test"
	url_ch <- "incorrecturl!!@#$**F(S)"

	worker.StartWorkers(url_ch, rating_ch)

	//testing logic
	for rating := range rating_ch {
		if rating.Busfactor != 0 && rating.Correctness != 0 && rating.License != false && rating.NetScore != 0 && rating.Rampup != 0 && rating.Responsiveness != 0 {
			t.Errorf("ratings should not have been created")
		}
	}
}

// removal because manager should expect npm urls
// func TestManagerNPMUrl(t *testing.T) {
// 	url_ch := make(chan string, url_ch_size)
// 	rating_ch := make(chan fileio.Rating, rating_ch_size)
// 	url_ch <- "https://www.npmjs.com/package/react"
// 	url_ch <- "https://www.npmjs.com/package/express"

// 	worker.StartWorkers(url_ch, rating_ch)

// 	//testing logic
// 	for rating := range rating_ch {
// 		if rating.Busfactor != 0 && rating.Correctness != 0 && rating.License != false && rating.NetScore != 0 && rating.Rampup != 0 && rating.Responsiveness != 0 {
// 			t.Errorf("ratings should not have been created")
// 		}
// 	}
// }
