package fish_catches

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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

func GenerateVis(data []*FishCatchesCSVModel) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Total Fish Catches And Production Over Time in Indonesia",
		Subtitle: "Catches in metric tons",
	}),
		charts.WithColorsOpts(opts.Colors{opts.HSLColor(168, 80, 20)}),
		charts.WithDataZoomOpts(opts.DataZoom{}),
	)
	bar.SetXAxis(getYearRange(data)).AddSeries("Values", generateBarItems(data))

	f, _ := os.Create("dataset/fish_catches.html")

	bar.Render(f)
}

func getYearRange(data []*FishCatchesCSVModel) (years []int) {
	for _, d := range data {
		years = append(years, d.Year)
	}
	return
}

func generateBarItems(data []*FishCatchesCSVModel) (items []opts.BarData) {
	for _, d := range data {
		items = append(items, opts.BarData{
			Value: float32(d.Prod) + d.Capture,
		})
	}
	return
}
