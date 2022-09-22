package alarm

import (
	"fmt"
	"time"

	"bili-monitor-system/config"
	"bili-monitor-system/db"
)

func Alarm(bvid string, comment db.Comment) error {
	defaultMailer := NewEmail(&SMTPInfo{
		Host:     config.MailConf.Host,
		Port:     config.MailConf.Port,
		IsSSL:    config.MailConf.IsSSL,
		UserName: config.MailConf.UserName,
		Password: config.MailConf.Password,
		From:     config.MailConf.From,
	})

	err := defaultMailer.SendMail(
		config.MailConf.To,
		fmt.Sprintf("[警告]评论区检测到敏感信息"),
		fmt.Sprintf("原视频：%s<br />评论内容：%s<br />用户：%s<br />UID：%s<br />发布时间：%s<br />点赞数：%d<br />",
			"https://www.bilibili.com/video/"+bvid,
			comment.Content,
			comment.Uname,
			comment.Mid,
			time.Unix(int64(comment.Ctime), 0).Format("2006-01-02 15:04:05"),
			comment.Like),
	)
	if err != nil {
		return err
	}
	return nil
}
