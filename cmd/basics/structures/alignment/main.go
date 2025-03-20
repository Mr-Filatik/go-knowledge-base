package main

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
	"unsafe"
)

type Bar struct {
	aaa  int8
	aaa2 int8
	ccc  int16
	ddd  int32
	bbb  int64
	ggg  int16
}

func main() {
	bar := Bar{
		aaa:  1,
		bbb:  2,
		ccc:  3,
		ddd:  4,
		aaa2: 5,
		ggg:  6,
	}

	si := CreateStructInfo(bar, "example one")
	PrintStructInfo(si)
}

func CreateStructInfo(s interface{}, name string) StructInfo {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)

	si := StructInfo{
		StructName:  name,
		StructType:  typ.Name(),
		StructSize:  int32(typ.Size()),
		StructAlign: int32(unsafe.Alignof(typ)),
		FieldInfos:  make([]FieldInfo, 0),
	}

	count := typ.NumField()
	for i := 0; i < count-1; i++ {
		field := typ.Field(i)
		nextField := typ.Field(i + 1)

		fieldValue := val.Field(i)

		fieldSize := int32(fieldValue.Type().Size())
		fieldOffset := int32(field.Offset)
		fieldPadding := int32(nextField.Offset) - fieldOffset - fieldSize

		fi := FieldInfo{
			FieldName:    field.Name,
			FieldType:    field.Type.Name(),
			FieldSize:    fieldSize,
			FieldOffset:  fieldOffset,
			FieldPadding: fieldPadding,
		}

		si.FieldInfos = append(si.FieldInfos, fi)
	}

	field := typ.Field(count - 1)
	fieldValue := val.Field(count - 1)
	fieldSize := int32(fieldValue.Type().Size())
	fieldOffset := int32(field.Offset)
	fieldPadding := si.StructSize - fieldOffset - fieldSize

	fi := FieldInfo{
		FieldName:    field.Name,
		FieldType:    field.Type.Name(),
		FieldSize:    fieldSize,
		FieldOffset:  fieldOffset,
		FieldPadding: fieldPadding,
	}

	si.FieldInfos = append(si.FieldInfos, fi)

	return si
}

func PrintStructInfo(si StructInfo) {
	fmt.Printf("----- Structure info -----\n")
	fmt.Printf("Name: %v\n", si.StructName)
	fmt.Printf("Type: %v, Size: %d, Align: %d\n", si.StructType, si.StructSize, si.StructAlign)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Fprintln(w, "==========\t==========\t======\t====\t=======")
	fmt.Fprintln(w, "Field Name\tField Type\tOffset\tSize\tPadding")
	fmt.Fprintln(w, "----------\t----------\t------\t----\t-------")

	fieldsSize := 0
	fieldsPadding := 0
	for _, field := range si.FieldInfos {
		fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\n",
			field.FieldName,
			field.FieldType,
			field.FieldOffset,
			field.FieldSize,
			field.FieldPadding,
		)
		fieldsSize += int(field.FieldSize)
		fieldsPadding += int(field.FieldPadding)
	}
	fmt.Fprintln(w, "----------\t----------\t------\t----\t-------")
	fmt.Fprintf(w, "%s\t%s\t%d\t%d\t%d\n",
		" ",
		"Summary",
		si.StructSize,
		fieldsSize,
		fieldsPadding,
	)
	fmt.Fprintln(w, "==========\t==========\t======\t====\t=======")
	w.Flush()
}

type StructInfo struct {
	StructName  string
	StructType  string
	StructSize  int32
	StructAlign int32
	FieldInfos  []FieldInfo
}

type FieldInfo struct {
	FieldName    string
	FieldType    string
	FieldOffset  int32
	FieldSize    int32
	FieldPadding int32
}
