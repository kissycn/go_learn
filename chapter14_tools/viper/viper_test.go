package viper

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
	"time"
)

func TestViper(t *testing.T) {
	viper.SetConfigName("config") // 配置文件名称
	viper.SetConfigType("toml")   // 培训文件类型
	viper.AddConfigPath(".")      // 搜索路径，支持多个，按照代码顺序搜索
	//viper.AddConfigPath("/Users/kissy/Workspace/Go/go_learn/src/ielee.com/go_learn/chapter14_tools/viper/")

	viper.SetDefault("redis.port", 6379)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))

	fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports: ", viper.GetIntSlice("server.ports"))
	fmt.Println("timeout: ", viper.GetDuration("server.timeout"))

	// 返回mysql下所有的key-value
	fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	// 获取redis下所有的key-value
	fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	fmt.Println("all settings: ", viper.AllSettings())
}

func TestViperSetDefault(t *testing.T) {
	viper.SetDefault("name", "xiaoming")
	fmt.Printf(viper.GetString("name"))
}

func TestGetEnv(t *testing.T) {
	env := os.Getenv("GOROOT")
	fmt.Printf(env)
}

func TestEnvByViper(t *testing.T) {
	//表示 先预加载匹配的环境变量
	viper.AutomaticEnv()
	//读取已经加载到default中的环境变量
	if env := viper.Get("GOROOT"); env == nil {
		println("unknown")
	} else {
		fmt.Printf("%#v\n", env)
	}
}

func TestIOReader(t *testing.T) {
	viper.SetConfigType("toml")
	tomlConfig := []byte(`
app_name = "awesome web"

# possible values: DEBUG, INFO, WARNING, ERROR, FATAL
log_level = "DEBUG"

[mysql]
ip = "127.0.0.1"
port = 3306
user = "dj"
password = 123456
database = "awesome"

[redis]
ip = "127.0.0.1"
port = 7381
`)
	err := viper.ReadConfig(bytes.NewBuffer(tomlConfig))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("redis port: ", viper.GetInt("redis.port"))
}

type Config struct {
	AppName  string
	LogLevel string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

func TestViperUnmarshal(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	var c Config
	viper.Unmarshal(&c)

	fmt.Println(c.MySQL)
}

func TestWatcher(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	time.Sleep(time.Second * 10)
	fmt.Println("redis port after sleep: ", viper.Get("redis.port"))
}

func TestConfigSave(t *testing.T) {
	viper.SetConfigName("config1")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.Set("app_name", "awesome web")
	viper.Set("log_level", "DEBUG")
	viper.Set("mysql.ip", "127.0.0.1")
	viper.Set("mysql.port", 3306)
	viper.Set("mysql.user", "root")
	viper.Set("mysql.password", "123456")
	viper.Set("mysql.database", "awesome")

	viper.Set("redis.ip", "127.0.0.1")
	viper.Set("redis.port", 6381)

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}
}
