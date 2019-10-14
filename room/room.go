package room

import (
	"digimon-cli/player"
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
		PlayerInfos: nil,
		Round:       1,
	}
}