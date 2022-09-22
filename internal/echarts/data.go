package echarts

import (
	"bili-monitor-system/db"
	"github.com/huichen/sego"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"unicode"
)

type WordFrequency struct {
	mu   sync.Mutex
	data map[string]interface{}
}

type SortedWordFrequency struct {
	Word string
	Num  int
}

var (
	TotalWordFrequency = WordFrequency{
		mu:   sync.Mutex{},
		data: make(map[string]interface{}),
	}
	UserWordFrequency = WordFrequency{
		mu:   sync.Mutex{},
		data: make(map[string]interface{}),
	}
	segmenter sego.Segmenter
)

func getTotalComments() int {
	return rand.Intn(50)
}

func getSensitiveComments() int {
	return rand.Intn(500)
}

func getTotalWordCloud() []SortedWordFrequency {
	//segmenter.LoadDictionary("sensitive/dictionary.txt")
	videos := make([]db.Video, 0)
	wg := sync.WaitGroup{}

	db.DB.Select("comments").Find(&videos)
	for _, video := range videos {
		for _, comment := range video.Comments {
			wg.Add(1)
			go TotalWordFrequency.calFrequency(comment.Content, &wg)
		}
	}
	wg.Wait()
	lstWordFrequency := TotalWordFrequency.sortFrequency()
	return lstWordFrequency
}

func getUserWordCloud() []SortedWordFrequency {
	videos := make([]db.Video, 0)
	wg := sync.WaitGroup{}

	db.DB.Where("mid = ?", "401742377").Select("comments").Find(&videos)
	for _, video := range videos {
		for _, comment := range video.Comments {
			wg.Add(1)
			go UserWordFrequency.calFrequency(comment.Content, &wg)
		}
	}
	wg.Wait()
	lstWordFrequency := UserWordFrequency.sortFrequency()
	return lstWordFrequency
}

func (wf *WordFrequency) calFrequency(originText string, wg *sync.WaitGroup) {
	//去掉分隔符
	f := func(c rune) bool {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && !unicode.IsPunct(c) && !unicode.IsSymbol(c) {
			return true
		}
		return false
	}
	words := strings.FieldsFunc(originText, f)
	//segments := segmenter.Segment([]byte(originText))
	//words := sego.SegmentsToSlice(segments, true)
	//如果字典里有该单词则加1，否则添加入字典赋值为1
	wf.mu.Lock()
	for _, v := range words {
		if _, ok := wf.data[v]; ok {
			wf.data[v] = wf.data[v].(int) + 1
		} else {
			wf.data[v] = 1
		}
	}
	wf.mu.Unlock()
	wg.Done()
	return
}

func (wf *WordFrequency) sortFrequency() []SortedWordFrequency {
	//按照单词出现的频率排序
	lstWordFrequencyNum := make([]SortedWordFrequency, 0)
	for k, v := range wf.data {
		lstWordFrequencyNum = append(lstWordFrequencyNum, SortedWordFrequency{k, v.(int)})
	}
	sort.Slice(lstWordFrequencyNum, func(i, j int) bool {
		return lstWordFrequencyNum[i].Num > lstWordFrequencyNum[j].Num
	})
	if len(lstWordFrequencyNum) > 100 {
		return lstWordFrequencyNum[:100]
	}
	return lstWordFrequencyNum
}
