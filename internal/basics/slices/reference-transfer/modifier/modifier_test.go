package modifier

import (
	"reflect"
	"testing"
)

func TestModifySliceDirectly(t *testing.T) {
	originalSlice := []int{1, 2, 3}

	ModifySliceDirectly(originalSlice, func(x int) int {
		return x * 2
	})

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(originalSlice, expected) {
		t.Errorf("TestModifySliceDirectly failed: actual %v, expected %v", originalSlice, expected)
	}
}

func TestModifySliceReference(t *testing.T) {
	originalSlice := []int{1, 2, 3}

	ModifySliceReference(&originalSlice, func(x int) int {
		return x * 2
	})

	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(originalSlice, expected) {
		t.Errorf("TestModifySliceReference failed: actual %v, expected %v", originalSlice, expected)
	}
}
