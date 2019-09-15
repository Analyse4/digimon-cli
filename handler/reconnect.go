package handler

//func Reconnect(ep string, reconnectTUI *tui.ReconnectTUI) {
//	reconnectTUI.Show()
//	switch reconnectTUI.Result() {
//	case reconnectTUI.RECONNECT:
//		c := wsconnection.New()
//		c.Connect(ep)
//		loginTUI := new(tui.Login)
//		LoginReq(c, loginTUI)
//		//TODO: main logic need write here
//	case reconnectTUI.QUIT:
//		os.Exit(0)
//	default:
//		fmt.Println("invalid input, please reselect")
//		reconnectTUI.Refresh()
//		Reconnect(ep, reconnectTUI)
//	}
//}
