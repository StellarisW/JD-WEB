package boot

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"main/app/global"
	"main/utils"
	"os"
	"path/filepath"
	"time"
)

func ViperSetup(path ...string) {
	var configPath string
	// 获取Config文件路径
	if len(path) == 0 {
		flag.StringVar(&configPath, "c", "", "choose configPath file.")
		flag.Parse()
		// 优先级: 命令行 > 环境变量 > 默认值
		if configPath == "" {
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				configPath = utils.ConfigFile
				fmt.Printf("Using the default value of configPath, Path: %v\n", utils.ConfigFile)
			} else {
				configPath = configEnv
				fmt.Printf("Using g.Config, Path: %v\n", configPath)
			}
		} else {
			fmt.Printf("Using the value pass by - C parameter of the command line, Path: %v\n", configPath)
		}
	} else {
		configPath = path[0]
		fmt.Printf("Using the value pass by func Viper(), Path: %v\n", configPath)
	}

	// 创建Viper对象
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configPath file, err: %s\n", err))
	}

	//实时读取配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("configPath file changed:", e.Name)
		if err := v.Unmarshal(&g.Config); err != nil {
			panic(fmt.Errorf("Unmarshal conf failed, err:%s \n", err))
		}
	})

	//反序列化到g.Config结构体
	if err := v.Unmarshal(&g.Config); err != nil {
		panic(fmt.Errorf("Unmarshal conf failed, err:%s \n", err))
	}

	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	g.Config.AutoCode.Root, _ = filepath.Abs("..")
	g.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(g.Config.Auth.JWT.ExpiresTime)),
	)
	g.VP = v
}
