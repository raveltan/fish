package main

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/raveltan/fish/dataset/controller"
)

func main() {
	mergedDataResult := controller.GenerateMergedDataset()
	mergedFishDatasetFile, err := os.OpenFile("dataset/fish.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer mergedFishDatasetFile.Close()

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = ','
		return gocsv.NewSafeCSVWriter(writer)
	})

	gocsv.MarshalFile(&mergedDataResult, mergedFishDatasetFile)
}
