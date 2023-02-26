package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/ahmetb/go-linq/v3"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Row struct {
	Date        string
	Amount      int
	Description string
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		return
	}

	err := processFiles(args)
	if err != nil {
		log.Fatal(err)
	}
}

func processFiles(fileNames []string) error {
	rows := []Row{}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		reader := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
		reader.LazyQuotes = true

		_, err = reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		for {
			columns, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}

				return err
			}

			if len(columns) < 4 {
				return fmt.Errorf("too few values: fileName=%s, columns=%v", fileName, columns)
			}

			amount, err := strconv.Atoi(columns[3])
			if err != nil {
				return err
			}

			rows = append(rows, Row{
				Date:        columns[1],
				Amount:      -1 * amount,
				Description: columns[2],
			})
		}
	}

	linq.From(rows).OrderBy(func(i interface{}) interface{} {
		return i.(Row).Date
	}).ToSlice(&rows)

	fmt.Println(`"Date","Amount","Description"`)

	for _, row := range rows {
		fmt.Printf(`"%s",%d,"%s"`, row.Date, row.Amount, row.Description)
		fmt.Println()
	}

	return nil
}
