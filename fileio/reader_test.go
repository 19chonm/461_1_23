package fileio

import (
	"testing"
)

// func Test_(t *testing.T) {

// }


func Test_MakeUrlChannel_Success(t *testing.T) {
	result := MakeUrlChannel() 
	if result == nil {
		t.Errorf("MakeUrlChannel failed, returned nil")
	}
}