package spider

import (
	"bili-monitor-system/api"
	"encoding/json"
)

func getVideoList(mid string) (UserVideoList, error) {
	videoList := UserVideoList{}
	headers := make(map[string]string)
	params := make(map[string]string)

	headers["accept"] = "application/json, text/plain, */*"
	headers["authority"] = "api.bilibili.com"
	headers["accept-encoding"] = "gzip, deflate, br"
	headers["accept-language"] = "zh-CN,zh;q=0.9"
	headers["referer"] = "https://space.bilibili.com/" + mid + "/?spm_id_from=333.999.0.0"
	headers["sec-ch-ua-mobile"] = "?0"
	headers["user-agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

	params["mid"] = mid
	params["pn"] = "1"
	params["ps"] = "25"
	params["index"] = "1"
	params["jsonp"] = "jsonp"

	body, err := api.HttpGetWithHeader(headers, "https://api.bilibili.com/x/space/arc/search", params)
	if err != nil {
		return UserVideoList{}, err
	}
	err = json.Unmarshal(body, &videoList)
	if err != nil {
		return UserVideoList{}, err
	}
	return videoList, nil
}

func getCommentList(oid string) (CommentList, error) {
	commentList := CommentList{}
	headers := make(map[string]string)
	params := make(map[string]string)

	headers["accept"] = "*/*"
	headers["authority"] = "api.bilibili.com"
	headers["accept-encoding"] = "gzip, deflate, br"
	headers["accept-language"] = "zh-CN,zh;q=0.9"
	headers["sec-ch-ua-mobile"] = "?0"
	headers["user-agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"

	params["next"] = "1"
	params["type"] = Video
	params["oid"] = oid
	params["mode"] = "3"
	params["plat"] = "1"

	body, err := api.HttpGetWithHeader(headers, "https://api.bilibili.com/x/v2/reply/main", params)
	err = json.Unmarshal(body, &commentList)
	if err != nil {
		return CommentList{}, err
	}
	return commentList, nil
}
