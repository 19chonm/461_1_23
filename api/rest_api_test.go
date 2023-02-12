package api

import (
	"fmt"
	"testing"
	"os"
	// api 
)
	// {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}
// Input URL Tests
func TestGoodInput(t *testing.T) {
	var goodInputUrl string = "https://github.com/qiangxue/go-rest-api"
	var correctUser string = ""
	var correctRepo string = ""
	correctToken, _ := os.LookupEnv("GITHUB_TOKEN")

	user, repo, token, ok := ValidateInput(goodInputUrl)
	if user != "" {
		t.Errorf("user got: %s, want: %s", user, correctUser)
	} 
	if repo != "" {
		t.Errorf("repo got: %s, want: %s", user, correctRepo)
	}
	if token != "" {
		t.Errorf("token got: %s, want: %s", user, correctToken)
	}
	if ok != nil {
		t.Errorf("ok was not nil: %s", ok.Error())
	}
}

func TestBadInput(t *testing.T) {
	var badInputUrl string = ""
	var badUser string = "badUser"
	var badRepo string = "badRepo"
	var badToken string = "badToken"
	badOk := fmt.Errorf("blah")

	user, repo, token, ok := ValidateInput(badInputUrl)
	if user != "" {
		t.Errorf("user got: %s, want: %s", user, badUser)
	} 
	if repo != "" {
		t.Errorf("repo got: %s, want: %s", user, badRepo)
	}
	if token != "" {
		t.Errorf("token got: %s, want: %s", user, badToken)
	}
	if ok != badOk {
		t.Errorf("ok got: %s, want: %s", ok.Error(), badOk.Error())
	}
}

