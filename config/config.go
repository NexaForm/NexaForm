package config

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Server struct {
	HttpPort int    `mapstructure:"http_port"`
	Host     string `mapstructre:"host"`
}

type Database struct {
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	DBName string `mapstructure:"db_name"`
}
