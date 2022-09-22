package view

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateTotalWordCloud() []opts.WordCloudData {
	wordCloudMap := getTotalWordCloud()
	items := make([]opts.WordCloudData, 0)
	for _, v := range wordCloudMap {
		items = append(items, opts.WordCloudData{Name: v.Word, Value: v.Num})
	}
	return items
}

func generateUserWordCloud(mid string) []opts.WordCloudData {
	wordCloudMap := getUserWordCloud(mid)
	items := make([]opts.WordCloudData, 0)
	for _, v := range wordCloudMap {
		items = append(items, opts.WordCloudData{Name: v.Word, Value: v.Num})
	}
	return items
}

func wordCloud(mid string) *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "bilibili评论关键字词云",
			Subtitle: "总站视频/个人视频",
		}),
		charts.WithToolboxOpts(opts.Toolbox{Show: true, Feature: &opts.ToolBoxFeature{
			SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{Show: true},
			DataZoom:    &opts.ToolBoxFeatureDataZoom{Show: true},
			DataView:    &opts.ToolBoxFeatureDataView{Show: true},
			Restore:     &opts.ToolBoxFeatureRestore{Show: true},
		}}),
		charts.WithLegendOpts(opts.Legend{Show: true}),
	)
	wc.AddSeries("total wordcloud", generateTotalWordCloud()).
		AddSeries("user wordcloud", generateUserWordCloud(mid)).
		SetSeriesOptions(
			charts.WithWorldCloudChartOpts(
				opts.WordCloudChart{
					SizeRange: []float32{14, 80},
					Shape:     "circle",
				}),
		)
	return wc
}
