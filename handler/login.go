package handler

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/peer"
	"digimon-cli/tui"
	"fmt"
	"github.com/sirupsen/logrus"
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

		log.WithFields(logrus.Fields{
			"router":   "digimon.login",
			"data_len": ack.XXX_Size(),
		}).Debug("send data")
	case loginTUI.QUIT:
		os.Exit(0)
	default:
		fmt.Println("invalid input, please reselect")
		loginTUI.Refresh()
		LoginReq(conn, loginTUI)
	}
	fmt.Println("login successful")
}
