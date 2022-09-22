package spider

const (
	Video = "1"
)

type UserVideoList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List struct {
			Vlist []struct {
				Comment     int    `json:"comment"`
				Play        int    `json:"play"`
				Pic         string `json:"pic"`
				Description string `json:"description"`
				Title       string `json:"title"`
				Author      string `json:"author"`
				Mid         int    `json:"mid"`
				Created     int    `json:"created"`
				Length      string `json:"length"`
				Aid         int    `json:"aid"`
				Bvid        string `json:"bvid"`
			} `json:"vlist"`
		} `json:"list"`
		Page struct {
			Pn    int `json:"pn"`
			Ps    int `json:"ps"`
			Count int `json:"count"`
		} `json:"page"`
	} `json:"data"`
}

type PopularVideoList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			Aid     int    `json:"aid"`
			Videos  int    `json:"videos"`
			Tname   string `json:"tname"`
			Pic     string `json:"pic"`
			Title   string `json:"title"`
			Pubdate int    `json:"pubdate"`
			Ctime   int    `json:"ctime"`
			Desc    string `json:"desc"`
			Owner   struct {
				Mid  int    `json:"mid"`
				Name string `json:"name"`
				Face string `json:"face"`
			} `json:"owner"`
			Stat struct {
				Aid      int `json:"aid"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Favorite int `json:"favorite"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Dislike  int `json:"dislike"`
			} `json:"stat"`
			Dynamic   string `json:"dynamic"`
			Cid       int    `json:"cid"`
			ShortLink string `json:"short_link"`
			Bvid      string `json:"bvid"`
		} `json:"list"`
	} `json:"data"`
}

type CommentList struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Cursor struct {
			IsBegin     bool   `json:"is_begin"`
			Prev        int    `json:"prev"`
			Next        int    `json:"next"`
			IsEnd       bool   `json:"is_end"`
			AllCount    int    `json:"all_count"`
			Mode        int    `json:"mode"`
			SupportMode []int  `json:"support_mode"`
			Name        string `json:"name"`
		} `json:"cursor"`
		Replies []struct {
			Oid    int `json:"oid"`
			Type   int `json:"type"`
			Ctime  int `json:"ctime"`
			Like   int `json:"like"`
			Member struct {
				Mid   string `json:"mid"`
				Uname string `json:"uname"`
			} `json:"member"`
			Content struct {
				Message string `json:"message"`
			} `json:"content"`
			SecReplies []struct {
				Oid    int `json:"oid"`
				Ctime  int `json:"ctime"`
				Like   int `json:"like"`
				Member struct {
					Mid   string `json:"mid"`
					Uname string `json:"uname"`
				} `json:"member"`
				Content struct {
					Message string `json:"message"`
				} `json:"content"`
				ReplyControl struct {
					TimeDesc string `json:"time_desc"`
				} `json:"reply_control"`
			} `json:"replies"`
			UpAction struct { // up主的操作
				Like  bool `json:"like"`
				Reply bool `json:"reply"`
			} `json:"up_action"`
			ReplyControl struct {
				UpLike            bool   `json:"up_like"`
				SubReplyEntryText string `json:"sub_reply_entry_text"`
				SubReplyTitleText string `json:"sub_reply_title_text"`
				TimeDesc          string `json:"time_desc"`
			} `json:"reply_control"`
		} `json:"replies"`
	} `json:"data"`
}
