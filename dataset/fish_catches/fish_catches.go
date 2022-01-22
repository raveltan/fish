package fish_catches

import (
	"os"

	"github.com/gocarina/gocsv"
)

type FishCatchesCSVModel struct {
	Country     string  `csv:"country"`
	CountryCode string  `csv:"code"`
	Year        int     `csv:"year"`
	Prod        int     `csv:"prod"`
	Capture     float32 `csv:"capture"`
}

func GetFishCatchesData(countryCode string, startYear int, endYear int) []*FishCatchesCSVModel {
	datasetFile, err := os.OpenFile("dataset/fish_catches/fish_catches.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}

	defer datasetFile.Close()

	datas := []*FishCatchesCSVModel{}
	results := []*FishCatchesCSVModel{}

	if err := gocsv.UnmarshalFile(datasetFile, &datas); err != nil {
		panic(err)
	}

	for _, data := range datas {
		if data.CountryCode == countryCode {
			if data.Year >= startYear && data.Year <= endYear {
				results = append(results, data)
			}
		}
	}

	return results
}
