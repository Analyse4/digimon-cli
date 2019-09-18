package player

type Player struct {
	ID        uint64
	NickName  string
	LoginType int32
}

func New() *Player {
	return &Player{
		ID:        0,
		NickName:  "",
		LoginType: -1,
	}
}
