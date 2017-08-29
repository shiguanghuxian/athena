package config

/* 这里定义了配置所需的结构体 */

// DbConfig 数据库配置
type DbConfig struct {
	Debug   bool   `toml:"debug"`
	DbType  string `toml:"db_type"`
	Address string `toml:"address"`
	Port    int    `toml:"port"`
	User    string `toml:"user"`
	Passwd  string `toml:"passwd"`
	DbName  string `toml:"db_name"`
	Prefix  string `toml:"prefix"`
	MaxIdle int    `toml:"max_idle"`
	MaxConn int    `toml:"max_conn"`
}

// NsqConfig nsq消息队列配置
type NsqConfig struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Address string   `toml:"address"`
	Port    int      `toml:"port"`
	Ping    Duration `toml:"ping"`
	MaxIdle int      `toml:"max_idle"`
}

// InfluxdbConfig influxdb配置文件
type InfluxdbConfig struct {
	Address   string   `toml:"address"`
	Port      int      `toml:"port"`
	DbName    string   `toml:"db_name"`
	SubDbName string   `toml:"sub_db_name"`
	Username  string   `toml:"username"`
	Password  string   `toml:"password"`
	Timeout   Duration `toml:"timeout"`
	MaxIdle   int      `toml:"max_idle"`
}

// SMTPConfig 短信发送配置
type SMTPConfig struct {
	Address  string `toml:"address"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	From     string `toml:"from"`
	Cc       string `toml:"cc"`
}

// AlidayuConfig 短信配置
type AlidayuConfig struct {
	AppKey       string `toml:"app_key"`
	AppSecret    string `toml:"app_secret"`
	SignName     string `toml:"sign_name"`
	TemplateCode string `toml:"template_code"`
}

// WechatConfig 微信配置
type WechatConfig struct {
	AppID         string `toml:"app_id"`
	AppSecret     string `toml:"app_secret"`
	EncodedAesKey string `toml:"encoded_aes_key"`
	Token         string `toml:"token"`
	TemplateID    string `toml:"template_id"`
	ServerURL     string `toml:"server_url"`
}

// HTTPListenerConfig http监听配置
type HTTPListenerConfig struct {
	Enable  bool     `toml:"enable"`
	Address string   `toml:"address"`
	Port    int      `toml:"port"`
	Timeout Duration `toml:"timeout"`
}

// TCPListenerConfig http监听配置
type TCPListenerConfig struct {
	Enable  bool     `toml:"enable"`
	Address string   `toml:"address"`
	Port    int      `toml:"port"`
	Timeout Duration `toml:"timeout"`
}

// LoggerConfig 日志记录对象配置
type LoggerConfig struct {
	Name      string `toml:"name"`
	Verbose   bool   `toml:"verbose"`
	SystemLog bool   `toml:"system_log"`
}
