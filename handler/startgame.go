package handler

import (
	"digimon-cli/pbprotocol"
	"fmt"
)

func (dc *digimonCli) StartGameAck(ack *pbprotocol.StartGameAck) error {
	fmt.Println("start game")
	return nil
}
