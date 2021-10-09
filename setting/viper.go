package setting

import (
	"fmt"
	"github.com/spf13/viper" //第三方配置管理神器
	"log"
)

var Conf = new(TotalConfig)

type TotalConfig struct {
	Mode         string `mapstructure:"mode"`
	*MySQLConfig `mapstructure:"mysql"`
	*LogConfig   `mapstructure:"log"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	Suffix       string `mapstructure:"suffix"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func ViperInit() {
	log.SetFlags(log.LstdFlags | log.Llongfile) //log显示行号

	viper.SetConfigFile("config.yaml") //读取配置
	//viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("viper.ReadInConfig failed, err:", err)
	} //捕捉viper读取错误

	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatalln("viper.Unmarshal failed, err:", err)
	} //捕捉viper反序列化错误
	//
	fmt.Println("config.yaml读取、反序列化成功")
}
