package player

import (
	"fmt"
	"github.com/Analyse4/digimon-cli/pbprotocol"
)

type Hero struct {
	Identity      pbprotocol.DigimonIdentity
	IdentityLevel int32
	SkillPoint    int32
	SkillType     int32
	SkillLevel    int32
	SkillName     string
	IsDead        bool
}

type Player struct {
	ID          uint64
	NickName    string
	LoginType   int32
	DigiMonster *Hero
	RoomID      uint64
	Seat        int32
}

func New() *Player {
	return &Player{
		ID:          0,
		NickName:    "",
		LoginType:   -1,
		DigiMonster: new(Hero),
		RoomID:      0,
		Seat:        -1,
	}
}

//Debug
func (pl *Player) Show() {
	fmt.Printf("ID: %d\n", pl.ID)
	fmt.Printf("Name: %s\n", pl.NickName)
	fmt.Printf("RoomID: %d\n", pl.RoomID)
	fmt.Printf("Seat: %d\n", pl.Seat)
	fmt.Printf("Dead: %v\n", pl.DigiMonster.IsDead)
	fmt.Printf("Identity: %d\n", pl.DigiMonster.Identity)
	fmt.Printf("IdentityLevel: %d\n", pl.DigiMonster.IdentityLevel)
	fmt.Printf("SkillType: %d\n", pl.DigiMonster.SkillType)
	fmt.Printf("SkillPoint: %d\n", pl.DigiMonster.SkillPoint)
	fmt.Printf("SkillLevel: %d\n", pl.DigiMonster.SkillLevel)
	fmt.Printf("SkillName: %s\n", pl.DigiMonster.SkillName)
}
