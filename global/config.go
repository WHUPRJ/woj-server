package global

type ConfigWebServer struct {
	Address       string `yaml:"Address"`
	Port          int    `yaml:"Port"`
	JwtSigningKey string `yaml:"JwtSigningKey"`
	JwtExpireHour int    `yaml:"JwtExpireHour"`
}

type ConfigRedis struct {
	Db       int    `yaml:"Db"`
	QueueDb  int    `yaml:"QueueDb"`
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
}

type ConfigDatabase struct {
	Host            string `yaml:"Host"`
	Port            int    `yaml:"Port"`
	User            string `yaml:"User"`
	Password        string `yaml:"Password"`
	Database        string `yaml:"Database"`
	Prefix          string `yaml:"Prefix"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	ConnMaxLifetime int    `yaml:"ConnMaxLifetime"`
}

type ConfigMetrics struct {
	Namespace string `yaml:"Namespace"`
	Subsystem string `yaml:"Subsystem"`
}

type Config struct {
	WebServer   ConfigWebServer `yaml:"WebServer"`
	Redis       ConfigRedis     `yaml:"Redis"`
	Database    ConfigDatabase  `yaml:"Database"`
	Metrics     ConfigMetrics   `yaml:"Metrics"`
	Development bool            `yaml:"Development"`
}
