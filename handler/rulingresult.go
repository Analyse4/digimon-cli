package handler

import (
	"digimon-cli/pbprotocol"
)

func (dc *digimonCli) RulingResultAck(ack *pbprotocol.RulingResultAck) error {
	if ack.AttackerID != dc.player.ID && ack.TargetID != dc.player.ID {
		return nil
	}
	//rpc battle tui
	return nil
}
