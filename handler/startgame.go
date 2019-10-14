package handler

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/tui"
)

func (dc *digimonCli) StartGameAck(ack *pbprotocol.StartGameAck) error {
	dc.room.Id = ack.RoomInfo.RoomId
	dc.room.CurrentNum = ack.RoomInfo.CurrentPlayerNum
	dc.room.Type = int32(ack.RoomInfo.Type)
	dc.room.IsStart = ack.RoomInfo.IsStart
	for _, v := range ack.RoomInfo.PlayerInfos {
		if v.Id == dc.player.ID {
			dc.player.DigiMonster.Identity = v.Hero.Identity
			dc.player.DigiMonster.IdentityLevel = v.Hero.IdentityLevel
			dc.player.DigiMonster.SkillPoint = v.Hero.SkillPoint
			dc.player.DigiMonster.SkillType = v.Hero.SkillType
			dc.player.DigiMonster.SkillLevel = v.Hero.SkillLevel
			dc.player.DigiMonster.SkillName = v.Hero.SkillName
		}
		for _, rp := range dc.room.PlayerInfos {
			if v.Id == rp.ID {
				rp.DigiMonster.Identity = v.Hero.Identity
				rp.DigiMonster.IdentityLevel = v.Hero.IdentityLevel
				rp.DigiMonster.SkillPoint = v.Hero.SkillPoint
				rp.DigiMonster.SkillType = v.Hero.SkillType
				rp.DigiMonster.SkillLevel = v.Hero.SkillLevel
				rp.DigiMonster.SkillName = v.Hero.SkillName
			}
		}
	}
	dc.player.Show()
	skillTUI := tui.NewSkill()
	skillTUI.SetHero(dc.player.DigiMonster)
	ReleaseSkillReq(dc.conn, skillTUI)
	return nil
}