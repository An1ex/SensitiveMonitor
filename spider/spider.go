package spider

import (
	"bili-monitor-system/filter"
	"encoding/json"
	"errors"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"log"
	"strconv"

	"bili-monitor-system/db"
)

var (
	MapSpider = make(map[string]*colly.Collector)
)

func Init() error {
	pvl := PopularVideoList{}
	s1 := newSpider("https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all", false)
	s1.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &pvl)
		if err != nil {
			return
		}
		err = switchSpider(pvl)
		if err != nil {
			return
		}
	})

	//uvl := UserVideoList{}
	//s2 := newSpider("https://api.bilibili.com/x/space/arc/search?pn=1&ps=25&index=1&mid=42543740", false)
	//s2.OnResponse(func(r *colly.Response) {
	//	err := json.Unmarshal(r.Body, &uvl)
	//	if err != nil {
	//		return
	//	}
	//	err = switchSpider(uvl)
	//	if err != nil {
	//		return
	//	}
	//})
	return nil
}

func newSpider(url string, sync bool) *colly.Collector {
	Spider := colly.NewCollector(colly.Async(sync))
	extensions.RandomUserAgent(Spider)
	MapSpider[url] = Spider
	return Spider
}

func StartSpider() {
	for url, spider := range MapSpider {
		err := spider.Visit(url)
		if err != nil {
			return
		}
	}
}

func switchSpider(i interface{}) error {
	switch i.(type) {
	case PopularVideoList:
		pvl := i.(PopularVideoList)
		for _, v := range pvl.Data.List {

			vData := db.Video{
				Author:   v.Owner.Name,
				Mid:      v.Owner.Mid,
				Title:    v.Title,
				Aid:      v.Aid,
				Bvid:     v.Bvid,
				Comment:  v.Stat.Reply,
				Play:     v.Stat.View,
				Pic:      v.Pic,
				Comments: getComments(v.Aid),
			}
			err := saveSpider(vData)
			if err != nil {
				return err
			}
		}
	case UserVideoList:
		uvl := i.(UserVideoList)
		for _, v := range uvl.Data.List.Vlist {

			vData := db.Video{
				Author:   v.Author,
				Mid:      v.Mid,
				Title:    v.Title,
				Aid:      v.Aid,
				Bvid:     v.Bvid,
				Comment:  v.Comment,
				Play:     v.Play,
				Pic:      v.Pic,
				Comments: getComments(v.Aid),
			}
			err := saveSpider(vData)
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return errors.New("UnsupportedSpiderDataType")
	}
	return nil
}

func getComments(aid int) db.Comments {
	clChan := make(chan CommentList)
	comments := db.Comments{}
	s3 := colly.NewCollector(colly.Async(true))
	extensions.RandomUserAgent(s3)

	s3.OnResponse(func(r *colly.Response) {
		var cl CommentList
		err := json.Unmarshal(r.Body, &cl)
		if err != nil {
			return
		}
		clChan <- cl
	})
	s3.Visit("https://api.bilibili.com/x/v2/reply/main?next=1&type=1&mode=3&plat=1&oid=" + strconv.Itoa(aid))
	cl := <-clChan
	for _, c := range cl.Data.Replies {
		comments = append(comments, db.Comment{
			Mid:     c.Member.Mid,
			Uname:   c.Member.Uname,
			Ctime:   c.Ctime,
			Like:    c.Like,
			Content: c.Content.Message,
		})
	}
	return comments
}

func saveSpider(vData db.Video) error {
	vOldData := db.Video{}
	err := db.DB.Where("aid = ?", vData.Aid).First(&vOldData).Error
	if err != nil {
		if err = db.DB.Create(&vData).Error; err != nil {
			log.Printf("[db] %v", err)
			return err
		}
		err = filter.Filter(vData.Bvid, vData.Comments)
		if err != nil {
			return err
		}
	} else {
		vData.Model = vOldData.Model
		if err = db.DB.Save(&vData).Error; err != nil {
			log.Printf("[db] %v", err)
			return err
		}
		if vData.Comment != vOldData.Comment {
			err = filter.Filter(vData.Bvid, vData.Comments)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
