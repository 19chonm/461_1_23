package fileio

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// 					--- NDJSON ---
// Each JSON text MUST conform to the [RFC8259]
// standard and MUST be written to the stream followed
// by the newline character \n (0x0A). The newline character
// MAY be preceded by a carriage return \r (0x0D). The JSON
// texts MUST NOT contain newlines or carriage returns.
//
// All serialized data MUST use the UTF8 encoding.
// https://github.com/ndjson/ndjson-spec

// ND JSON follows RFC 8259
// Go's JSON library follows RFC 7159
// - However, the only major change between the two is that 8259 supports UTF8, which Go does by default
// - Assuming this different is negligible, and that using Go's "json" library is okay

const worker_output_ch_size = 100 // Size of the buffer for the Worker's output

type Rating struct {
	NetScore       float64 `json:"NetScore"`
	Url            string  `json:"URL"`
	License        float64 `json:"License"`
	Rampup         float64 `json:"RampUp"`
	Correctness    float64 `json:"Correctness"`
	Responsiveness float64 `json:"ResponsiveMaintainer"`
	Busfactor      float64 `json:"BusFactor"`
}

type WorkerOutput struct {
	WorkerRating Rating
	WorkerErr    error
}

func MakeWorkerOutputChannel() chan WorkerOutput {
	return make(chan WorkerOutput, worker_output_ch_size)
}

func ReadWorkerResults(ch chan WorkerOutput) ([]Rating, []error) {
	// Create a slice to hold the values from the channel
	sorted_ratings := []Rating{}
	errors := []error{}

	// Read in ratings from channel
	for {
		wo, ok := <-ch
		if !ok { // Channel has been closed
			break
		}

		if wo.WorkerErr != nil { // If error, record the error
			errors = append(errors, wo.WorkerErr)
		} else { // Else, record the result
			sorted_ratings = append(sorted_ratings, wo.WorkerRating)
		}
	}

	// Sort the slice
	sort.Slice(sorted_ratings, func(p, q int) bool {
		return sorted_ratings[p].NetScore > sorted_ratings[q].NetScore
	})

	return sorted_ratings, errors
}

func make_json_string(r Rating) string {
	// Convert the Rating struct into a json string
	jsonString, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("for: %+v\n", r)
		panic("fileio: make_json_string fail")
	}

	return string(jsonString)
}

func Print_sorted_output(ratings []Rating) {
	fmt.Println("\n\n----------------Sorted Ratings-----------------")
	for r := range ratings {
		fmt.Println(ratings[r].Url, "has a rating of:", make_json_string(ratings[r]))
	}
	fmt.Println("-----------------------------------------------")
}

func PrintErrors(errs []error) {
	fmt.Fprintln(os.Stderr, "Errors in the worker stage: ")
	for _, e := range errs {
		fmt.Fprintln(os.Stderr, e)
	}
}
