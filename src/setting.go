package src

import (
	"fmt"
	"gbase/glog"
	"github.com/spf13/viper"
	"path/filepath"
)

func (m *SRC) LoadConfig() (err error) {
	confFilePath := filepath.Join(m.RootPath, m.ConfigFileName)
	vip := viper.New()
	//导入配置文件
	vip.SetConfigType("yaml")
	vip.SetConfigFile(confFilePath)
	//读取配置文件
	err = vip.ReadInConfig()
	if err != nil {
		glog.XWarning(fmt.Sprintf("LoadConfig error: %v\n", err.Error()))
		return err
	}

	//将配置文件读到结构体中
	err = vip.Unmarshal(&m)
	if err != nil {
		glog.XWarning(fmt.Sprintf("vip.Unmarshal error: %v\n", err))
		return err
	}
	return nil
}

func (m *SRC) UpdateConfig(data map[string]interface{}) (err error) {
	confFilePath := filepath.Join(m.RootPath, "config.yaml")
	vip := viper.New()
	//导入配置文件
	vip.SetConfigType("yaml")
	vip.SetConfigFile(confFilePath)
	err = vip.ReadInConfig()
	if err != nil {
		glog.XWarning(fmt.Sprintf("无法读取配置文件 %s : %v\n", "config.yaml", err))
		return
	}
	for k, v := range data {
		vip.Set(k, v)
	}
	err = vip.WriteConfig()
	if err != nil {
		glog.XWarning(fmt.Sprintf("无法写入配置文件 %s : %v\n", "config.yaml", err))
		return
	}
	return nil
}
