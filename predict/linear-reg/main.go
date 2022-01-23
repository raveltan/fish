package linearreg

import (
	"fmt"
	"math"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/raveltan/fish/dataset/controller"
	"github.com/sajari/regression"
)

func Regression(data []*controller.MergedFishModel, testData []*controller.MergedFishModel) regression.Regression {
	var r regression.Regression
	r.SetVar(0, "Production")
	r.SetVar(1, "Year")
	r.SetVar(2, "Capture")
	r.SetObserved("Consumption")

	for _, record := range data {
		// d, _ := json.Marshal(record)
		// fmt.Println(string(d))
		r.Train(regression.DataPoint(float64(record.Consumption), []float64{float64(record.Production), float64(record.Year), float64(record.Capture)}))
	}
	r.Run()
	fmt.Println("Multivariable Linear Regression")
	fmt.Printf("\nRegression Formula:\n%v\n", r.Formula)

	var mAE float64

	for _, record := range testData {
		prediction, _ := r.Predict([]float64{float64(record.Production), float64(record.Year), float64(record.Capture)})
		// fmt.Println("Prediction")
		// TODO: ASK MISS NURUL
		// fmt.Println(record.Year)
		// fmt.Println(prediction)
		// fmt.Println(prediction - float64(record.Consumption))
		// fmt.Println("--------")
		mAE += math.Abs(float64(record.Consumption)-prediction) / float64(len(testData))
	}
	fmt.Printf("MAE = %0.2f\n\n", mAE)
	return r
}

func GenerateVis(data []*controller.MergedFishModel, r regression.Regression) {
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Prediction of Total Fish Consumption in Indonesia",
		Subtitle: "Based on Multivariable Linear Regression",
	}),
		charts.WithColorsOpts(opts.Colors{opts.HSLColor(168, 80, 20)}),
		charts.WithDataZoomOpts(opts.DataZoom{}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWesteros,
		}),
	)
	bar.SetXAxis(getYearRange(data)).AddSeries("Prediction", generatePredictedData(data, r)).AddSeries("Actual Data", generateActualData(data))

	f, _ := os.Create("predict/linear-reg/viz.html")

	bar.Render(f)
}

func getYearRange(data []*controller.MergedFishModel) (years []int) {
	for _, d := range data {
		years = append(years, d.Year)
	}
	return
}

func generatePredictedData(data []*controller.MergedFishModel, r regression.Regression) (items []opts.LineData) {
	for _, record := range data {
		res, _ := r.Predict([]float64{float64(record.Production), float64(record.Year), float64(record.Capture)})
		items = append(items, opts.LineData{
			Symbol: "diamond",
			Value:  res,
		})
	}
	return
}
func generateActualData(data []*controller.MergedFishModel) (items []opts.LineData) {
	for _, record := range data {
		items = append(items, opts.LineData{
			Value:  record.Consumption,
			Symbol: "rect",
		})
	}
	return
}
