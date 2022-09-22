package filter

import (
	"bili-monitor-system/db"
	"bili-monitor-system/internal/alarm"
	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
	"log"
)

var (
	memStore *store.MemoryStore
)

func Init() error {
	var err error
	//sf, _ := os.Open("sensitive/guns.txt")
	//defer sf.Close()
	//s := bufio.NewReader(sf)
	memStore, err = store.NewMemoryStore(store.MemoryConfig{
		//Reader:     s,
		DataSource: []string{"回声"},
	})
	if err != nil {
		return err
	}
	return nil
}

func Filter(bid string, comments db.Comments) error {
	for _, comment := range comments {
		originText := comment.Content
		filterManage := filter.NewDirtyManager(memStore)
		result, _ := filterManage.Filter().Filter(originText, '*', '@', '#')
		if len(result) != 0 {
			err := alarm.Alarm(bid, comment)
			if err != nil {
				log.Printf("[alarm] %v", err)
			}
		}
	}
	return nil
}
