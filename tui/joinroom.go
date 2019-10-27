package tui

import "fmt"

type JoinRoom struct {
	JOINROOM uint8
	QUIT     uint8
	result   uint8
}

func (l *JoinRoom) Show() {
	fmt.Println("please choose JoinRoom or Quit, input number only")
	fmt.Println("1.JoinRoom		2.Quit")
	fmt.Scan(&l.result)
}

func (l *JoinRoom) Result() uint8 {
	return l.result
}

func NewJoinRoom() *JoinRoom {
	return &JoinRoom{
		JOINROOM: 1,
		QUIT:     2,
	}
}

func (l *JoinRoom) Refresh() {
	l.result = 0
}
