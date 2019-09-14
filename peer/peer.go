package peer

import (
	"fmt"
	"github.com/Analyse4/digimon/logger"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/url"
)

var (
	log *logrus.Entry
)

func init() {
	log = logger.GetLogger().WithField("pkg", "peer")
}

func ConnectGameServer(ep string) {
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: ep, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	c := NewConnection(conn)

	go c.readLoop()
	go c.writeLoop()
	go sentinelLoop(c)
}

func sentinelLoop(c *Connection) {
	fmt.Println("sentinel loop start")
	c.wg.Wait()
	fmt.Println("all loop finished")

	c.conn.Close()
	//TODO: prompt reconnect or quit
}

func processMsg(data []byte) {
	fmt.Println("process msg")
}
