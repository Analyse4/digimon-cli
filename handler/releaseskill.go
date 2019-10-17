package handler

import (
	"fmt"
	"github.com/Analyse4/digimon-cli/pbprotocol"
	"github.com/Analyse4/digimon-cli/peer"
	"github.com/Analyse4/digimon-cli/tui"
	"github.com/sirupsen/logrus"
)

func ReleaseSkillReq(conn peer.Connection, skillTUI *tui.Skill) {
	req := new(pbprotocol.ReleaseSkillReq)
	req.SkillTargets = make([]uint64, 0)
	skillTUI.Show()
	skillTUI.SkillType = skillTUI.Result()
	if skillTUI.IsSkillTypValid() {
		if (skillTUI.H.SkillPoint >= 2 && skillTUI.Result() == skillTUI.ATTACK) || skillTUI.Result() == skillTUI.DEFENCE || (skillTUI.H.SkillPoint >= 2 && skillTUI.Result() == skillTUI.EVOLVE) {
			secondShow(skillTUI)
		} else {
			skillTUI.SkillLevel = 1
		}
		if skillTUI.SkillType == skillTUI.ATTACK {
			attackTargetShow(skillTUI)
		}
		req.SkillType = int32(skillTUI.SkillType)
		req.SkillLevel = int32(skillTUI.SkillLevel)
		req.SkillTargets = append(req.SkillTargets, skillTUI.SkillTargets...)
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

func secondShow(skillTUI *tui.Skill) {
	skillTUI.ShowLevel()
	skillTUI.SkillLevel = skillTUI.Result()
	if !skillTUI.IsSkillLevelValid() {
		fmt.Println("invalid input, please reselect")
		skillTUI.Refresh()
		skillTUI.RefreshLevel()
		secondShow(skillTUI)
	}
}

func attackTargetShow(skillTUI *tui.Skill) {
	skillTUI.ShowAttackTarget()
	skillTUI.SkillTargets = append(skillTUI.SkillTargets, skillTUI.RoomIDList[int(skillTUI.Result())])
	if !skillTUI.IsSkillTargetValid() {
		fmt.Println("invalid input, please reselect")
		skillTUI.Refresh()
		skillTUI.RefreshSkillTargets()
		attackTargetShow(skillTUI)
	}
}
