package tui

import (
	"fmt"
	"github.com/Analyse4/digimon-cli/pbprotocol"
	"github.com/Analyse4/digimon-cli/player"
	"github.com/sirupsen/logrus"
)

const (
	ROOKIE = iota + 1
	CHAMPION
	ULTIMATE
	MEGA
)

type Skill struct {
	H            *player.Hero
	result       uint8
	RoomIDList   []uint64
	SkillType    uint8
	SkillLevel   uint8
	SkillTargets []uint64
	POWERUP      uint8
	DEFENCE      uint8
	ESCAPE       uint8
	QUIT         uint8
	ATTACK       uint8
	EVOLVE       uint8
}

func (s *Skill) Show() {
	if s.H.SkillPoint < 0 {
		return
	}
	switch s.H.SkillPoint {
	case 0:
		fmt.Println("please choose Power Up or Defence or Escape or Quit, input number only")
		fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit")
		fmt.Scan(&s.result)
	case 1:
		fmt.Println("please choose Power Up or Defence or Escape or Quit or Attack, input number only")
		fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit		5.Attack")
		fmt.Scan(&s.result)
	default:
		if s.H.IdentityLevel == CHAMPION && s.H.SkillPoint < 3 {
			fmt.Println("please choose Power Up or Defence or Escape or Quit or Attack, input number only")
			fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit		5.Attack")
			fmt.Scan(&s.result)
		} else if s.H.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON || s.H.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON {
			fmt.Println("please choose Power Up or Defence or Escape or Quit or Attack, input number only")
			fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit		5.Attack")
			fmt.Scan(&s.result)
		} else {
			fmt.Println("please choose Power Up or Defence or Escape or Quit or Attack or Evolve, input number only")
			fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit		5.Attack		6.Evolve")
			fmt.Scan(&s.result)
		}
	}
}

func (s *Skill) ShowLevel() {
	switch s.SkillType {
	case s.ATTACK:
		if s.H.Identity == pbprotocol.DigimonIdentity_TOGEMON || s.H.Identity == pbprotocol.DigimonIdentity_ANGEMON || s.H.Identity == pbprotocol.DigimonIdentity_LILLYMON || s.H.Identity == pbprotocol.DigimonIdentity_ZUDOMON || s.H.Identity == pbprotocol.DigimonIdentity_GARUDAMON || s.H.Identity == pbprotocol.DigimonIdentity_MAGNAANGEMON || s.H.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON || s.H.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_ANGEWOMON || s.H.Identity == pbprotocol.DigimonIdentity_WARGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_METALGARURUMON {
			if s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_WARGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_METALGARURUMON) {
				fmt.Println("please choose attack level, input number only")
				fmt.Println("1.one		2.two		3.five")
				fmt.Scan(&s.result)
				if s.result == 3 {
					s.result = s.result + 2
				}
			} else {
				fmt.Println("please choose attack level, input number only")
				fmt.Println("1.one		2.two")
				fmt.Scan(&s.result)
			}
		} else {
			s.result = 1
		}
	case s.DEFENCE:
		fmt.Println("please choose defence level, input number only")
		fmt.Println("1.one		2.two		3.three		4.four")
		fmt.Scan(&s.result)
	case s.EVOLVE:
		if s.H.IdentityLevel == ROOKIE {
			if s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_AGUMON || s.H.Identity == pbprotocol.DigimonIdentity_GABUMON) {
				fmt.Println("please choose evolve type, input number only")
				fmt.Println("1.evolve		2.mega-evolve")
				fmt.Scan(&s.result)
				if s.result == 2 {
					s.result++
				}
			} else {
				fmt.Println("please choose evolve type, input number only")
				fmt.Println("1.evolve")
				fmt.Scan(&s.result)
			}
		} else if s.H.IdentityLevel == CHAMPION && s.H.SkillPoint >= 3 {
			if s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_GREYMON || s.H.Identity == pbprotocol.DigimonIdentity_GARURUMON) {
				fmt.Println("please choose evolve type, input number only")
				fmt.Println("1.super-evolve		2.mega-evolve")
				fmt.Scan(&s.result)
			} else {
				fmt.Println("please choose evolve type, input number only")
				fmt.Println("1.super-evolve")
				fmt.Scan(&s.result)
			}
			s.result++
		} else if s.H.IdentityLevel == ULTIMATE && s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON) {
			fmt.Println("please choose evolve type, input number only")
			fmt.Println("1.mega-evolve")
			fmt.Scan(&s.result)
			s.result += 2
		}
	}
}

func (s *Skill) ShowAttackTarget() {
	//TODO: show
	s.result = 0
}

func (s *Skill) Result() uint8 {
	return s.result
}

func NewSkill() *Skill {
	return &Skill{
		H:            new(player.Hero),
		result:       0,
		RoomIDList:   make([]uint64, 0),
		SkillType:    0,
		SkillLevel:   0,
		SkillTargets: make([]uint64, 0),
		POWERUP:      1,
		DEFENCE:      2,
		ESCAPE:       3,
		QUIT:         4,
		ATTACK:       5,
		EVOLVE:       6,
	}
}

func (s *Skill) Refresh() {
	s.result = 0
}

func (s *Skill) RefreshType() {
	s.SkillType = 0
}

func (s *Skill) RefreshLevel() {
	s.SkillLevel = 0
}

func (s *Skill) RefreshSkillTargets() {
	s.SkillTargets = nil
	s.SkillTargets = make([]uint64, 0)
}

func (s *Skill) IsSkillTypValid() bool {
	switch s.H.SkillPoint {
	case 0:
		if s.SkillType == 1 || s.SkillType == 2 || s.SkillType == 3 || s.SkillType == 4 {
			return true
		} else {
			return false
		}
	case 1:
		if s.SkillType == 1 || s.SkillType == 2 || s.SkillType == 3 || s.SkillType == 4 || s.SkillType == 5 {
			return true
		} else {
			return false
		}
	default:
		if s.H.IdentityLevel == CHAMPION && s.H.SkillPoint < 3 {
			if s.SkillType == 1 || s.SkillType == 2 || s.SkillType == 3 || s.SkillType == 4 || s.SkillType == 5 {
				return true
			} else {
				return false
			}
		} else if s.H.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON || s.H.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON {
			if s.SkillType == 1 || s.SkillType == 2 || s.SkillType == 3 || s.SkillType == 4 || s.SkillType == 5 {
				return true
			} else {
				return false
			}
		} else {
			if s.SkillType == 1 || s.SkillType == 2 || s.SkillType == 3 || s.SkillType == 4 || s.SkillType == 5 || s.SkillType == 6 {
				return true
			} else {
				return false
			}
		}
	}
}

func (s *Skill) IsSkillLevelValid() bool {
	switch s.SkillType {
	case s.ATTACK:
		if s.H.Identity == pbprotocol.DigimonIdentity_TOGEMON || s.H.Identity == pbprotocol.DigimonIdentity_ANGEMON || s.H.Identity == pbprotocol.DigimonIdentity_LILLYMON || s.H.Identity == pbprotocol.DigimonIdentity_ZUDOMON || s.H.Identity == pbprotocol.DigimonIdentity_GARUDAMON || s.H.Identity == pbprotocol.DigimonIdentity_MAGNAANGEMON || s.H.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON || s.H.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_ANGEWOMON || s.H.Identity == pbprotocol.DigimonIdentity_WARGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_METALGARURUMON {
			if s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_WARGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_METALGARURUMON) {
				if s.SkillLevel == 1 || s.SkillLevel == 2 || s.SkillLevel == 5 {
					return true
				} else {
					return false
				}
			} else {
				if s.SkillLevel == 1 || s.SkillLevel == 2 {
					return true
				} else {
					return false
				}
			}
		} else {
			if s.SkillLevel == 1 {
				return true
			} else {
				return false
			}
		}
	case s.DEFENCE:
		if s.SkillLevel == 1 || s.SkillLevel == 2 || s.SkillLevel == 3 || s.SkillLevel == 4 {
			return true
		} else {
			return false
		}
	case s.EVOLVE:
		if s.H.IdentityLevel == ROOKIE {
			if s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_AGUMON || s.H.Identity == pbprotocol.DigimonIdentity_GABUMON) {
				if s.SkillLevel == 1 || s.SkillLevel == 3 {
					return true
				} else {
					return false
				}
			} else {
				if s.SkillLevel == 1 {
					return true
				} else {
					return false
				}
			}
		} else if s.H.IdentityLevel == CHAMPION && s.H.SkillPoint >= 3 {
			if s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_GREYMON || s.H.Identity == pbprotocol.DigimonIdentity_GARURUMON) {
				if s.SkillLevel == 2 || s.SkillLevel == 3 {
					return true
				} else {
					return false
				}
			} else {
				if s.SkillLevel == 2 {
					return true
				} else {
					return false
				}
			}
		} else if s.H.IdentityLevel == ULTIMATE && s.H.SkillPoint >= 5 && (s.H.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || s.H.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON) {
			if s.SkillLevel == 3 {
				return true
			} else {
				return false
			}
		} else {
			logrus.Errorln("attribute invalid")
			return false
		}
	default:
		logrus.Errorln("attribute invalid")
		return false
	}
}

func (s *Skill) IsLevelType() bool {
	switch s.SkillType {
	case s.ATTACK:
		return true
	case s.DEFENCE:
		return true
	case s.EVOLVE:
		return true
	default:
		return false
	}
}

func (s *Skill) SetHero(hero *player.Hero) {
	s.H = hero
}

func (s *Skill) IsSkillTargetValid() bool {
	if s.SkillTargets[0] != s.RoomIDList[0] {
		return false
	}
	return true
}
