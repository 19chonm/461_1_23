package main

import (
	"encoding/json"
	"fmt"
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

type rating struct {
	NetScore       float64 `json:"NetScore"`
	Url            string  `json:"URL"`
	License        bool    `json:"License"`
	Rampup         float64 `json:"RampUp"`
	Correctness    float64 `json:"Correctness"`
	Responsiveness float64 `json:"ResponsiveMaintainer"`
	Busfactor      float64 `json:"BusFactor"`
}

func output_row(r rating) string {
	jsonString, err := json.Marshal(r)
	if err != nil {
		panic("bad error")
	}

	return string(jsonString) + "\n"
}

func NDJSON_test() {
	datas := make([]rating, 10)

	for i := 0; i < 10; i++ {
		datas[i] = rating{NetScore: float64(i), License: (i != 0)}
	}

	for _, data := range datas {
		// data := rating{NetScore: 1}
		o := output_row(data)
		fmt.Print(o)
	}

	// jsonString, err := json.Marshal(data)
	// fmt.Println(err)

	// fmt.Println(data)
	// fmt.Println(string(jsonString))
}

func Sort_modules(ratings chan rating) []rating {
	// create a slice to hold the values from the channel
	sorted_ratings := []rating{}
	for r := range ratings {
		sorted_ratings = append(sorted_ratings, r)
	}

	// sort the slice
	sort.Slice(sorted_ratings, func(p, q int) bool {
		return sorted_ratings[p].NetScore > sorted_ratings[q].NetScore
	})

	Print_sorted_output(sorted_ratings)
	
	return sorted_ratings
}

func Print_sorted_output(ratings []rating) {
	fmt.Println("----------------Sorted Ratings-----------------")
	for r := range ratings {
		fmt.Println(ratings[r].Url, "has a rating of:", ratings[r].NetScore)
	}
	fmt.Println("-----------------------------------------------")
}