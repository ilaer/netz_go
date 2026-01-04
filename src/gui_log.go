package src

import (
	"fmt"
	"time"
)

func (m *SRC) GuiLog() {
	for {
		select {
		case logText, ok := <-m.GuiLogChannel:
			if !ok {
				continue
			}

			if len(m.MultiLogTextList) <= 50 {
				m.MultiLogTextList = append([]string{fmt.Sprintf("%s    %s", time.Now().Format("2006-01-02 15:04:05"), logText)}, m.MultiLogTextList...)

			} else {
				m.MultiLogTextList = append([]string{fmt.Sprintf("%s    %s", time.Now().Format("2006-01-02 15:04:05"), logText)}, m.MultiLogTextList[0:50]...)

			}
		}
	}
}
