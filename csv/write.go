package csv

import (
    "encoding/csv"
    "log"
    "os"
)

func Write() {

	// initialize the slice
    records := [][]string{
        {"first_name", "last_name", "occupation"},
        {"John", "Doe", "gardener"},
        {"Lucy", "Smith", "teacher"},
    }

	// append single element to the slice
	record := []string{}
	record = append(record, "Brian", "Bethamy", "programmer")
	records = append(records, record)

	// create the file
    file, err := os.Create("users.csv")
    defer file.Close()
    if err != nil {
        log.Fatalln("failed to open file", err)
    }

	// write slice into file
    writer := csv.NewWriter(file)
    err = writer.WriteAll(records)
    if err != nil {
        log.Fatal(err)
    }
}