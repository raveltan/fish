package controller

import (
	"github.com/raveltan/fish/dataset/fish_catches"
	"github.com/raveltan/fish/dataset/fish_consumption"
)

type MergedFishModel struct {
	Year        int     `csv:"year" json:"year"`
	Production  int     `csv:"production" json:"production"`
	Capture     float32 `csv:"capture" json:"capture"`
	Consumption float32 `csv:"consumption" json:"consumption"`
}

func GenerateMergedDataset() []*MergedFishModel {
	consumption := fish_consumption.GetFishConsumptionData("IDN", 1961, 2017)
	// Generate Visualisation
	fish_consumption.GenerateVis(consumption)

	catches := fish_catches.GetFishCatchesData("IDN", 1961, 2017)
	// Generate Visualisation
	fish_catches.GenerateVis(catches)

	mergedFishModelData := []*MergedFishModel{}
	for index := 0; index < len(consumption); index++ {
		catchData := catches[index]
		consumptionData := consumption[index]
		mergedFishModelData = append(mergedFishModelData, &MergedFishModel{
			Year:        catchData.Year,
			Production:  catchData.Prod,
			Capture:     catchData.Capture,
			Consumption: consumptionData.Consumption,
		})
	}

	// DEBUG Print dataset data
	// for _, data := range mergedFishModelData {
	// 	d, _ := json.Marshal(*data)
	// 	fmt.Println(string(d))
	// }

	return mergedFishModelData

}
