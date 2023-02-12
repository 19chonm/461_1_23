package fileio

import (
	"testing"
	// "fmt"
	// "sort"
)

// I don't think I have to do this
func Test_MakeRatingsChannel_Success(t *testing.T) {
	rating := MakeRatingsChannel()
	if rating == nil {
		t.Errorf("MakeRatingChannelError occurred")
		// fmt.Printf(rating)
	}
}

// This don't work and I don't know why
// func Test_SortModule_Success(t *testing.T) {
// 	ch := MakeRatingsChannel() 
// 	sorted_ratings := Sort_modules(ch)
// 	// correct_ratings := ????
// 	if sorted_ratings == nil {
// 		t.Errorf("Sort Module Failed")
// 	}
// 	// if sorted_ratings != correct_ratings {
// 	// 	t.Errorf("Sort Module Failed")
// 	// }
// }
// func Test_SortModule_Fail(t *testing.T) {
// }

func Test_Make_json_string_Success(t *testing.T) {	
	goodUrl := "https://github.com/facebook/react"
	r := Rating{75, goodUrl, 5, 10, 15, 20, 25}
	jsonStringString := Make_json_string(r)

	if jsonStringString == "" {
		t.Errorf("got blank jsonStringString")
	}
}

func Test_Make_json_string_Fail(t *testing.T) {

}

// These... Are debug tests.. prob unneeded
// func Test_Print_sorted_output_Success(t *testing.T) {

// }
// func Test_Print_sorted_output_Fail(t *testing.T) {

// }
