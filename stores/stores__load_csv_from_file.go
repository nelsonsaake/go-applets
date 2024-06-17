package main

import (
	"encoding/csv"
	"os"
	"projects/saelections/pkg/sysout"
	"projects/saelections/pkg/ufs"
)

func loadItemsFromCsvFile(filename string) (items []Item) {
	if exists, err := ufs.Exists(filename); !exists {
		sysout.Print("file doesn't exist, %v", err)
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6

	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range record {
		// name, price, units, subunits, unitprice, subunitprice
		item := Item{
			Name:         row[0],
			Price:        row[1],
			Units:        row[2],
			SubUnits:     row[3],
			UnitPrice:    row[4],
			SubUnitPrice: row[5],
		}
		items = append(items, item)
	}

	return
}
