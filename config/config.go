package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	MysqlConf = Mysql{}
	MailConf  = Mail{}
)

func handlePanic() {
	r := recover()
	if r != nil {
		log.Printf("hot update failed: %v", r)
	}
}

func Init() error {
	err := loadConf()
	if err != nil {
		return err
	}
	return nil
}

func loadConf() error {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")

	err := vp.ReadInConfig()
	if err != err {
		return err
	}

	err = vp.UnmarshalKey("mysql", &MysqlConf)
	if err != nil {
		return err
	}
	err = vp.UnmarshalKey("mail", &MailConf)
	if err != nil {
		return err
	}

	// 热更新
	defer handlePanic()
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		if err = vp.ReadInConfig(); err != nil {
			panic(err)
		}
		_ = vp.UnmarshalKey("database", &MysqlConf)
	})

	return nil
}
