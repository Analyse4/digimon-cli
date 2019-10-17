package room

import (
	"github.com/Analyse4/digimon-cli/player"
)

type Room struct {
	Id          uint64
	IsStart     bool
	Type        int32
	CurrentNum  uint32
	PlayerInfos []*player.Player
	Round       int32
}

func New() *Room {
	return &Room{
		Id:          0,
		IsStart:     false,
		Type:        0,
		CurrentNum:  0,
		PlayerInfos: make([]*player.Player, 0),
		Round:       1,
	}
}
