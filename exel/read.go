package exel

import (
	"github.com/tealeg/xlsx"
	"strings"
)

//[]string OK, []string filter err, error
func ReadColumn(file string, columnIdx int, filter func(s string) error) ([]string, []string, error) {
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		return nil, nil, err
	}

	ret, filterRet := []string{}, []string{}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for idx, cell := range row.Cells {
				if idx != columnIdx {
					continue
				}
				text, _ := cell.String()
				num := strings.TrimSpace(text)

				if err := filter(num); err != nil {
					filterRet = append(filterRet, num)
				} else {
					ret = append(ret, num)
				}
			}
		}
	}

	return ret, filterRet, nil
}
