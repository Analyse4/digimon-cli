package wsconnection

import (
	"digimon-cli/codec"
	"digimon-cli/handler"
	"digimon-cli/tui"
	"digimon/logger"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"reflect"
	"sync"
)

var (
	log *logrus.Entry
)

func init() {
	log = logger.GetLogger().WithField("pkg", "wsconnection")
}

//TODO: redesign buffer
type WSConnection struct {
	wg     *sync.WaitGroup
	conn   *websocket.Conn
	done   chan struct{}
	buffer chan []byte
}

func New() *WSConnection {
	c := &WSConnection{
		wg:     new(sync.WaitGroup),
		conn:   nil,
		done:   make(chan struct{}),
		buffer: make(chan []byte, 100),
	}
	c.wg.Add(2)
	return c
}

func (w *WSConnection) Connect(ep string) {
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: ep, Path: "/echo"}
	log.Debugf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	w.conn = conn

	go w.ReadLoop()
	go w.WriteLoop()
	go w.SentinelLoop()

	log.WithFields(logrus.Fields{
		"endpoint": w.conn.RemoteAddr().String(),
	}).Debug("game server connect successful")
	fmt.Println("Welcome to Digimon World")
}

func (w *WSConnection) Send(router string, data interface{}) {
	pbc := codec.NewProtobufCodec()
	// TODO: error handler
	msg, _ := pbc.Marshal(router, data)
	w.buffer <- msg

	log.WithFields(logrus.Fields{
		"router":   router,
		"data_len": len(msg),
	}).Debug("send data")
}

func (w *WSConnection) ReadLoop() {
	defer w.wg.Done()
	log.Debug("read loop start")
	for {
		_, data, err := w.conn.ReadMessage()
		if err != nil {
			log.WithFields(logrus.Fields{
				"endpoint": w.conn.RemoteAddr().String(),
			}).Debug("read loop finished: " + err.Error())

			w.done <- struct{}{}
			w.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseTryAgainLater, ""))
			return
		} else {
			log.WithFields(logrus.Fields{
				"len": len(data),
			}).Debug("receive data")

			processMsg(data)
		}
	}
}

func (w *WSConnection) WriteLoop() {
	defer w.wg.Done()
	log.Debug("write loop start")
	for {
		select {
		case <-w.done:
			log.WithFields(logrus.Fields{
				"endpoint": w.conn.RemoteAddr().String(),
			}).Debug("write loop finished")

			return
		case data := <-w.buffer:
			err := w.conn.WriteMessage(websocket.BinaryMessage, data)
			if err != nil {
				//TODO: process err
				log.Println(err)
			}
		}
	}
}

func (w *WSConnection) SentinelLoop() {
	ep := w.conn.RemoteAddr().String()
	log.Debug("sentinel loop start")
	w.wg.Wait()
	log.Debug("all loop finished")

	w.conn.Close()
	reconnectTUI := tui.NewReconnect()
	reconnect(ep, reconnectTUI)
}

// TODO: check prompt
func (w *WSConnection) Refresh() {
	w = New()
}

func processMsg(data []byte) {
	pbc := codec.NewProtobufCodec()
	router, ack, err := pbc.UnMarshal(data)
	if err != nil {
		log.WithFields(logrus.Fields{
			"router": router,
		}).Error("unmarshal failed: " + err.Error())
	}
	h, _ := handler.GetDigimonCli().GetHandler(router)
	resultSet := h.Method.Func.Call([]reflect.Value{reflect.ValueOf(handler.GetDigimonCli()), reflect.ValueOf(ack)})
	if resultSet[0].Interface() != nil {
		log.WithFields(logrus.Fields{
			"router": router,
		}).Error("handler execute failed: " + resultSet[0].Interface().(error).Error())
	}
}

func reconnect(ep string, reconnectTUI *tui.ReconnectTUI) {
	reconnectTUI.Show()
	switch reconnectTUI.Result() {
	case reconnectTUI.RECONNECT:
		c := New()
		c.Connect(ep)
		loginTUI := tui.NewLogin()
		handler.LoginReq(c, loginTUI)
		//TODO: main logic need write here
	case reconnectTUI.QUIT:
		os.Exit(0)
	default:
		fmt.Println("invalid input, please reselect")
		reconnectTUI.Refresh()
		reconnect(ep, reconnectTUI)
	}
}
