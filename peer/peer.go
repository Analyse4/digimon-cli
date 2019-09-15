package peer

import (
	"github.com/Analyse4/digimon/logger"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry
)

func init() {
	log = logger.GetLogger().WithField("pkg", "peer")
}

type Connection interface {
	Connect(string)
	Send(string, interface{})
	ReadLoop()
	WriteLoop()
	SentinelLoop()
	Refresh()
}

//func ConnectGameServer(ep string) *Connection{
//	//interrupt := make(chan os.Signal, 1)
//	//signal.Notify(interrupt, os.Interrupt)
//
//	u := url.URL{Scheme: "ws", Host: ep, Path: "/echo"}
//	log.Debugf("connecting to %s", u.String())
//
//	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		log.Fatal("dial:", err)
//	}
//
//	c := NewConnection(conn)
//
//	go c.readLoop()
//	go c.writeLoop()
//	go sentinelLoop(c)
//
//	log.WithFields(logrus.Fields{
//		"endpoint": c.conn.RemoteAddr().String(),
//	}).Debug("game server connect successful")
//	fmt.Println("Welcome to Digimon World")
//	return c
//}
//
//func sentinelLoop(c *Connection) {
//	ep := c.conn.RemoteAddr().String()
//	log.Debug("sentinel loop start")
//	c.wg.Wait()
//	log.Debug("all loop finished")
//
//	c.conn.Close()
//	reconnectTUI := new(tui.ReconnectTUI)
//	handler.Reconnect(ep, reconnectTUI)
//}
//
//func processMsg(data []byte) {
//	fmt.Println("process msg")
//}
