package handler

import (
	"fmt"
	"github.com/Analyse4/digimon-cli/pbprotocol"
)

func (dc *digimonCli) EndGameAck(ack *pbprotocol.EndGameAck) error {
	if ack.WinnerId == dc.player.ID {
		fmt.Println("Game Over, You Win")
	} else {
		fmt.Println("Game Over, You Lose")
	}
	return nil
}
