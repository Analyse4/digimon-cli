package handler

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/tui"
	"fmt"
)

func (dc *digimonCli) RoundResultAck(ack *pbprotocol.RoundResultAck) error {
	if ack.IsEnd {
		// settlement
	} else {
		dc.room.Round = ack.Round

		fmt.Printf("ROUND: %d\n", dc.room.Round)
		dc.player.Show()

		skillTUI := tui.NewSkill()
		skillTUI.SetHero(dc.player.DigiMonster)
		ReleaseSkillReq(dc.conn, skillTUI)
	}
	return nil
}
