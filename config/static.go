package config

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
}

type Mail struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	IsSSL    bool     `yaml:"isSSL"`
	UserName string   `yaml:"userName"`
	Password string   `yaml:"password"`
	From     string   `yaml:"from"`
	To       []string `yaml:"to"`
}
