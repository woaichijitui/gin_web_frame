package core

import (
	"flag"
	"fmt"
	"gin_web_frame/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

const ConfigEnv = "config"
const ConfigDefaultFile = "config.yaml"

// Viper 配置
func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//global.CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

// getConfigPath 获取配置文件路径, 优先级: 命令行 > 环境变量 > 默认值
func getConfigPath() (config string) {
	// `-c` flag parse
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config != "" { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", config)
		return
	}
	if env := os.Getenv(ConfigEnv); env != "" { // 判断环境变量 CONFIG
		config = env
		fmt.Printf("您正在使用 %s 环境变量, config 的路径为 %s\n", ConfigEnv, config)
		return
	}

	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		config = ConfigDefaultFile
		fmt.Printf("未配置文件路径或配置文件路径不存在, 使用默认配置文件路径: %s\n", config)
	}

	return
}
