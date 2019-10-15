package handler

import "digimon-cli/pbprotocol"

func (dc *digimonCli) PlayerLeaveAck(ack *pbprotocol.PlayerLeaveAck) error {
	if dc.room.Id != ack.RoomId {
		return nil
	}
	dc.room.PlayerInfos[ack.Seat] = nil
	return nil
}
