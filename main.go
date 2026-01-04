package main

import (
	gcmd "gbase/cmd"
	"github.com/sqweek/dialog"
	"netz_go/src"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

const (
	SM_CXSCREEN = uintptr(0) // X Size of screen
	SM_CYSCREEN = uintptr(1) // Y Size of screen
)

func main() {
	width_, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(SM_CXSCREEN)
	height_, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(SM_CYSCREEN)
	var width float32
	var height float32
	width = 1024.0
	height = 768.0
	if width_ > 1024 {
		width = float32(width_) * 0.84
		height = float32(height_) * 0.8
	}
	args := os.Args //[]string{"", "download_basic"}
	rootPath, _ := os.Getwd()
	/*判断配置文件*/
	configFileName := "config.yaml"
	if len(args) > 1 {
		if strings.Contains(args[1], ".yaml") == true {
			configFileName = args[1]
		}
	}
	exists, _ := gcmd.FileExists(filepath.Join(rootPath, configFileName))
	if exists == false {
		dialog.Message("%s", "找不到配置文件 : "+configFileName).Title("提示").Error()
	}

	m := &src.SRC{
		Width:          width,
		Height:         height,
		RootPath:       rootPath,
		ConfigFileName: configFileName,
	}
	m.Initialize(rootPath)
	go m.GuiLog()
	go m.FileLog()

	err := m.MainWindow()
	if err != nil {
		panic(err)
	}
}
