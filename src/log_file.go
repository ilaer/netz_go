package src

import (
	"fmt"
	"netz_go/util"
)

func (m *SRC) FileLog() {
	for {
		select {
		case logText, ok := <-m.LogChannel:
			if !ok {
				continue
			}
			util.XWarning(fmt.Sprintf("%s", logText))
		}
	}
}
