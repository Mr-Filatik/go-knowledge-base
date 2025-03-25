package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanSliceCapacityIncrease_CorrectData(t *testing.T) {
	result := ScanSliceCapacityIncrease(4, 10)

	expected := []int{4, 8}
	assert.Equal(t, expected, result)
}

func TestScanSliceCapacityIncrease_CorrectData_FromZero(t *testing.T) {
	result := ScanSliceCapacityIncrease(0, 10)

	expected := []int{0, 1, 2, 4, 8}
	assert.Equal(t, expected, result)
}

func TestScanSliceCapacityIncrease_IncorrectData(t *testing.T) {
	result := ScanSliceCapacityIncrease(4, 0)

	expected := []int{4}
	assert.Equal(t, expected, result)
}
