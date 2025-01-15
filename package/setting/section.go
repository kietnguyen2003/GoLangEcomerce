package setting

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	MySql MySqlSetting `mapstructure:"database"`
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
