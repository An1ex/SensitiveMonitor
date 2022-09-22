package spider

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	"bili-monitor-system/db"
	"bili-monitor-system/filter"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

var (
	MapSpider        = make(map[string]*colly.Collector)
	domain2collector = map[string]*colly.Collector{}
)

func Init() error {
	domain2collector["api.bilibili.com"] = initPopularVideoCollector()
	//domain2collector["api.bilibili.com"] = initUserVideoCollector()
	return nil
}

func Start() {
	urls := []string{"https://api.bilibili.com/x/web-interface/ranking/v2?rid=0&type=all"}
	//user_urls := []string{"https://api.bilibili.com/x/space/arc/search?pn=1&ps=25&index=1&mid=42543740"}
	for _, u := range urls {
		instance := factory(u)
		instance.Visit(u)
	}
}

func factory(s string) *colly.Collector {
	u, _ := url.Parse(s)
	return domain2collector[u.Host]
	//return domain2collector[s]
}

func initPopularVideoCollector() *colly.Collector {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	pvl := PopularVideoList{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &pvl)
		if err != nil {
			return
		}
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
			err = saveSpider(vData)
			if err != nil {
				return
			}
		}
	})
	return c
}

func initUserVideoCollector() *colly.Collector {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	uvl := UserVideoList{}
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &uvl)
		if err != nil {
			return
		}
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
			err = saveSpider(vData)
			if err != nil {
				return
			}
		}
	})
	return c
}

func getComments(aid int) db.Comments {
	clChan := make(chan CommentList)
	comments := db.Comments{}
	c := colly.NewCollector(colly.Async(true))
	extensions.RandomUserAgent(c)

	c.OnResponse(func(r *colly.Response) {
		var cl CommentList
		err := json.Unmarshal(r.Body, &cl)
		if err != nil {
			return
		}
		clChan <- cl
	})
	c.Visit("https://api.bilibili.com/x/v2/reply/main?next=1&type=1&mode=3&plat=1&oid=" + strconv.Itoa(aid))
	cl := <-clChan
	for _, comment := range cl.Data.Replies {
		comments = append(comments, db.Comment{
			Mid:     comment.Member.Mid,
			Uname:   comment.Member.Uname,
			Ctime:   comment.Ctime,
			Like:    comment.Like,
			Content: comment.Content.Message,
		})
	}
	return comments
}

func saveSpider(vData db.Video) error {
	vOldData := db.Video{}
	err := db.DB.Where("aid = ?", vData.Aid).First(&vOldData).Error
	if err != nil { //视频已未爬取过
		vData.Sensitive = filter.Filter(vData.Bvid, vData.Comments)
		if err = db.DB.Create(&vData).Error; err != nil {
			log.Printf("[db] %v", err)
			return err
		}
	} else { //视频已爬取过
		vData.Model = vOldData.Model
		//if vData.Comment != vOldData.Comment { //有新评论，才需要过滤
		//	vData.Sensitive = filter.Filter(vData.Bvid, vData.Comments)
		//}
		vData.Sensitive = filter.Filter(vData.Bvid, vData.Comments)
		if err = db.DB.Save(&vData).Error; err != nil {
			log.Printf("[db] %v", err)
			return err
		}
	}
	return nil
}
