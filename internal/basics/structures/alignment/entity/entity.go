package entity

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
