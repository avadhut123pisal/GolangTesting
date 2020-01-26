package calc

import (
	"testing"
)

type TestDataItem struct {
	inputs         []int
	expectedResult int
	hasError       bool
}

func TestMathAdd(t *testing.T) {
	testCases := []TestDataItem{
		{
			[]int{1, 2, 3, 4}, 10, false,
		},
		{
			[]int{-1, -2}, -3, false,
		},
		{
			[]int{1}, 0, true,
		},
	}
	for _, data := range testCases {
		sum, err := Add(data.inputs...)
		if data.hasError {
			if err == nil {
				t.Errorf("Add() with args %v: FAILED, expected error but got value %v", data.inputs, sum)
			} else {
				t.Logf("Add() with args %v: PASSED, expected error, got error %v", data.inputs, err)
			}
		} else {
			if sum != data.expectedResult {
				t.Errorf("Add() with args %v: FAILED, expected %v but got %v", data.inputs, data.expectedResult, sum)
			} else {
				t.Logf("Add() with args %v: PASSED, expected %v but got %v", data.inputs, data.expectedResult, sum)
			}
		}
	}

}
