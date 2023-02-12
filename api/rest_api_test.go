package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	// "strings"
	"bytes"
	"net/url"
	// "encoding/json"
)

type TestType struct {
	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}

// {"license":{"key":"mit","name":"MIT License","url":"https://api.github.com/licenses/mit"}}
// Input URL Tests
func Test_Input_Success(t *testing.T) {
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

func Test_Input_BadURL(t *testing.T) {
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

func Test_DecodeResponse(t *testing.T) {
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

func Test_SendGithubRequestHelper_Success(t *testing.T) {
	endpoint := "https://api.github.com/users/octocat/orgs"
	token, _ := os.LookupEnv("GITHUB_TOKEN")

	// retry_count := 0
	res, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if res == nil {
		t.Errorf("res is nil")
	}

	if err != nil {
		t.Errorf("Got unexpected error %s", err.Error())
	} 
	if (statusCode != 200) {
		t.Errorf("GitHub request responded with error code %d", statusCode)
	}
}

func Test_SendGithubRequestHelper_BadEndpoint(t *testing.T) {
	endpoint := "https://api.github.com/bad_endpoint"
	token, _ := os.LookupEnv("GITHUB_TOKEN")

	// retry_count := 0
	_, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if err == nil {
		t.Errorf("err is nil")
	} 
	if (statusCode == 200) {
		t.Errorf("Got 200 status code")
	}
}

func Test_SendGithubRequestHelper_NotFound(t *testing.T) {
	endpoint := "https://api.github.com/repos/fakeuser/fakerepohopefully"
	token, _ := os.LookupEnv("GITHUB_TOKEN")

	// retry_count := 0
	_, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if err == nil {
		t.Errorf("err is nil")
	} 
	if (statusCode != 404) {
		t.Errorf("want 404 status code, got %d", statusCode)
	}
}

func Test_SendGithubRequestHelper_BadToken(t *testing.T) {
	endpoint := "https://api.github.com/user/octocat/orgs"
	token := "invalid_token"

	// retry_count := 0
	_, err, statusCode := SendGithubRequestHelper(endpoint, token)

	if err == nil {
		t.Errorf("err is nil")
	} 
	if (statusCode == 200) {
		t.Errorf("Got 200 status code")
	}
}

func Test_SetQueryParameter_SuccessNotExists(t *testing.T) {
	endpoint := "https://example.com/path?foo=bar"
	SetQueryParameter(&endpoint, "baz", "quux")
	urlObject, _ := url.Parse(endpoint)
	query := urlObject.Query()
	
	if query.Get("baz") != "quux" {
		t.Errorf("want baz to be quux, got %s", query.Get("baz"))
	}
	if query.Get("foo") != "bar" {
		t.Errorf("want foo to be bar, got %s", query.Get("foo"))
	}
}

func Test_SetQueryParameter_SuccessExists(t *testing.T) {
	endpoint := "https://example.com/path?foo=bar"
	SetQueryParameter(&endpoint, "foo", "baz")
	urlObject, _ := url.Parse(endpoint)
	query := urlObject.Query()

	if query.Get("foo") != "baz" {
		t.Errorf("want foo to be bar, got %s", query.Get("foo"))
	}
}

func Test_SetQueryParameter_Error(t *testing.T) {
	endpoint := "\n" // bad endpoint
	SetQueryParameter(&endpoint, "foo", "bar")
	_, err := url.Parse(endpoint)

	if err == nil {
		t.Errorf("err is nil")
	}
}
