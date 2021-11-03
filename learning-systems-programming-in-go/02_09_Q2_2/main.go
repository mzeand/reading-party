package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// see https://pkg.go.dev/encoding/csv
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"Apple", "100"})
	writer.Write([]string{"Orange", "80"})
	writer.Write([]string{"Remon", "150"})
	writer.Flush()
}
