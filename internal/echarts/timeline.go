package echarts

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func generateTotalLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 24; i++ {
		items = append(items, opts.LineData{Value: getTotalComments()})
	}
	return items
}

func generateSensitiveLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 24; i++ {
		items = append(items, opts.LineData{Value: getSensitiveComments()})
	}
	return items
}

func TimelineHandler(w http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Bilibili舆情评论当日监控图",
			Subtitle: "总评论/敏感评论",
		}),
		charts.WithToolboxOpts(opts.Toolbox{Show: true, Feature: &opts.ToolBoxFeature{
			SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{Show: true},
			DataZoom:    &opts.ToolBoxFeatureDataZoom{Show: true},
			DataView:    &opts.ToolBoxFeatureDataView{Show: true},
			Restore:     &opts.ToolBoxFeatureRestore{Show: true},
		}}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
	)
	line.SetXAxis([]string{"0:00", "1:00", "2:00", "3:00", "4:00", "5:00", "6:00", "7:00", "8:00", "9:00", "10:00", "11:00", "12:00",
		"13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"}).
		AddSeries("total comments", generateTotalLineItems()).
		AddSeries("sensitive comments", generateSensitiveLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
			charts.WithLabelOpts(opts.Label{Show: true}),
		)

	err := line.Render(w)
	if err != nil {
		return
	}

	// generate html
	//f, _ := os.Create("template/line.html")
	//line.Render(f)
}
