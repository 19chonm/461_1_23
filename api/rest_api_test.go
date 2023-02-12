package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	// "strings"
	"bytes"
	// "encoding/json"
)

type TestType struct {
	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}

// {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}
// Input URL Tests
func TestGoodInput(t *testing.T) {
	var goodInputUrl string = "https://github.com/facebook/react"
	var correctUser string = "facebook"
	var correctRepo string = "react"
	correctToken, _ := os.LookupEnv("GITHUB_TOKEN")
	user, repo, token, ok := ValidateInput(goodInputUrl)
	if user != correctUser {
		t.Errorf("user got: %s, want: %s.", user, correctUser)
	}
	if repo != correctRepo {
		t.Errorf("repo got: %s, want: %s.", repo, correctRepo)
	}
	if token != correctToken {
		t.Errorf("token got: %s, want: %s.", token, correctToken)
	}
	if ok != nil {
		t.Errorf("ok was not nil: %s", ok.Error())
	}
}

func TestBadInput(t *testing.T) {
	var badInputUrl string = "https://google.com/facebok/reat"
	var badUser string = ""
	var badRepo string = ""
	var badToken string = ""
	badOk := fmt.Errorf("some api error")

	user, repo, token, ok := ValidateInput(badInputUrl)
	if user != badUser {
		t.Errorf("user got: %s, want: %s.", user, badUser)
	}
	if repo != badRepo {
		t.Errorf("repo got: %s, want: %s.", repo, badRepo)
	}
	if token != badToken {
		t.Errorf("token got: %s, want: %s.", token, badToken)
	}
	if ok == nil {
		t.Errorf("ok got: %s, want: %s.", ok, badOk.Error())
	}
}

func TestDecodeResponse(t *testing.T) {
	res := http.Response{
		Body: io.NopCloser(bytes.NewBufferString("{\"foo\": 461, \"bar\": \"Project\"}")),
	}
	correctFoo := 461
	correctBar := "Project"
	jsonRes, err := DecodeResponse[TestType](&res)

	if jsonRes.Foo != correctFoo {
		t.Errorf("jsonRes.Foo got: %d want: %d.", jsonRes.Foo, correctFoo)
	}
	if jsonRes.Bar != correctBar {
		t.Errorf("jsonRes.Bar got: %s want: %s.", jsonRes.Bar, correctBar)
	}
	if err != nil {
		t.Errorf("err got: %v", err)
	}
}
