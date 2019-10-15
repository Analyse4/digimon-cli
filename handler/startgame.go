package handler

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/player"
	"digimon-cli/tui"
	"fmt"
)

func (dc *digimonCli) StartGameAck(ack *pbprotocol.StartGameAck) error {
	dc.room.Id = ack.RoomInfo.RoomId
	dc.room.CurrentNum = ack.RoomInfo.CurrentPlayerNum
	dc.room.Type = int32(ack.RoomInfo.Type)
	dc.room.IsStart = ack.RoomInfo.IsStart
	dc.room.Round = ack.RoomInfo.Round
	for _, v := range ack.RoomInfo.PlayerInfos {
		if v.Id == dc.player.ID {
			dc.player.DigiMonster.Identity = v.Hero.Identity
			dc.player.DigiMonster.IdentityLevel = v.Hero.IdentityLevel
			dc.player.DigiMonster.SkillPoint = v.Hero.SkillPoint
			dc.player.DigiMonster.SkillType = v.Hero.SkillType
			dc.player.DigiMonster.SkillLevel = v.Hero.SkillLevel
			dc.player.DigiMonster.SkillName = v.Hero.SkillName
			dc.player.DigiMonster.IsDead = v.Hero.IsDead
		}
	}

	for i, v := range ack.RoomInfo.PlayerInfos {
		if IsNewPlayer(v, dc.room.PlayerInfos) {
			tmpPlayer := new(player.Player)
			tmpPlayer.ID = v.Id
			tmpPlayer.NickName = v.Nickname
			tmpPlayer.RoomID = v.RoomId
			tmpPlayer.Seat = v.Seat
			tmpPlayer.DigiMonster = new(player.Hero)
			if int(v.Seat) < len(dc.room.PlayerInfos) {
				dc.room.PlayerInfos[v.Seat] = tmpPlayer
			} else {
				dc.room.PlayerInfos = append(dc.room.PlayerInfos, tmpPlayer)
			}
		} else {
			dc.room.PlayerInfos[i].DigiMonster.Identity = v.Hero.Identity
			dc.room.PlayerInfos[i].DigiMonster.IdentityLevel = v.Hero.IdentityLevel
			dc.room.PlayerInfos[i].DigiMonster.SkillPoint = v.Hero.SkillPoint
			dc.room.PlayerInfos[i].DigiMonster.SkillType = v.Hero.SkillType
			dc.room.PlayerInfos[i].DigiMonster.SkillLevel = v.Hero.SkillLevel
			dc.room.PlayerInfos[i].DigiMonster.SkillName = v.Hero.SkillName
			dc.room.PlayerInfos[i].DigiMonster.IsDead = v.Hero.IsDead
		}
	}

	fmt.Printf("Round: %v\n", dc.room.Round)
	dc.player.Show()
	skillTUI := tui.NewSkill()
	for _, pl := range dc.room.PlayerInfos {
		if pl.ID != dc.player.ID {
			skillTUI.RoomIDList = append(skillTUI.RoomIDList, pl.ID)
		}
	}
	skillTUI.SetHero(dc.player.DigiMonster)
	ReleaseSkillReq(dc.conn, skillTUI)
	return nil
}
