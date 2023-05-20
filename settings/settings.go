package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// @file      : settings.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"` // mapstructure viper中使用的tag
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*Auth        `mapstructure:"auth"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	DBName             string `mapstructure:"dbname"`
	MaxOpenConnections int    `mapstructure:"max_open_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
}

type Auth struct {
	JwtExpire int `mapstructure:"jwt_expire"`
}

// Init 使用viper初始化配置文件
func Init(configFilePath string) (err error) {
	viper.SetConfigFile(configFilePath)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%#v\n", err)
		return
	}

	// 将读取到配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%#v\n", err)
	}

	// 实时监听配置文件的修改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("%s is modify\n", configFilePath)
		// 配置文件修改后，继续反序列化到Conf中
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%#v\n", err)
		}
	})
	return
}
