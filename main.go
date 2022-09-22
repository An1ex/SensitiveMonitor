package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"bili-monitor-system/config"
	"bili-monitor-system/db"
	"bili-monitor-system/filter"
	"bili-monitor-system/spider"

	echarts2 "bili-monitor-system/echarts"
	"github.com/robfig/cron/v3"
)

func init() {
	if err := config.Init(); err != nil {
		log.Fatalf("[config] %v", err)
	}
	if err := db.Init(); err != nil {
		log.Fatalf("[db] %v", err)
	}
	if err := filter.Init(); err != nil {
		log.Fatalf("[filter] %v", err)
	}
	if err := spider.Init(); err != nil {
		log.Fatalf("[spider] %v", err)
	}
}

func main() {
	c := cron.New()
	if _, err := c.AddFunc("@every 2s", spider.StartSpider); err != nil {
		log.Fatalf("[colly] %v", err)
	}
	c.Start()

	http.HandleFunc("/timeline", echarts2.TimelineHandler)
	http.HandleFunc("/wordcloud", echarts2.WordCloudHandler)
	http.ListenAndServe(":8081", nil)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	c.Stop()
	fmt.Println("\rPrepare to stop...")
}
