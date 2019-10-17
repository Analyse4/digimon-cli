package handler

import (
	"fmt"
	"github.com/Analyse4/digimon-cli/pbprotocol"
	"github.com/Analyse4/digimon-cli/peer"
	"github.com/Analyse4/digimon-cli/tui"
	"os"
)

func LoginReq(conn peer.Connection, loginTUI *tui.Login) {
	loginTUI.Show()
	switch loginTUI.Result() {
	case loginTUI.LOGIN:
		fmt.Println("send login req")
		ack := new(pbprotocol.LoginReq)
		ack.Type = pbprotocol.LoginReq_Visitor
		conn.Send("digimon.login", ack)
	case loginTUI.QUIT:
		os.Exit(0)
	default:
		fmt.Println("invalid input, please reselect")
		loginTUI.Refresh()
		LoginReq(conn, loginTUI)
	}
	fmt.Println("login successful")
}

func (dc *digimonCli) LoginAck(ack *pbprotocol.LoginAck) error {
	dc.player.ID = ack.PlayerInfo.Id
	dc.player.NickName = ack.PlayerInfo.Nickname
	fmt.Println(dc.player.ID)
	fmt.Println(dc.player.NickName)
	return nil
}
