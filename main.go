package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	type Row struct {
		Text     string `csv:"テキスト"`
		Numeric  string `csv:"数値"`
		URL      string `csv:"URL"`
		Textlrln string `csv:"セル内改行テキスト"`
	}

	fileName := "window-excel.csv"
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())
	rows := []*Row{}
	if err := gocsv.Unmarshal(reader, &rows); err != nil {
		panic(err)
	}
	for _, row := range rows {
		fmt.Printf("Text: %s\n", row.Text)
		fmt.Printf("Textlrln: %s\n", row.Textlrln)
		fmt.Printf("Numeric: %s\n", row.Numeric)
		fmt.Printf("URL: %s\n", row.URL)
	}
}
