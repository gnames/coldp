package coldp

import "fmt"

type FieldNumberWarning struct {
	FieldsNum, RowFieldsNum int
	Row                     []string
	Message                 string
}

func (w *FieldNumberWarning) Error() string {
	return w.Message
}

func RowToMap(headers, row []string) (map[string]string, error) {
	diff := len(headers) - len(row)
	var warning error
	if diff > 0 {
		for range diff {
			row = append(row, "")
		}
		warning = &FieldNumberWarning{
			FieldsNum:    len(headers),
			RowFieldsNum: len(row),
			Row:          row,
			Message:      fmt.Sprintf("not enough fields, filled with empty strings"),
		}
	}
	if diff < 0 {
		warning = &FieldNumberWarning{
			FieldsNum:    len(headers),
			RowFieldsNum: len(row),
			Row:          row,
			Message:      fmt.Sprintf("too many fields, extras will be ignored"),
		}
	}

	res := make(map[string]string)
	for i := range headers {
		if i < len(row) { // Prevent index out of range
			res[headers[i]] = row[i]
		}
	}

	return res, warning
}
