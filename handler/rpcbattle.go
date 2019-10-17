package handler

import (
	"fmt"
	"github.com/Analyse4/digimon-cli/pbprotocol"
	"github.com/Analyse4/digimon-cli/peer"
	"github.com/Analyse4/digimon-cli/tui"
	"github.com/sirupsen/logrus"
)

func RPCBattleReq(conn peer.Connection, RPCBattleTUI *tui.RPCBattle) {
	RPCBattleTUI.Show()
	if RPCBattleTUI.Result() != RPCBattleTUI.Rock && RPCBattleTUI.Result() != RPCBattleTUI.Paper && RPCBattleTUI.Result() != RPCBattleTUI.Scissors {
		fmt.Println("invalid input, please reselect")
		RPCBattleTUI.Refresh()
		RPCBattleReq(conn, RPCBattleTUI)
	} else {
		req := new(pbprotocol.RPCBattleReq)
		req.Role = RPCBattleTUI.Identity
		req.OtherSideId = RPCBattleTUI.OtherSideID
		req.Rpc = RPCBattleTUI.Result()
		conn.Send("digimon.rpcbattle", req)
	}
}

func (dc *digimonCli) RPCBattleAck(ack *pbprotocol.RPCBattleAck) error {
	if ack.Base.Result != 0 {
		logrus.Error(ack.Base.Msg)
		return fmt.Errorf(ack.Base.Msg)
	}
	if ack.LastWinId == 0 {
		fmt.Println("The Last Battle is Draw")
	} else if ack.LastWinId == dc.player.ID {
		fmt.Println("Congratulations, the Winner of Last Battle is You")
	} else {
		fmt.Println("You lose last battle")
	}

	if ack.IsHaveNext {
		var rpcBattleTUI *tui.RPCBattle
		switch dc.player.ID {
		case ack.AttackerId:
			rpcBattleTUI = tui.NewRPCBattle(ATTACKER, ack.TargetId)
		case ack.TargetId:
			rpcBattleTUI = tui.NewRPCBattle(TARGET, ack.AttackerId)
		default:
			return fmt.Errorf("invalid patameter")
		}
		RPCBattleReq(dc.conn, rpcBattleTUI)
	}
	return nil
}
