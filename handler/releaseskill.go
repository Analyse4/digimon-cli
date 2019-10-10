package handler

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/peer"
	"digimon-cli/tui"
	"fmt"
	"github.com/sirupsen/logrus"
)

func ReleaseSkillReq(conn peer.Connection, skillTUI *tui.Skill) {
	req := new(pbprotocol.ReleaseSkillReq)
	skillTUI.Show()
	skillTUI.SkillType = skillTUI.Result()
	if skillTUI.IsSkillTypValid() {
		if (skillTUI.H.SkillPoint >= 2 && skillTUI.Result() == skillTUI.ATTACK) || skillTUI.Result() == skillTUI.DEFENCE || (skillTUI.H.SkillPoint >= 2 && skillTUI.Result() == skillTUI.EVOLVE) {
			secondShow(skillTUI, req, conn)
		} else {
			skillTUI.SkillLevel = 1
		}
		req.SkillType = int32(skillTUI.SkillType)
		req.SkillLevel = int32(skillTUI.SkillLevel)
		conn.Send("digimon.releaseskill", req)
	} else {
		fmt.Println("invalid input, please reselect")
		skillTUI.Refresh()
		skillTUI.RefreshType()
		ReleaseSkillReq(conn, skillTUI)
	}
}

func (dc *digimonCli) ReleaseSkillAck(ack *pbprotocol.ReleaseSkillAck) error {
	if ack.Base.Result != 0 {
		logrus.Println(ack.Base.Msg)
		return fmt.Errorf(ack.Base.Msg)
	}
	dc.player.DigiMonster.Identity = ack.Hero.Identity
	dc.player.DigiMonster.IdentityLevel = ack.Hero.IdentityLevel
	dc.player.DigiMonster.SkillType = ack.Hero.SkillType
	dc.player.DigiMonster.SkillLevel = ack.Hero.SkillLevel
	dc.player.DigiMonster.SkillName = ack.Hero.SkillName
	dc.player.DigiMonster.SkillPoint = ack.Hero.SkillPoint
	return nil
}

func secondShow(skillTUI *tui.Skill, req *pbprotocol.ReleaseSkillReq, conn peer.Connection) {
	skillTUI.ShowLevel()
	skillTUI.SkillLevel = skillTUI.Result()
	if skillTUI.IsSkillLevelValid() {
		req.SkillLevel = int32(skillTUI.SkillLevel)
	} else {
		fmt.Println("invalid input, please reselect")
		skillTUI.Refresh()
		skillTUI.RefreshLevel()
		secondShow(skillTUI, req, conn)
	}
}
