package scanner

import (
	"reflect"
	"strconv"
	"unsafe"

	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/entity"
)

func ScanStructInfo(s interface{}, name string) entity.StructInfo {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)

	si := entity.StructInfo{
		StructName:  name,
		StructType:  typ.Name(),
		StructSize:  int32(typ.Size()),
		StructAlign: int32(unsafe.Alignof(typ)),
		FieldInfos:  make([]entity.FieldInfo, 0),
	}

	count := typ.NumField()
	if count != 0 {
		for i := 0; i < count-1; i++ {
			field := typ.Field(i)
			nextField := typ.Field(i + 1)

			fieldValue := val.Field(i)
			fieldSize := int32(fieldValue.Type().Size())
			fieldOffset := int32(field.Offset)
			fieldPadding := int32(nextField.Offset) - fieldOffset - fieldSize

			fi := entity.FieldInfo{
				FieldName:    field.Name,
				FieldType:    getTypeName(field.Type),
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

		fi := entity.FieldInfo{
			FieldName:    field.Name,
			FieldType:    getTypeName(field.Type),
			FieldSize:    fieldSize,
			FieldOffset:  fieldOffset,
			FieldPadding: fieldPadding,
		}

		si.FieldInfos = append(si.FieldInfos, fi)
	}
	return si
}

func getTypeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Array:
		return "[" + strconv.Itoa(t.Len()) + "]" + getTypeName(t.Elem())
	case reflect.Slice:
		return "[]" + getTypeName(t.Elem())
	case reflect.Pointer:
		return "*" + getTypeName(t.Elem())
	case reflect.Map:
		return "map[" + getTypeName(t.Key()) + "]" + getTypeName(t.Elem())
	default:
		return t.Name()
	}
}
