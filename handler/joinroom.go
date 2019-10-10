package handler

import (
	"digimon-cli/pbprotocol"
	"digimon-cli/peer"
	"digimon-cli/player"
	"digimon-cli/tui"
	"fmt"
	"os"
)

func JoinRoomReq(conn peer.Connection, JoinRoomTUI *tui.JoinRoom) {
	JoinRoomTUI.Show()
	switch JoinRoomTUI.Result() {
	case JoinRoomTUI.JOINROOM:
		fmt.Println("send join room req")
		req := new(pbprotocol.JoinRoomReq)
		conn.Send("digimon.joinroom", req)
	case JoinRoomTUI.QUIT:
		os.Exit(0)
	default:
		fmt.Println("invalid input, please reselect")
		JoinRoomTUI.Refresh()
		JoinRoomReq(conn, JoinRoomTUI)
	}
	fmt.Println("login successful")
}

func (dc *digimonCli) JoinRoomAck(ack *pbprotocol.JoinRoomAck) error {
	if ack.Base.Result != 0 {
		//TODO: try again later
		log.Println(ack.Base.Msg)
	}
	dc.room.Id = ack.RoomInfo.RoomId
	dc.room.Type = int32(ack.RoomInfo.Type)
	dc.room.CurrentNum = ack.RoomInfo.CurrentPlayerNum
	dc.room.IsStart = ack.RoomInfo.IsStart
	for _, v := range ack.RoomInfo.PlayerInfos {
		tmpPlayer := new(player.Player)
		tmpPlayer.ID = v.Id
		tmpPlayer.NickName = v.Nickname
		tmpPlayer.DigiMonster = new(player.Hero)
		dc.room.PlayerInfos = append(dc.room.PlayerInfos, tmpPlayer)
	}
	dc.ShowRoomInfo()
	return nil
}

func (dc *digimonCli) ShowRoomInfo() {
	fmt.Printf("room_id: %d\n", dc.room.Id)
	fmt.Printf("room_type: %d\n", dc.room.Type)
	fmt.Printf("is_start: %t\n", dc.room.IsStart)
	fmt.Printf("current_num: %d\n", dc.room.CurrentNum)
	for i := 0; i < int(dc.room.CurrentNum); i++ {
		fmt.Printf("player_id: %d, player_nickname: %s\n", dc.room.PlayerInfos[i].ID, dc.room.PlayerInfos[i].NickName)
	}
}
