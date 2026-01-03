package singleton

import (
	"errors"
	"log"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/ducthangng/geofleet/user-service/service/copier"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServerHost        string `mapstructure:"SERVER_HOST"`
	Port              string `mapstructure:"PORT"`
	ReadTimeout       int    `mapstructure:"READ_TIMEOUT"`
	ReadHeaderTimeout int    `mapstructure:"READ_HEADER_TIMEOUT"`
	WriteTimeout      int    `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout       int    `mapstructure:"IDLE_TIMEOUT"`
	MaxHeaderBytes    int    `mapstructure:"MAX_HEADER_BYTES"`
	HTTPDomain        string `mapstructure:"DOMAIN"`
}

type DatabaseConfig struct {
	Type            string `mapstructure:"TYPE"`
	User            string `mapstructure:"DB_USER"`
	Password        string `mapstructure:"PASSWORD"`
	Host            string `mapstructure:"HOST"`
	Name            string `mapstructure:"NAME"`
	Port            string `mapstructure:"PORT"`
	SSLMode         string `mapstructure:"SSL_MODE"`
	CACERTBASE64    string `mapstructure:"CACERT_BASE64"`
	MaxIdleConns    int    `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns    int    `mapstructure:"MAX_OPEN_CONNS"`
	MaxConnLifeTime int    `mapstructure:"MAX_CONN_LIFE_TIME"`
	MaxConnIdleTime int    `mapstructure:"MAX_CONN_IDLE_TIME"`
}

type CookieConfig struct {
	CookieName     string `mapstructure:"COOKIE_NAME"`
	CookieDomain   string `mapstructure:"COOKIE_DOMAIN"`
	CookieHTTPOnly bool   `mapstructure:"COOKIE_HTTP_ONLY"`
	CookieSecure   bool   `mapstructure:"COOKIE_SECURE"`
	MaxAge         int    `mapstructure:"MAX_AGE"`
	JWTKey         string `mapstructure:"JWT_KEY"`

	HTTPSameSiteOption int `mapstructure:"SAME_SITE_OPTION"`
}

type HTTPConfig struct {
}

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
	Cookie CookieConfig
	HTTP   HTTPConfig
}

var (
	centralConfig *Config
	once          sync.Once
	configMu      sync.Mutex
	serverEnv     = "server.dev.env"
	dbEnv         = "db.dev.env"
	cookieEnv     = "cookie.dev.env"
)

func ReadConfig(devenv string) Config {
	once.Do(func() {
		log.Println("initialized CentralConfig")
		if centralConfig == nil {
			centralConfig = &Config{}
		}

		switch devenv {
		case "development":
			readDevelopmentEnv()

		case "production":
			readDevelopmentEnv()

		default:
			panic(errors.New("need to have config environment"))
		}
	})

	// return the variable only

	var config Config
	copier.MustCopy(&config, centralConfig)

	return config
}

func GetConfig() Config {
	configMu.Lock()
	defer configMu.Unlock()

	var config Config
	copier.MustCopy(&config, centralConfig)

	return config
}

func readDevelopmentEnv() {
	viper.SetConfigType("env")
	viper.AddConfigPath("./environment")
	viper.AutomaticEnv()

	readServerConfig()
	readDBConfig()
	readCookieConfig()
}

func readServerConfig() {
	// read Server Config
	viper.SetConfigName(serverEnv)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&centralConfig.Server); err != nil {
		panic(err)
	}
	// LogString(centralConfig.Server)
}

func readDBConfig() {
	// read Server Config
	viper.SetConfigName(dbEnv)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&centralConfig.DB); err != nil {
		panic(err)
	}

	// LogString(centralConfig.DB)
}

func readCookieConfig() {
	// read Server Config
	viper.SetConfigName(cookieEnv)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&centralConfig.Cookie); err != nil {
		panic(err)
	}

	LogString(centralConfig.Cookie)
}

func LogString(data any) {
	stringtype, _ := sonic.MarshalString(data)
	log.Println("config ne: ", stringtype)
}
