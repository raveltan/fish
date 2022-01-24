package dt

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/raveltan/fish/dataset/controller"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

const (
	MAE string = "mae"
	MSE string = "mse"
)

func DecisionTree() (*trees.CARTDecisionTreeRegressor, base.FixedDataGrid) {
	// Importing Data
	data, err := base.ParseCSVToInstances("dataset/fish.csv", true)
	if err != nil {
		panic(err)
	}
	//Model
	model := trees.NewDecisionTreeRegressor("mae", 10)

	// Training Testing Split
	trainData, testData := base.InstancesTrainTestSplit(data, 0.70)
	model.Fit(trainData)

	predictions := model.Predict(testData)

	var mAE float64
	var mSE float64

	_, size := trainData.Size()
	for i := 0; i < size; i++ {
		consumption, _ := strconv.Atoi(strings.Split(trainData.RowString(i), " ")[3])
		mAE += math.Abs(float64(consumption)-predictions[i]) / float64(size)
		mSE += (float64(consumption) - predictions[i]) * (float64(consumption) - predictions[i]) / float64(size)
	}
	rMSE := math.Sqrt(mSE)
	nRMSE := rMSE / (44 - 10)
	fmt.Println("Decision Tree Regression")
	fmt.Printf("MAE = %0.2f\n", mAE)
	fmt.Printf("RMSE = %0.2f\n", rMSE)
	fmt.Printf("NRMSE = %0.2f\n\n", nRMSE)

	chartData, _ := base.InstancesTrainTestSplit(data, 0.0)
	return model, chartData
}

func GenerateVis(data []*controller.MergedFishModel, r *trees.CARTDecisionTreeRegressor, chartData base.FixedDataGrid) {
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Fish Consumption in Indonesia",
		Subtitle: "Based on Dicision Tree Regression",
	}),
		charts.WithColorsOpts(opts.Colors{opts.HSLColor(168, 80, 20)}),
		charts.WithDataZoomOpts(opts.DataZoom{}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeWesteros,
		}),
	)
	bar.SetXAxis(getYearRange(data)).AddSeries("Prediction", generatePredictedData(data, r, chartData)).AddSeries("Actual Data", generateActualData(data))

	f, _ := os.Create("predict/dt/viz.html")

	bar.Render(f)
}

func getYearRange(data []*controller.MergedFishModel) (years []int) {
	for _, d := range data {
		years = append(years, d.Year)
	}
	return
}

func generatePredictedData(data []*controller.MergedFishModel, r *trees.CARTDecisionTreeRegressor, chartData base.FixedDataGrid) (items []opts.LineData) {
	sortedPredictions := []float64{}
	for range data {
		sortedPredictions = append(sortedPredictions, 0)
	}
	predictions := r.Predict(chartData)
	_, size := chartData.Size()
	for i := 0; i < size; i++ {
		year, _ := strconv.Atoi(strings.Split(chartData.RowString(i), " ")[0])
		sortedPredictions[year-1961] = predictions[i]
	}
	for _, prediction := range sortedPredictions {
		items = append(items, opts.LineData{
			Symbol: "diamond",
			Value:  prediction,
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
