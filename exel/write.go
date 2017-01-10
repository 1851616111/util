package exel

import (
	"errors"
	reflectutil "github.com/1851616111/util/reflect"
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
	B int    `xlsx:"年龄"`
	C child  `xlsx:"struct"`
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

	firstLine := []string{}
	elemFieldNum := elemT.NumField()
	for i := 0; i < elemFieldNum; i++ {
		var cell string
		if cell = elemT.Field(i).Tag.Get("xlsx"); cell == "" {
			cell = elemT.Field(i).Name
		}

		firstLine = append(firstLine, cell)
	}

	v := reflect.ValueOf(obj)
	rows := [][]string{firstLine}

	for sliceIdx := 0; sliceIdx < v.Len(); sliceIdx++ {
		ele := v.Index(sliceIdx)

		cells := []string{}
		for fieldIdx := 0; fieldIdx < elemFieldNum; fieldIdx++ {
			cell := reflectutil.ValueToString(ele.Field(fieldIdx))
			cells = append(cells, cell)
		}
		rows = append(rows, cells)
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
