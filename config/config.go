package config

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Logging  Logging  `mapstructure:"logging"`
}

type Server struct {
	HTTPPort               int    `mapstructure:"http_port"`
	Host                   string `mapstructure:"host"`
	TokenExpMinutes        uint   `mapstructure:"token_exp_minutes"`
	RefreshTokenExpMinutes uint   `mapstructure:"refresh_token_exp_minutes"`
	TokenSecret            string `mapstructure:"token_secret"`
}

type Database struct {
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DBName string `mapstructure:"db_name"`
}
type LoggerConfig struct {
	Name        string `mapstructure:"name"`
	LogFilePath string `mapstructure:"log_file_path"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
	Level       string `mapstructure:"level"`
}

type Logging struct {
	LokiURL string         `mapstructure:"loki_url"`
	Loggers []LoggerConfig `mapstructure:"loggers"`
}
