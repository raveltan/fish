package main

import (
	"math/rand"

	"github.com/raveltan/fish/dataset/controller"
	"github.com/raveltan/fish/predict/dt"
	linearreg "github.com/raveltan/fish/predict/linear-reg"
)

func removeIndex(s []*controller.MergedFishModel, index int) []*controller.MergedFishModel {
	return append(s[:index], s[index+1:]...)
}

func main() {
	data := controller.GenerateMergedDataset()
	traningSplit := int(len(data) * 7 / 10)

	trainingData := []*controller.MergedFishModel{}
	for i := 0; i < traningSplit; i++ {
		index := rand.Intn(len(data))
		trainingData = append(trainingData, data[index])
		data = removeIndex(data, index)
	}

	regression := linearreg.Regression(trainingData, data)
	data = controller.GenerateMergedDataset()
	linearreg.GenerateVis(data, regression)

	model, chartData := dt.DecisionTree()
	dt.GenerateVis(data, model, chartData)
}
