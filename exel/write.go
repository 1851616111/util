package exel

import (
	"errors"
	"github.com/tealeg/xlsx"
	"reflect"
	"strings"
)

var DefaultTTargetFile string = "exel.xlsx"
var DefaultSheet string = "sheet"

type child struct {
	C string
}
type test struct {
	A string `xlsx:"姓名"`
	B int    `xlsx:"-"`
	C child  `xlsx:"xxx"`
}

func MarshalToFile(obj interface{}) error {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Slice {
		return errors.New("unsupported object type, only struct")
	}

	elemT := t.Elem()
	if elemT.Kind() == reflect.Ptr {
		elemT = elemT.Elem()
	}

	if elemT.Kind() != reflect.Struct {
		return errors.New("unsupported object type, only struct")
	}

	v := reflect.ValueOf(obj)
	rows := [][]string{columnNames(elemT)}

	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)

		rows = append(rows, ValueToSlice(item)[1:])
	}

	return Write(rows)
}

func Write(rows [][]string) error {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var err error

	if !strings.Contains(DefaultTTargetFile, ".xlsx") {
		DefaultTTargetFile = DefaultTTargetFile + ".xlsx"
	}
	file = xlsx.NewFile()
	sheet, err = file.AddSheet(DefaultSheet)
	if err != nil {
		return err
	}

	writeSheet(sheet, rows)

	if err = file.Save(DefaultTTargetFile); err != nil {
		return err
	}

	return nil
}

func writeSheet(sheet *xlsx.Sheet, rows [][]string) {
	for _, row := range rows {
		newRow := sheet.AddRow()
		for _, cell := range row {
			newCell := newRow.AddCell()
			newCell.Value = cell
		}
	}
}
