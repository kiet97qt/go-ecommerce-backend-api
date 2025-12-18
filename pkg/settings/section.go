package settings

type Config struct {
	Server    Server      `mapstructure:"server"`
	Databases []Databases `mapstructure:"databases"`
	Security  Security    `mapstructure:"security"`
	MySQL     MySQL       `mapstructure:"mysql"`
	Logging   Logging     `mapstructure:"logging"`
	Redis     Redis       `mapstructure:"redis"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type Databases struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type Security struct {
	JWT struct {
		Key string `mapstructure:"key"`
	} `mapstructure:"jwt"`
}

type MySQL struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type Logging struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Compress   bool   `mapstructure:"compress"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
