package config

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
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
