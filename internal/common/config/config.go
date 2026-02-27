package config

type AppConfig struct {
	Mysql struct { //mysql 服务配置
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		Database string `json:"database" yaml:"database"`
	}
	Redis struct { //redis 服务配置
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		Password string `json:"password" yaml:"password"`
		DB       int    `json:"db" yaml:"db"`
	}
	Huyi struct { // Huyi 服务配置
		APIID  string `json:"APIID,omitempty" yaml:"APIID,omitempty"`
		APIKEY string `json:"APIKEY,omitempty" yaml:"APIKEY,omitempty"`
	}
	Zap struct { //zap 日志服务配置
		LogDir      string `json:"logDir" yaml:"logDir"`
		MaxAge      int    `json:"maxAge" yaml:"maxAge"`
		Compress    bool   `json:"compress" yaml:"compress"`
		Level       string `json:"level" yaml:"level"`
		Development bool   `json:"development" yaml:"development"`
	}
}
