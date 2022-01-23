package fish_consumption

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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

func GenerateVis(data []*FishConsumptionCSVModel) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Total Fish Consumption Time in Indonesia",
		Subtitle: "Food supply quantity (kg/capita/yr)",
	}),
		charts.WithColorsOpts(opts.Colors{opts.HSLColor(168, 80, 20)}),
		charts.WithDataZoomOpts(opts.DataZoom{}),
	)
	bar.SetXAxis(getYearRange(data)).AddSeries("Values", generateBarItems(data))

	f, _ := os.Create("dataset/fish_consumption.html")

	bar.Render(f)
}

func getYearRange(data []*FishConsumptionCSVModel) (years []int) {
	for _, d := range data {
		years = append(years, d.Year)
	}
	return
}

func generateBarItems(data []*FishConsumptionCSVModel) (items []opts.BarData) {
	for _, d := range data {
		items = append(items, opts.BarData{
			Value: d.Consumption,
		})
	}
	return
}
