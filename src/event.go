package src

func (m *SRC) ScanEvent() {
	go func() {
		m.DataList.DeleteRowAll()
		m.LogLabel.SetText("正在扫描...")
		m.LogLabel.Redraw()
		results := m.Seek(m.Segment, m.Ports)
		for _, result := range results {
			idx := m.DataList.AddRowText("")
			m.DataList.SetItemText(idx, 0, result.([]string)[0])
			m.DataList.SetItemText(idx, 1, result.([]string)[1])
			m.DataList.SetItemText(idx, 2, result.([]string)[2])

		}
		m.ScanButton.Enable(true)
		m.ScanButton.Redraw(true)
		m.DataList.Redraw(true)
		m.LogLabel.SetText("扫描完毕!")
		m.LogLabel.Redraw()
	}()

}

func (m *SRC) DoubleClickEvent(idx int32) {
	if idx < 0 {
		return
	}
	ip := m.DataList.GetItemText(idx, 0)
	port := m.DataList.GetItemText(idx, 1)
	m.GuiLogChannel <- "IP: " + ip + " Port: " + port
}
