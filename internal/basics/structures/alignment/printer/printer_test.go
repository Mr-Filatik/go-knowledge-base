package printer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/entity"
	"github.com/stretchr/testify/assert"
)

func TestPrintStructInfo(t *testing.T) {
	testStructInfo := entity.StructInfo{
		StructName:  "TestStruct",
		StructType:  "testStruct",
		StructSize:  24,
		StructAlign: 8,
		FieldInfos: []entity.FieldInfo{
			{
				FieldName:    "Field1",
				FieldType:    "int8",
				FieldOffset:  0,
				FieldSize:    1,
				FieldPadding: 0,
			},
			{
				FieldName:    "Field2",
				FieldType:    "int16",
				FieldOffset:  2,
				FieldSize:    2,
				FieldPadding: 1,
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

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintStructInfo(testStructInfo)

	w.Close()
	os.Stdout = oldStdout

	var capturedOutput bytes.Buffer
	io.Copy(&capturedOutput, r)

	expectedOutput := `----- Structure info -----
Name: TestStruct
Type: testStruct, Size: 24, Align: 8
========== ========== ====== ==== =======
Field Name Field Type Offset Size Padding
---------- ---------- ------ ---- -------
Field1     int8       0      1    0
Field2     int16      2      2    1
Field3     int32      4      4    0
Field4     int64      8      8    0
Field5     int16      16     2    6
---------- ---------- ------ ---- -------
           Summary    24     17   7
========== ========== ====== ==== =======
`
	assert.Equal(t, expectedOutput, capturedOutput.String())
}
