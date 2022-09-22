package view

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var (
	minute = []string{"0s", "10s", "20s", "30s", "40s", "50s"} //per minute
	hour   = []string{"0m", "10m", "20m", "30m", "40m", "50m"} //per hour
	day    = []string{"0:00", "1:00", "2:00", "3:00", "4:00", "5:00", "6:00", "7:00", "8:00", "9:00", "10:00", "11:00", "12:00",
		"13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"} //per day
	timeMap = map[string][]string{"minute": minute, "hour": hour, "day": day}
)

func generateTotalLineItems(mid string, len int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len; i++ {
		items = append(items, opts.LineData{Value: getTotalComments(mid)})
	}
	return items
}

func generateSensitiveLineItems(mid string, len int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len; i++ {
		items = append(items, opts.LineData{Value: getSensitiveComments(mid)})
	}
	return items
}

func timeLine(mid, time string) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "bilibili评论折线图",
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
	line.SetXAxis(timeMap[time]).
		AddSeries("total comments", generateTotalLineItems(mid, len(timeMap[time]))).
		AddSeries("sensitive comments", generateSensitiveLineItems(mid, len(timeMap[time]))).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
			charts.WithLabelOpts(opts.Label{Show: true}),
		)
	return line
}
