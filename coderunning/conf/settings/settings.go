package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() error {
	// 读取配置文件
	viper.SetConfigName("config") // 配置文件名称
	viper.SetConfigType("yaml")   // 如果配置文件中没有拓展名，需要配置此项
	//viper.SetConfigFile("../config.yaml")
	// 会从多个地方寻找配置文件
	viper.AddConfigPath("./etc")  // 在工作目录中查找配置文件，现在当前目录找，找不到走下面的目录，添加多个路径
	viper.AddConfigPath("../etc") // 在工作目录中查找配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			fmt.Println("配置文件未找到")
			return err
		} else {
			// 其他错误
			fmt.Println("加载配置文件错误")
			return err
		}
	}

	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error(err.Error())
	}

	// 支持配置热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件修改了...\n")
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error(err.Error())
		}
	})

	return nil
}
