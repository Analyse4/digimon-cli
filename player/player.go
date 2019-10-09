package player

import "digimon-cli/pbprotocol"

type Hero struct {
	Identity      pbprotocol.DigimonIdentity
	IdentityLevel int32
	SkillPoint    int32
	SkillType     int32
	SkillLevel    int32
	SkillName     string
}

type Player struct {
	ID          uint64
	NickName    string
	LoginType   int32
	DigiMonster *Hero
}

func New() *Player {
	return &Player{
		ID:          0,
		NickName:    "",
		LoginType:   -1,
		DigiMonster: new(Hero),
	}
}
