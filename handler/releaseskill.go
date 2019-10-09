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
	if (pl.DigiMonster.SkillPoint >= 2 && skillTUI.Result() == skillTUI.ATTACK) || skillTUI.Result() == skillTUI.DEFENCE || (pl.DigiMonster.SkillPoint >= 2 && skillTUI.Result() == skillTUI.EVOLVE) {
		skillTUI.SecondShow(pl.DigiMonster)
		skillTUI.SkillLevel = skillTUI.Result()
	}
	switch skillTUI.SkillType {
	case skillTUI.POWERUP:
		req.SkillType = int32(skillTUI.POWERUP)
	case skillTUI.DEFENCE:
		req.SkillType = int32(skillTUI.DEFENCE)
		req.SkillLevel = int32(skillTUI.SkillLevel)
	case skillTUI.ESCAPE:
		req.SkillType = int32(skillTUI.ESCAPE)
	case skillTUI.QUIT:
		req.SkillType = int32(skillTUI.QUIT)
	case skillTUI.ATTACK:
		req.SkillType = int32(skillTUI.ATTACK)
		req.SkillLevel = int32(skillTUI.SkillLevel)
	case skillTUI.EVOLVE:
		req.SkillType = int32(skillTUI.EVOLVE)
		req.SkillLevel = int32(skillTUI.SkillLevel)
	default:
		fmt.Println("invalid input, please reselect")
		skillTUI.Refresh()
		ReleaseSkillReq(conn, pl, skillTUI)
	}
	conn.Send("releaseskill", req)
}
