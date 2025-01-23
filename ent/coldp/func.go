package coldp

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gnames/coldp/config"
	"github.com/gnames/gnfmt/gncsv"
	csvConfig "github.com/gnames/gnfmt/gncsv/config"
	"github.com/gnames/gnlib"
)

func ReadJSON[T DataLoader](
	path string,
	chIn chan T) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	slog.Info("Parsing references JSON-CSL file")
	var cslRec CSLRecords
	err = json.Unmarshal(byteValue, &cslRec)
	if err != nil {
		return err
	}
	for _, csl := range cslRec {
		chIn <- csl.ToReference().(T)
	}
	close(chIn)
	return nil
}

func ReadJSONL[T DataLoader](
	path string,
	chIn chan T) error {
	f, err := os.Open(path) // Replace "my_file.txt" with your file name
	if err != nil {
		return err
	}
	defer f.Close() // Ensure the file is closed when done

	slog.Info("Reading references from JSON-CSL file")

	// Create a new scanner
	scanner := bufio.NewScanner(f)

	// Loop over all lines in the file
	for scanner.Scan() {
		bs := scanner.Bytes()
		var csl CSL
		err = json.Unmarshal(bs, &csl)
		if err != nil {
			return err
		}
		chIn <- csl.ToReference().(T)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	close(chIn)
	return nil
}

func Read[T DataLoader](
	cfg config.Config,
	path string,
	chOut chan T) error {
	chIn := make(chan []string)
	var wg sync.WaitGroup
	wg.Add(1)

	opts := []csvConfig.Option{
		csvConfig.OptPath(path),
		csvConfig.OptBadRowMode(cfg.BadRow),
		csvConfig.OptWithQuotes(cfg.WithQuotes),
	}
	csvCfg, err := csvConfig.New(opts...)
	headers := gnlib.Map(csvCfg.Headers, func(s string) string {
		s = strings.ToLower(s)
		els := strings.Split(s, ":")
		return els[len(els)-1]
	})

	go func() {
		defer wg.Done()
		for row := range chIn {
			var t T
			dl, warn := t.Load(headers, row)
			t = dl.(T)
			if warn != nil {
				slog.Warn("Cannot read row", "warn", warn, "row", row)
				continue
			}
			chOut <- t
		}
	}()

	if err != nil {
		return err
	}

	csv := gncsv.New(csvCfg)
	_, err = csv.Read(context.Background(), chIn)
	if err != nil {
		return err
	}
	close(chIn)
	wg.Wait()

	return nil
}

type FieldNumberWarning struct {
	FieldsNum, RowFieldsNum int
	Row                     []string
	Message                 string
}

func (w *FieldNumberWarning) Error() string {
	return w.Message
}

func ToInt(s string) sql.NullInt64 {
	var res sql.NullInt64
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	v, err := strconv.Atoi(s)
	if err != nil {
		return res
	}
	res.Int64 = int64(v)
	res.Valid = true
	return res
}

func ToBool(s string) sql.NullBool {
	var res sql.NullBool
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	if s == "" {
		return res
	}
	res.Valid = true
	b := s[0]
	switch b {
	case '1', 'y', 't':
		res.Bool = true
	default:
		res.Bool = false
	}
	return res
}

func ToFloat(s string) sql.NullFloat64 {
	var res sql.NullFloat64
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return res
	}
	res.Float64 = v
	res.Valid = true
	return res
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
			fld := strings.ToLower(headers[i])
			fld = strings.ReplaceAll(fld, "_", "")
			res[headers[i]] = row[i]
		}
	}

	return res, warning
}
