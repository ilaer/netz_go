package src

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/twgh/xcgui/widget"
	"netz_go/util"
	"path/filepath"
)

type SRC struct {
	Width          float32 //GUI width
	Height         float32 //GUI height
	RootPath       string  `json:"root_path"`
	ConfigFileName string  `json:"config_file_name"`

	Segment string `yaml:"segment"`
	Ports   string `yaml:"ports"`

	MultiLogTextList []string `json:"multi_log_text_list"`

	DataList   *widget.List
	ScanButton *widget.Button
	LogLabel   *widget.ShapeText

	GuiLogChannel chan string `json:"guiLogChannel"`
	LogChannel    chan string `json:"logChannel"`
}

func (m *SRC) Initialize(rootPath string) (err error) {
	confFilePath := filepath.Join(m.RootPath, m.ConfigFileName)
	vip := viper.New()
	//导入配置文件
	vip.SetConfigType("yaml")
	vip.SetConfigFile(confFilePath)
	//读取配置文件
	err = vip.ReadInConfig()
	if err != nil {
		util.XWarning(fmt.Sprintf("ReadInConfig error: %v\n", err.Error()))
		return err
	}

	//将配置文件读到结构体中
	err = vip.Unmarshal(&m)
	if err != nil {
		util.XWarning(fmt.Sprintf("vip.Unmarshal error: %v\n", err))
		return err
	}
	m.GuiLogChannel = make(chan string)
	m.LogChannel = make(chan string)

	m.MultiLogTextList = []string{}

	return nil
}
