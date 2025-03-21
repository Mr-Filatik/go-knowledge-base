package scanner

import (
	"testing"

	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/entity"
	"github.com/stretchr/testify/assert"
)

type emptyStruct struct{}

func TestScanStructInfo_EmptyStruct(t *testing.T) {
	testInstance := emptyStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_EmptyStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_EmptyStruct",
		StructType:  "emptyStruct",
		StructSize:  0,
		StructAlign: 8,
		FieldInfos:  []entity.FieldInfo{},
	}
	assert.Equal(t, expected, result)
}

type singleFieldStruct struct {
	Field1 int64
}

func TestScanStructInfo_SingleFieldStruct(t *testing.T) {
	testInstance := singleFieldStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_SingleFieldStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_SingleFieldStruct",
		StructType:  "singleFieldStruct",
		StructSize:  8,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int64",
				FieldOffset:  0,
				FieldSize:    8,
				FieldPadding: 0,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type pointerStruct struct {
	Field1 int32
	Field2 *int
	Field3 int8
}

func TestScanStructInfo_PointerStruct(t *testing.T) {
	testInstance := pointerStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_PointerStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_PointerStruct",
		StructType:  "pointerStruct",
		StructSize:  24,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int32",
				FieldOffset:  0,
				FieldSize:    4,
				FieldPadding: 4,
			},
			{
				FieldName:    "Field2",
				FieldType:    "*int",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int8",
				FieldOffset:  16,
				FieldSize:    1,
				FieldPadding: 7,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type arrayStruct struct {
	Field1 [3]int16
	Field2 int32
	Field3 int8
}

func TestScanStructInfo_ArrayStruct(t *testing.T) {
	testInstance := arrayStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_ArrayStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_ArrayStruct",
		StructType:  "arrayStruct",
		StructSize:  16,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "[3]int16",
				FieldOffset:  0,
				FieldSize:    6,
				FieldPadding: 2,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int32",
				FieldOffset:  8,
				FieldSize:    4,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int8",
				FieldOffset:  12,
				FieldSize:    1,
				FieldPadding: 3,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type sliceStruct struct {
	Field1 []int32
	Field2 int32
	Field3 int8
}

func TestScanStructInfo_SliceStruct(t *testing.T) {
	testInstance := sliceStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_SliceStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_SliceStruct",
		StructType:  "sliceStruct",
		StructSize:  32,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "[]int32",
				FieldOffset:  0,
				FieldSize:    24,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int32",
				FieldOffset:  24,
				FieldSize:    4,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int8",
				FieldOffset:  28,
				FieldSize:    1,
				FieldPadding: 3,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type mapStruct struct {
	Field1 int64
	Field2 map[string]int32
	Field3 int32
}

func TestScanStructInfo_MapStruct(t *testing.T) {
	testInstance := mapStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_MapStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_MapStruct",
		StructType:  "mapStruct",
		StructSize:  24,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int64",
				FieldOffset:  0,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "map[string]int32",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int32",
				FieldOffset:  16,
				FieldSize:    4,
				FieldPadding: 4,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type nestedInnerStruct struct {
	InnerField1 int8
	InnerField2 int32
}

type nestedOuterStruct struct {
	Field1 int16
	Field2 nestedInnerStruct
	Field3 int8
}

func TestScanStructInfo_NestedStruct(t *testing.T) {
	testInstance := nestedOuterStruct{
		Field2: nestedInnerStruct{},
	}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_NestedStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_NestedStruct",
		StructType:  "nestedOuterStruct",
		StructSize:  16,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int16",
				FieldOffset:  0,
				FieldSize:    2,
				FieldPadding: 2,
			},
			{
				FieldName:    "Field2",
				FieldType:    "nestedInnerStruct",
				FieldOffset:  4,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int8",
				FieldOffset:  12,
				FieldSize:    1,
				FieldPadding: 3,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type boolStruct struct {
	Field1 bool
	Field2 int16
	Field3 bool
}

func TestScanStructInfo_BoolStruct(t *testing.T) {
	testInstance := boolStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_BoolStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_BoolStruct",
		StructType:  "boolStruct",
		StructSize:  6,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "bool",
				FieldOffset:  0,
				FieldSize:    1,
				FieldPadding: 1,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int16",
				FieldOffset:  2,
				FieldSize:    2,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "bool",
				FieldOffset:  4,
				FieldSize:    1,
				FieldPadding: 1,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type stringStruct struct {
	Field1 int64
	Field2 string
	Field3 int32
}

func TestScanStructInfo_StringStruct(t *testing.T) {
	testInstance := stringStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_StringStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_StringStruct",
		StructType:  "stringStruct",
		StructSize:  32,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int64",
				FieldOffset:  0,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "string",
				FieldOffset:  8,
				FieldSize:    16,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int32",
				FieldOffset:  24,
				FieldSize:    4,
				FieldPadding: 4,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type overlapStruct struct {
	Field1 int32
	Field2 int8
	Field3 int64
}

func TestScanStructInfo_OverlapStruct(t *testing.T) {
	testInstance := overlapStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_OverlapStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_OverlapStruct",
		StructType:  "overlapStruct",
		StructSize:  16,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int32",
				FieldOffset:  0,
				FieldSize:    4,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int8",
				FieldOffset:  4,
				FieldSize:    1,
				FieldPadding: 3,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int64",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type aligned8Struct struct {
	Field1 int32
	Field2 int32
	Field3 int8
}

func TestScanStructInfo_Aligned8Struct(t *testing.T) {
	testInstance := aligned8Struct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_Aligned8Struct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_Aligned8Struct",
		StructType:  "aligned8Struct",
		StructSize:  12,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int32",
				FieldOffset:  0,
				FieldSize:    4,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int32",
				FieldOffset:  4,
				FieldSize:    4,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int8",
				FieldOffset:  8,
				FieldSize:    1,
				FieldPadding: 3,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type aligned16Struct struct {
	Field1 int64
	Field2 int64
	Field3 int8
}

func TestScanStructInfo_Aligned16Struct(t *testing.T) {
	testInstance := aligned16Struct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_Aligned16Struct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_Aligned16Struct",
		StructType:  "aligned16Struct",
		StructSize:  24,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int64",
				FieldOffset:  0,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int64",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int8",
				FieldOffset:  16,
				FieldSize:    1,
				FieldPadding: 7,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type unorderedStruct struct {
	Field1 int8
	Field2 int64
	Field3 int16
}

func TestScanStructInfo_UnorderedStruct(t *testing.T) {
	testInstance := unorderedStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_UnorderedStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_UnorderedStruct",
		StructType:  "unorderedStruct",
		StructSize:  24,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int8",
				FieldOffset:  0,
				FieldSize:    1,
				FieldPadding: 7,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int64",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int16",
				FieldOffset:  16,
				FieldSize:    2,
				FieldPadding: 6,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type orderedStruct struct {
	Field1 int8
	Field2 int16
	Field3 int64
}

func TestScanStructInfo_OrderedStruct(t *testing.T) {
	testInstance := orderedStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_OrderedStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_OrderedStruct",
		StructType:  "orderedStruct",
		StructSize:  16,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int8",
				FieldOffset:  0,
				FieldSize:    1,
				FieldPadding: 1,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int16",
				FieldOffset:  2,
				FieldSize:    2,
				FieldPadding: 4,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int64",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
		},
	}
	assert.Equal(t, expected, result)
}

type largeTypeStruct struct {
	Field1 int64
	Field2 float64
	Field3 int32
	Field4 bool
}

func TestScanStructInfo_LargeTypeStruct(t *testing.T) {
	testInstance := largeTypeStruct{}
	result := ScanStructInfo(testInstance, "TestScanStructInfo_LargeTypeStruct")

	expected := entity.StructInfo{
		StructName:  "TestScanStructInfo_LargeTypeStruct",
		StructType:  "largeTypeStruct",
		StructSize:  24,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int64",
				FieldOffset:  0,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "float64",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int32",
				FieldOffset:  16,
				FieldSize:    4,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field4",
				FieldType:    "bool",
				FieldOffset:  20,
				FieldSize:    1,
				FieldPadding: 3,
			},
		},
	}
	assert.Equal(t, expected, result)
}
