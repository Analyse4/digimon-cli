package tui

import "fmt"

type RPCBattle struct {
	Rock        int32
	Paper       int32
	Scissors    int32
	Identity    int32
	OtherSideID uint64
	result      int32
}

func (l *RPCBattle) Show() {
	fmt.Println("please chose Rock, Paper or Scissors, input number only")
	fmt.Println("1.Rock		2.Paper		3.Scissors")
	fmt.Scan(&l.result)
}

func (l *RPCBattle) Result() int32 {
	return l.result
}

func NewRPCBattle(i int32, otherID uint64) *RPCBattle {
	return &RPCBattle{
		Rock:        1,
		Paper:       2,
		Scissors:    3,
		Identity:    i,
		OtherSideID: otherID,
		result:      0,
	}
}

func (l *RPCBattle) Refresh() {
	l.Identity = 0
	l.result = 0
}
