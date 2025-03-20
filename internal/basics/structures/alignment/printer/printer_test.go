package printer

import (
	"testing"
)

func TestPrintStructInfo(t *testing.T) {
	// 	// Создаем моковые данные StructInfo
	// 	mockStructInfo := entity.StructInfo{
	// 		StructName:  "ExampleStruct",
	// 		StructType:  "struct",
	// 		StructSize:  24,
	// 		StructAlign: 8,
	// 		FieldInfos: []entity.FieldInfo{
	// 			{
	// 				FieldName:    "Field1",
	// 				FieldType:    "int8",
	// 				FieldOffset:  0,
	// 				FieldSize:    1,
	// 				FieldPadding: 0,
	// 			},
	// 			{
	// 				FieldName:    "Field2",
	// 				FieldType:    "int16",
	// 				FieldOffset:  2,
	// 				FieldSize:    2,
	// 				FieldPadding: 1,
	// 			},
	// 			{
	// 				FieldName:    "Field3",
	// 				FieldType:    "int32",
	// 				FieldOffset:  4,
	// 				FieldSize:    4,
	// 				FieldPadding: 0,
	// 			},
	// 			{
	// 				FieldName:    "Field4",
	// 				FieldType:    "int64",
	// 				FieldOffset:  8,
	// 				FieldSize:    8,
	// 				FieldPadding: 0,
	// 			},
	// 			{
	// 				FieldName:    "Field5",
	// 				FieldType:    "int16",
	// 				FieldOffset:  16,
	// 				FieldSize:    2,
	// 				FieldPadding: 6,
	// 			},
	// 		},
	// 	}

	// 	// Перехватываем вывод в буфер
	// 	var buf bytes.Buffer
	// 	oldStdout := os.Stdout
	// 	os.Stdout = &buf
	// 	defer func() { os.Stdout = oldStdout }()

	// 	// Вызываем функцию PrintStructInfo
	// 	PrintStructInfo(mockStructInfo)

	// 	// Получаем вывод
	// 	output := buf.String()

	// 	// Ожидаемый вывод
	// 	expectedOutput := `----- Structure info -----
	// Name: ExampleStruct
	// Type: struct, Size: 24, Align: 8
	// ==========	========== ====== ==== =======
	// Field Name	Field Type	Offset Size Padding
	// ----------	---------- ------ ---- -------
	// Field1	int8	0	1	0
	// Field2	int16	2	2	1
	// Field3	int32	4	4	0
	// Field4	int64	8	8	0
	// Field5	int16	16	2	6
	// ----------	----------	------	----	-------
	//      Summary	24	17	7
	// ==========	==========	======	====	=======
	// `

	// 	// Проверяем, что вывод совпадает с ожидаемым
	// 	assert.Equal(t, expectedOutput, output)
}
