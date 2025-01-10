package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestRequire(t *testing.T) {
	fmt.Println("Test 2 test case")
	require.Equal(t, 2, AddOne(1))
	fmt.Println("AddOne test 1 passed")
	require.Equal(t, 2, AddOne(2))
	fmt.Println("AddOne test 2 passed")
}

func TestAssert(t *testing.T) {
	fmt.Println("Test 2 test case")
	assert := assert.New(t)
	assert.Equal(t, 2, AddOne(1))
	fmt.Println("AddOne test 1")
	assert.Equal(t, 2, AddOne(2))
	fmt.Println("AddOne test 2")

}
