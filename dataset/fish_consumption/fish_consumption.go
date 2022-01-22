package fish_consumption

import (
	"os"

	"github.com/gocarina/gocsv"
)

type FishConsumptionCSVModel struct {
	Country     string  `csv:"country"`
	CountryCode string  `csv:"code"`
	Year        int     `csv:"year"`
	Consumption float32 `csv:"consumption"`
}

func GetFishConsumptionData(countryCode string, startYear int, endYear int) []*FishConsumptionCSVModel {
	datasetFile, err := os.OpenFile("dataset/fish_consumption/fish_consumption.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer datasetFile.Close()

	datas := []*FishConsumptionCSVModel{}
	results := []*FishConsumptionCSVModel{}

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
