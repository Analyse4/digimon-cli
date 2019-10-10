package handler

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/peer"
	"digimon-cli/player"
	"digimon-cli/tui"
	"fmt"
)

func ReleaseSkillReq(conn peer.Connection, pl *player.Player, skillTUI *tui.Skill) {
	req := new(pbprotocol.ReleaseSkillReq)
	skillTUI.Show()
	skillTUI.SkillType = skillTUI.Result()
	if skillTUI.IsSkillTypValid() {
		if (pl.DigiMonster.SkillPoint >= 2 && skillTUI.Result() == skillTUI.ATTACK) || skillTUI.Result() == skillTUI.DEFENCE || (pl.DigiMonster.SkillPoint >= 2 && skillTUI.Result() == skillTUI.EVOLVE) {
			secondShow(skillTUI, pl.DigiMonster, req, conn)
		}
		req.SkillType = int32(skillTUI.SkillType)
		conn.Send("digimon.releaseskill", req)
	} else {
		fmt.Println("invalid input, please reselect")
		skillTUI.Refresh()
		skillTUI.RefreshType()
		ReleaseSkillReq(conn, pl, skillTUI)
	}
}

func (dc *digimonCli) ReleaseSkillAck(ack *pbprotocol.ReleaseSkillAck) error {
	fmt.Println(ack)
	//TODO: update status
	//TODO: release skill req
	return nil
}

func secondShow(skillTUI *tui.Skill, h *player.Hero, req *pbprotocol.ReleaseSkillReq, conn peer.Connection) {
	skillTUI.ShowLevel(h)
	skillTUI.SkillLevel = skillTUI.Result()
	if skillTUI.IsSkillLevelValid(h) {
		req.SkillLevel = int32(skillTUI.SkillLevel)
	} else {
		fmt.Println("invalid input, please reselect")
		skillTUI.Refresh()
		skillTUI.RefreshLevel()
		secondShow(skillTUI, h, req, conn)
	}
}
