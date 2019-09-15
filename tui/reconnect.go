package tui

import "fmt"

type ReconnectTUI struct {
	RECONNECT uint8
	QUIT      uint8
	result    uint8
}

func (l *ReconnectTUI) Show() {
	fmt.Println("Encounter service issues, please chose Reconnect or Quit, input number only")
	fmt.Println("1.Reconnect		2.Quit")
	fmt.Scan(&l.result)
}

func (l *ReconnectTUI) Result() uint8 {
	return l.result
}

func NewReconnect() *ReconnectTUI {
	return &ReconnectTUI{
		RECONNECT: 1,
		QUIT:      2,
	}
}

func (l *ReconnectTUI) Refresh() {
	l.result = 0
}
