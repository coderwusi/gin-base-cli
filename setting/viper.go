package setting

import (
	"fmt"
	"github.com/spf13/viper" //第三方配置管理神器
)

var Conf = new(TotalConfig)

type TotalConfig struct {
	*MySQLConfig `mapstructure:"mysql"`
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

func Init() {

	viper.SetConfigFile("config.yaml") //读取配置
	//viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	} //捕捉viper读取错误

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return
	} //捕捉viper反序列化错误
	//
	fmt.Println("config.yaml读取反序列化成功")

}
