package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/entity"
)

func PrintStructInfo(si entity.StructInfo) {
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
