package peer

////TODO: redesign buffer
//type Connection struct {
//	wg     *sync.WaitGroup
//	conn   *websocket.Conn
//	done   chan struct{}
//	buffer chan []byte
//}
//
//func NewConnection(conn *websocket.Conn) *Connection {
//	c :=  &Connection{
//		wg:     new(sync.WaitGroup),
//		conn:   conn,
//		done:   make(chan struct{}),
//		buffer: make(chan []byte, 100),
//	}
//	c.wg.Add(2)
//	return c
//}
//
//func (c *Connection) Send(router string, data interface{}) {
//	pbc := codec.NewProtobufCodec()
//	msg := pbc.Marshal(router, data)
//	c.buffer <- msg
//}
//
//func (c *Connection) readLoop() {
//	defer c.wg.Done()
//	log.Debug("read loop start")
//	for {
//		_, data, err := c.conn.ReadMessage()
//		if err != nil {
//			log.WithFields(logrus.Fields{
//				"endpoint": c.conn.RemoteAddr().String(),
//			}).Debug("read loop finished: " + err.Error())
//
//			c.done <- struct{}{}
//			c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseTryAgainLater, ""))
//			return
//		} else {
//			processMsg(data)
//		}
//	}
//}
//
//func (c *Connection) writeLoop() {
//	defer c.wg.Done()
//	log.Debug("write loop start")
//	for {
//		select {
//		case <-c.done:
//			log.WithFields(logrus.Fields{
//				"endpoint": c.conn.RemoteAddr().String(),
//			}).Debug("write loop finished")
//
//			return
//		case data := <-c.buffer:
//			err := c.conn.WriteMessage(websocket.BinaryMessage, data)
//			if err != nil {
//				//TODO: process err
//				log.Println(err)
//			}
//		}
//	}
//}
