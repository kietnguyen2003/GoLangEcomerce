package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	MySql  MySqlSetting  `mapstructure:"database"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySqlSetting struct {
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	DbName          string `mapstructure:"dbName"`
	Port            int    `mapstructure:"port"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Level      string `mapstructure:"level"`
	File       string `mapstructure:"file"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}
