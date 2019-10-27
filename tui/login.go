package tui

import "fmt"

type Login struct {
	LOGIN  uint8
	QUIT   uint8
	result uint8
}

func (l *Login) Show() {
	fmt.Println("please choose Login or Quit, input number only")
	fmt.Println("1.Login		2.Quit")
	fmt.Scan(&l.result)
}

func (l *Login) Result() uint8 {
	return l.result
}

func NewLogin() *Login {
	return &Login{
		LOGIN: 1,
		QUIT:  2,
	}
}

func (l *Login) Refresh() {
	l.result = 0
}
