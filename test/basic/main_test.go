package basic

import (
	"testing"
)

// để biết 1 file test nào thì phải có tên file test là tên file cần test + _test.go

func TestAddONe(t *testing.T) {
	inputAddOne := 1
	outputAddOne := 2

	actual := AddOne(inputAddOne)
	if actual != outputAddOne {
		t.Errorf("AddOne(%d) = %d; want %d", inputAddOne, actual, outputAddOne)
	}

}
