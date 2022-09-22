package db

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Author   string   `json:"author"`
	Mid      int      `json:"mid"`
	Title    string   `json:"title"`
	Aid      int      `json:"aid"`
	Bvid     string   `json:"bvid"`
	Comment  int      `json:"comment"`
	Play     int      `json:"play"`
	Pic      string   `json:"pic"`
	Comments Comments `json:"comments"`
}

type Dynamic struct {
	gorm.Model
	Mode int `json:"mode"`
}

type Comments []Comment

type Comment struct {
	Mid     string `json:"mid"`
	Uname   string `json:"uname"`
	Ctime   int    `json:"ctime"`
	Like    int    `json:"like"`
	Content string `json:"content"`
}

func (c Comments) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Comments) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}
