package testfiles

import (
	"testing"

	"github.com/19chonm/461_1_23/fileio"
	"github.com/19chonm/461_1_23/worker"
)

/*
Test to make sure rating channel is filled up
*/

func TestWorkerGoodInput(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)
	worker.RunTask("https://github.com/expressjs/express", rating_ch)
	if len(rating_ch) == 0 {
		t.Errorf("rating channel was not updated")
	}
}

func TestWorkerPositiveScore(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)
	worker.RunTask("https://github.com/nullivex/nodist", rating_ch)
	for rating := range rating_ch {
		if rating.NetScore == 0 {
			t.Errorf("rating should be more than 0")
		}
	}
}

func TestWorkerBadInput(t *testing.T) {
	rating_ch := make(chan fileio.Rating, rating_ch_size)
	worker.RunTask("https://badurl.com/blabla/test", rating_ch)
	if len(rating_ch) != 0 {
		t.Errorf("rating channel should not have been updated")
	}
}

func TestWorkerRatingFail(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)

	//change this
	incorrectUrl := "https://github.com/incorrectownername/react"
	worker.RunTask(incorrectUrl, rating_ch)

	for rating := range rating_ch {
		if rating.NetScore != 0 {
			t.Errorf("rating should have been 0 for: %s", incorrectUrl)
		}
	}
}

func TestWorkerLicenseFail(t *testing.T) {
	rating_ch := make(chan fileio.Rating, 1)

	//change this
	incorrectUrl := "https://github.com/expressjs/express"
	worker.RunTask(incorrectUrl, rating_ch)

	for rating := range rating_ch {
		if rating.NetScore != 0 && rating.License == false {
			t.Errorf("rating should have been 0 for: %s", incorrectUrl)
		}
	}
}
