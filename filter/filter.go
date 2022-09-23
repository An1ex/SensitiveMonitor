package filter

import (
	"bufio"
	"io"
	"log"
	"os"

	"bili-monitor-system/alarm"
	"bili-monitor-system/db"

	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

var (
	memStore   *store.MemoryStore
	datasource = make([]string, 0)
)

func Init() error {
	sf, err := os.Open("sensitive/test.txt")
	if err != nil {
		return err
	}
	defer sf.Close()
	r := bufio.NewReader(sf)
	var line string
	for {
		line, err = r.ReadString('\n')
		datasource = append(datasource, line)
		if err != nil || err == io.EOF {
			break
		}
	}
	memStore, err = store.NewMemoryStore(store.MemoryConfig{
		//Reader:     s,
		DataSource: datasource,
	})
	if err != nil {
		return err
	}
	return nil
}

func Filter(bid string, comments db.Comments) int {
	count := 0
	for _, comment := range comments {
		originText := comment.Content
		filterManage := filter.NewDirtyManager(memStore)
		result, _ := filterManage.Filter().Filter(originText, '@', '，', '。', '[', ']', '！')
		if len(result) != 0 {
			count += 1
			err := alarm.Alarm(bid, comment)
			if err != nil {
				log.Printf("[alarm] %v", err)
			}
		}
	}
	return count
}
