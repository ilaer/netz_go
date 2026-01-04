package src

import (
	"fmt"
	"net"
	"path/filepath"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func (m *SRC) MainWindow() (err error) {
	xc.SetXcguiPath("ui.dll")
	ap := app.New(true)
	ap.SetWindowIcon(imagex.NewByFile(filepath.Join(m.RootPath, "icon.ico")).Handle)
	//ap.EnableDPI(true)
	//ap.EnableAutoDPI(true)

	ap.SetWindowIcon(2)
	f := font.NewEX("simhei", 13, xcc.FontStyle_Regular)
	ap.SetDefaultFont(f.Handle)

	w := window.New(int32(m.Width*0.1), int32(m.Height*0.15), int32(m.Width), int32(m.Height), "NetZ", 0, xcc.Window_Style_Modal|xcc.Window_Style_Drag_Window)
	w.SetIcon(imagex.NewByFile(filepath.Join(m.RootPath, "icon.ico")).Handle)

	mainLayout := widget.NewLayoutEle(int32(m.Width*0.01), int32(m.Height*0.05), int32(m.Width*0.84), int32(m.Height*0.8), w.Handle)
	homeTabButton := widget.NewButton(0, 0, int32(m.Width*0.065), int32(m.Height*0.05), "首页", mainLayout.Handle)
	homeTabButton.SetTypeEx(xcc.Button_Type_Radio)
	homeTabButton.SetStyle(xcc.Button_Style_Default)
	homeTabButton.SetCheck(true)

	settingTabButton := widget.NewButton(0, 0, int32(m.Width*0.065), int32(m.Height*0.05), "设置", mainLayout.Handle)
	settingTabButton.SetTypeEx(xcc.Button_Type_Radio)
	settingTabButton.SetStyle(xcc.Button_Style_Default)

	layoutPages := widget.NewLayoutEle(int32(m.Width*0.011), int32(m.Height*0.11), int32(m.Width), int32(m.Height), w.Handle)

	layoutPage1 := widget.NewLayoutEle(0, 0, int32(m.Width), int32(m.Height), layoutPages.Handle)
	layoutPage1.Show(true)
	homeTabButton.SetBindEle(layoutPage1.Handle)

	layoutPage2 := widget.NewLayoutEle(0, 0, int32(m.Width), int32(m.Height), layoutPages.Handle)
	layoutPage2.Show(false)
	settingTabButton.SetBindEle(layoutPage2.Handle)

	widget.NewShapeLine(0, int32(m.Height*0.01), int32(m.Width*1), int32(m.Height*0.01), layoutPage1.Handle)
	widget.NewShapeText(0, 0, int32(m.Width*0.07), int32(m.Height*0.05), "网段:", layoutPage1.Handle)

	networkEntry := widget.NewEdit(0, 0, int32(m.Width*0.15), int32(m.Height*0.05), layoutPage1.Handle)
	networkEntry.SetText(m.Segment)

	widget.NewShapeText(0, 0, int32(m.Width*0.02), int32(m.Height*0.05), " ", layoutPage1.Handle)
	widget.NewShapeText(0, 0, int32(m.Width*0.07), int32(m.Height*0.05), "端口:", layoutPage1.Handle)
	portEntry := widget.NewEdit(0, 0, int32(m.Width*0.3), int32(m.Height*0.05), layoutPage1.Handle)
	portEntry.SetText(m.Ports)

	widget.NewShapeText(0, 0, int32(m.Width*0.2), int32(m.Height*0.05), " ", layoutPage1.Handle)
	m.ScanButton = widget.NewButton(0, 0, int32(m.Width*0.1), int32(m.Height*0.05), "开始扫描", layoutPage1.Handle)

	m.ScanButton.Event_BnClick(func(sbHandled *bool) int {
		routeIP := fmt.Sprintf("%s.1", networkEntry.GetTextEx())
		parsedIP := net.ParseIP(routeIP)
		if parsedIP == nil {
			w.MessageBox("错误", "无效的网段格式，请输入正确的网段格式，例如：192.168.1", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, xcc.Window_Style_Modal)
			return 1
		}
		m.ScanButton.Enable(false)
		m.ScanButton.Redraw(true)
		go m.ScanEvent()
		return 0
	})

	widget.NewShapeText(0, 0, int32(m.Width*0.02), int32(m.Height*0.07), " ", layoutPage1.Handle)
	m.DataList = widget.NewList(0, 0, int32(m.Width*0.98), int32(m.Height*0.76), layoutPage1.Handle)
	m.DataList.CreateAdapterHeader()
	m.DataList.CreateAdapters(3)
	m.DataList.SetRowHeightDefault(int32(m.Height*0.05), int32(m.Height*0.05))
	m.DataList.AddColumnText(int32(m.Width*0.25), "name1", "主机")
	m.DataList.AddColumnText(int32(m.Width*0.46), "name2", "端口")
	m.DataList.AddColumnText(int32(m.Width*0.25), "name3", "时间")
	m.DataList.SetSort(0, 0, true)

	m.DataList.Event_LBUTTONDBCLICK(func(nFlags int, pPt *xc.POINT, pbHandled *bool) int {
		// 取鼠标点击的行和列
		var row, column int32
		m.DataList.HitTestOffset(pPt, &row, &column)

		if row < 0 || column < 0 {
			return 0
		}
		cloumnCount := m.DataList.GetCountColumn_AD()
		fmt.Printf("cloumnCount: %d\n", cloumnCount)
		fmt.Printf("nFlags: %d, pPt: (%d, %d), row: %d, column: %d\n", nFlags, pPt.X, pPt.Y, row, column)
		rcText := m.DataList.GetItemText(row, column)
		println("双击行索引:", row, "列索引:", column, "文本内容:", rcText)
		return 0
	})
	m.LogLabel = widget.NewShapeText(0, 0, int32(m.Width*0.98), int32(m.Height*0.05), "提示栏 ", layoutPage1.Handle)
	m.LogLabel.SetTextColor(xc.RGBA(100, 149, 237, 255))
	w.Show(true)
	ap.Run()
	ap.Exit()

	return nil
}
