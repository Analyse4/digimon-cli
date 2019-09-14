package peer

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
)

//TODO: redesign buffer
type Connection struct {
	wg     *sync.WaitGroup
	conn   *websocket.Conn
	done   chan struct{}
	buffer chan []byte
}

func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		wg:     new(sync.WaitGroup),
		conn:   conn,
		done:   make(chan struct{}),
		buffer: make(chan []byte, 100),
	}
}

func (c *Connection) readLoop() {
	defer c.wg.Done()
	fmt.Println("read loop start")
	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("lost connection with server")
			log.WithFields(logrus.Fields{
				"endpoint": c.conn.RemoteAddr().String(),
			}).Debug(err)

			c.done <- struct{}{}
			c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseTryAgainLater, ""))
			return
		} else {
			processMsg(data)
		}
	}
}

func (c *Connection) writeLoop() {
	defer c.wg.Done()
	fmt.Println("write loop start")
	for {
		select {
		case <-c.done:
			return
		case data := <-c.buffer:
			err := c.conn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				//TODO: process err
			}
		}
	}
}
