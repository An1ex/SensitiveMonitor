package view

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/components"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	page := components.NewPage()

	params := r.URL.Query()
	mid := params.Get("id")
	time := params.Get("time")
	page.AddCharts(
		timeLine(mid, time),
		wordCloud(mid),
	)
	err := page.Render(w)
	if err != nil {
		return
	}
}
