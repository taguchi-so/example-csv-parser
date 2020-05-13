package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gocarina/gocsv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	example1()
	example2()
}

func example1() {
	fileName := "window-excel.csv"
	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("encoding/csv\n")
	reader := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())
	r := csv.NewReader(reader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("Text: %s\n", record[0])
		fmt.Printf("Textlrln: %s\n", record[1])
		fmt.Printf("Numeric: %s\n", record[2])
		fmt.Printf("URL: %s\n", record[3])
	}
	fmt.Printf("\n")
}

func example2() {
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
	fmt.Printf("github.com/gocarina/gocsv\n")
	for _, row := range rows {
		fmt.Printf("Text: %s\n", row.Text)
		fmt.Printf("Textlrln: %s\n", row.Textlrln)
		fmt.Printf("Numeric: %s\n", row.Numeric)
		fmt.Printf("URL: %s\n", row.URL)
	}
	fmt.Printf("\n")
}
