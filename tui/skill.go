package tui

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/player"
	"fmt"
)

const (
	ROOKIE = iota + 7
	CHAMPION
	ULTIMATE
	MEGA
)

type Skill struct {
	hero       *player.Hero
	result     uint8
	SkillType  uint8
	SkillLevel uint8
	POWERUP    uint8
	DEFENCE    uint8
	ESCAPE     uint8
	QUIT       uint8
	ATTACK     uint8
	EVOLVE     uint8
}

func (s *Skill) Show() {
	if s.hero.SkillPoint < 0 {
		return
	}
	switch s.hero.SkillPoint {
	case 0:
		fmt.Println("please chose Power Up or Defence or Escape or Quit, input number only")
		fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit")
		fmt.Scan(&s.result)
	case 1:
		fmt.Println("please chose Power Up or Defence or Escape or Quit or Attack, input number only")
		fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit		5.Attack")
		fmt.Scan(&s.result)
	default:
		fmt.Println("please chose Power Up or Defence or Escape or Quit or Attack or Evolve, input number only")
		fmt.Println("1.Power Up		2.Defence		3.Escape		4.Quit		5.Attack		6.Evolve")
		fmt.Scan(&s.result)
	}
}

func (s *Skill) ShowLevel(h *player.Hero) {
	switch s.SkillType {
	case s.ATTACK:
		if h.Identity == pbprotocol.DigimonIdentity_TOGEMON || h.Identity == pbprotocol.DigimonIdentity_ANGEMON || h.Identity == pbprotocol.DigimonIdentity_LILLYMON || h.Identity == pbprotocol.DigimonIdentity_ZUDOMON || h.Identity == pbprotocol.DigimonIdentity_GARUDAMON || h.Identity == pbprotocol.DigimonIdentity_MAGNAANGEMON || h.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON || h.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || h.Identity == pbprotocol.DigimonIdentity_ANGEWOMON || h.Identity == pbprotocol.DigimonIdentity_WARGREYMON || h.Identity == pbprotocol.DigimonIdentity_METALGARURUMON {
			if h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_WARGREYMON || h.Identity == pbprotocol.DigimonIdentity_METALGARURUMON) {
				fmt.Println("please chose attack level, input number only")
				fmt.Println("1.one		2.two		3.five")
				fmt.Scan(&s.result)
				if s.result == 3 {
					s.result = s.result + 2
				}
			} else {
				fmt.Println("please chose attack level, input number only")
				fmt.Println("1.one		2.two")
				fmt.Scan(&s.result)
			}
		}
	case s.DEFENCE:
		fmt.Println("please chose defence level, input number only")
		fmt.Println("1.one		2.two		3.three		4.four")
		fmt.Scan(&s.result)
	case s.EVOLVE:
		if h.IdentityLevel == ROOKIE {
			if h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_AGUMON || h.Identity == pbprotocol.DigimonIdentity_GABUMON) {
				fmt.Println("please chose evolve type, input number only")
				fmt.Println("1.evolve		2.mega-evolve")
				fmt.Scan(&s.result)
				if s.result == 2 {
					s.result++
				}
			} else {
				fmt.Println("please chose evolve type, input number only")
				fmt.Println("1.evolve")
				fmt.Scan(&s.result)
			}
		} else if h.IdentityLevel == CHAMPION && h.SkillPoint >= 3 {
			if h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_GREYMON || h.Identity == pbprotocol.DigimonIdentity_GARURUMON) {
				fmt.Println("please chose evolve type, input number only")
				fmt.Println("1.super-evolve		2.mega-evolve")
				fmt.Scan(&s.result)
			} else {
				fmt.Println("please chose evolve type, input number only")
				fmt.Println("1.super-evolve")
				fmt.Scan(&s.result)
			}
			s.result++
		} else if h.IdentityLevel == ULTIMATE && h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || h.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON) {
			fmt.Println("please chose evolve type, input number only")
			fmt.Println("1.mega-evolve")
			fmt.Scan(&s.result)
			s.result += 2
		}
	}
}

func (s *Skill) Result() uint8 {
	return s.result
}

func NewSkill() *Skill {
	return &Skill{
		hero:       new(player.Hero),
		result:     0,
		SkillType:  0,
		SkillLevel: 0,
		POWERUP:    1,
		DEFENCE:    2,
		ESCAPE:     3,
		QUIT:       4,
		ATTACK:     5,
		EVOLVE:     6,
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

func (s *Skill) IsSkillTypValid() bool {
	switch s.SkillType {
	case s.POWERUP:
		return true
	case s.DEFENCE:
		return true
	case s.ESCAPE:
		return true
	case s.QUIT:
		return true
	case s.ATTACK:
		return true
	case s.EVOLVE:
		return true
	default:
		return false
	}
}

func (s *Skill) IsSkillLevelValid(h *player.Hero) bool {
	switch s.SkillType {
	case s.ATTACK:
		if h.Identity == pbprotocol.DigimonIdentity_TOGEMON || h.Identity == pbprotocol.DigimonIdentity_ANGEMON || h.Identity == pbprotocol.DigimonIdentity_LILLYMON || h.Identity == pbprotocol.DigimonIdentity_ZUDOMON || h.Identity == pbprotocol.DigimonIdentity_GARUDAMON || h.Identity == pbprotocol.DigimonIdentity_MAGNAANGEMON || h.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON || h.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || h.Identity == pbprotocol.DigimonIdentity_ANGEWOMON || h.Identity == pbprotocol.DigimonIdentity_WARGREYMON || h.Identity == pbprotocol.DigimonIdentity_METALGARURUMON {
			if h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_WARGREYMON || h.Identity == pbprotocol.DigimonIdentity_METALGARURUMON) {
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
			return false
		}
	case s.DEFENCE:
		if s.SkillLevel == 1 || s.SkillLevel == 2 || s.SkillLevel == 3 || s.SkillLevel == 4 {
			return true
		} else {
			return false
		}
	case s.EVOLVE:
		if h.IdentityLevel == ROOKIE {
			if h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_AGUMON || h.Identity == pbprotocol.DigimonIdentity_GABUMON) {
				if s.SkillLevel == 1 || s.SkillLevel == 2 {
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
		} else if h.IdentityLevel == CHAMPION && h.SkillPoint >= 3 {
			if h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_GREYMON || h.Identity == pbprotocol.DigimonIdentity_GARURUMON) {
				if s.SkillLevel == 1 || s.SkillLevel == 2 {
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
		} else if h.IdentityLevel == ULTIMATE && h.SkillPoint >= 5 && (h.Identity == pbprotocol.DigimonIdentity_SKULLGREYMON || h.Identity == pbprotocol.DigimonIdentity_WEREGARURUMON) {
			if s.SkillLevel == 1 {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	default:
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
