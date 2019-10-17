package handler

import (
	"fmt"
	"github.com/Analyse4/digimon-cli/pbprotocol"
	"github.com/Analyse4/digimon-cli/tui"
)

const (
	ATTACKER = iota + 1
	TARGET
)

func (dc *digimonCli) RulingResultAck(ack *pbprotocol.RulingResultAck) error {
	if ack.AttackerID != dc.player.ID && ack.TargetID != dc.player.ID {
		return nil
	}
	var rpcBattleTUI *tui.RPCBattle
	switch dc.player.ID {
	case ack.AttackerID:
		rpcBattleTUI = tui.NewRPCBattle(ATTACKER, ack.TargetID)
	case ack.TargetID:
		rpcBattleTUI = tui.NewRPCBattle(TARGET, ack.AttackerID)
	default:
		return fmt.Errorf("invaild parameter")
	}
	RPCBattleReq(dc.conn, rpcBattleTUI)
	return nil
}
