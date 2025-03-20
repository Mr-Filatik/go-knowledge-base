package scanner

import (
	"testing"

	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/entity"
	"github.com/stretchr/testify/assert"
)

func TestScanStructInfo(t *testing.T) {
	testInstance(t)
	testPointerInstance(t)
	testArrayInstance(t)
	testNestedStruct(t)
	testEmptyStruct(t)
	testSingleFieldStruct(t)
	testBoolStruct(t)
	testOverlapStruct(t)
}

type TestStruct struct {
	Field1 int8
	Field2 int16
	Field3 int32
	Field4 int64
	Field5 int16
}

func testInstance(t *testing.T) {
	testInstance := TestStruct{}

	result := ScanStructInfo(testInstance, "TestStruct")

	// Ожидаемый результат
	expected := entity.StructInfo{
		StructName:  "TestStruct",
		StructType:  "TestStruct",
		StructSize:  24,
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
				FieldPadding: 0,
			},
			{
				FieldName:    "Field3",
				FieldType:    "int32",
				FieldOffset:  4,
				FieldSize:    4,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field4",
				FieldType:    "int64",
				FieldOffset:  8,
				FieldSize:    8,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field5",
				FieldType:    "int16",
				FieldOffset:  16,
				FieldSize:    2,
				FieldPadding: 6,
			},
		},
	}

	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type TestPointerStruct struct {
	Field1 int32
	Field2 *int
	Field3 int8
}

func testPointerInstance(t *testing.T) {
	testInstance := TestPointerStruct{}

	result := ScanStructInfo(testInstance, "TestPointerStruct")

	// Ожидаемый результат
	expected := entity.StructInfo{
		StructName:  "TestPointerStruct",
		StructType:  "TestPointerStruct",
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
				FieldType:    "",
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

	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type ArrayStruct struct {
	Field1 [3]int16
	Field2 int32
	Field3 int8
}

func testArrayInstance(t *testing.T) {
	testInstance := ArrayStruct{}

	result := ScanStructInfo(testInstance, "TestArrayStruct")

	// Ожидаемый результат
	expected := entity.StructInfo{
		StructName:  "TestArrayStruct",
		StructType:  "ArrayStruct",
		StructSize:  16,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "",
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

	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type NestedStructInner struct {
	InnerField1 int8
	InnerField2 int32
}

type NestedStructOuter struct {
	Field1 int16
	Field2 NestedStructInner
	Field3 int8
}

func testNestedStruct(t *testing.T) {
	testInstance := NestedStructOuter{
		Field2: NestedStructInner{},
	}

	result := ScanStructInfo(testInstance, "TestNestedStruct")

	// Ожидаемый результат
	expected := entity.StructInfo{
		StructName:  "TestNestedStruct",
		StructType:  "NestedStructOuter",
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
				FieldType:    "NestedStructInner",
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

	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type emptyStruct struct{}

func testEmptyStruct(t *testing.T) {
	testInstance := emptyStruct{}
	result := ScanStructInfo(testInstance, "TestEmptyStruct")
	expected := entity.StructInfo{
		StructName:  "TestEmptyStruct",
		StructType:  "emptyStruct",
		StructSize:  0,
		StructAlign: 8,
		FieldInfos:  []entity.FieldInfo{},
	}
	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type singleFieldStruct struct {
	Field1 int64
}

func testSingleFieldStruct(t *testing.T) {
	testInstance := singleFieldStruct{}
	result := ScanStructInfo(testInstance, "TestSingleFieldStruct")
	expected := entity.StructInfo{
		StructName:  "TestSingleFieldStruct",
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
	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type boolStruct struct {
	Field1 bool
	Field2 int16
	Field3 bool
}

func testBoolStruct(t *testing.T) {
	testInstance := boolStruct{}
	result := ScanStructInfo(testInstance, "TestBoolStruct")
	expected := entity.StructInfo{
		StructName:  "TestBoolStruct",
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
	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type overlapStruct struct {
	Field1 int32
	Field2 int8
	Field3 int64
}

func testOverlapStruct(t *testing.T) {
	testInstance := overlapStruct{}
	result := ScanStructInfo(testInstance, "TestOverlapStruct")
	expected := entity.StructInfo{
		StructName:  "TestOverlapStruct",
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
	// Проверяем, что результат совпадает с ожидаемым
	assert.Equal(t, expected, result)
}

type aligned16Struct struct {
	Field1 int64
	Field2 int64
	Field3 int8
}

type aligned8Struct struct {
	Field1 int32
	Field2 int32
	Field3 int8
}

type unorderedStruct struct {
	Field1 int8
	Field2 int64
	Field3 int16
}

type largeTypeStruct struct {
	Field1 int64
	Field2 float64
	Field3 int32
}

type stringStruct struct {
	Field1 int64
	Field2 string
	Field3 int32
}
