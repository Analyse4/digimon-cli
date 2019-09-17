package handler

import (
	"digimon-cli/peer"
	"digimon/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
	"sync"
)

var (
	log *logrus.Entry
	dc  *digimonCli
)

type handler struct {
	AckType reflect.Type
	Method  reflect.Method
}

type digimonCli struct {
	conn            peer.Connection
	handlerRegister sync.Map
	playerID        uint16
}

func init() {
	log = logger.GetLogger().WithField("pkg", "handler")
}

func New() *digimonCli {
	dc = &digimonCli{
		conn:            nil,
		handlerRegister: sync.Map{},
	}
	return dc
}

func (dc *digimonCli) Register() {
	typ := reflect.TypeOf(dc)
	for i := 0; i < typ.NumMethod(); i++ {
		m := typ.Method(i)
		if isProperMethod(m) {
			s := m.Name
			router := "digimon." + strings.ToLower(s[:len(s)-3])
			h := new(handler)
			h.AckType = m.Type.In(1).Elem()
			h.Method = m
			dc.handlerRegister.Store(router, h)
		}
	}
}

func (dc *digimonCli) Show() {
	dc.handlerRegister.Range(func(key, value interface{}) bool {
		fmt.Println("---------------")
		fmt.Println(value.(*handler).AckType)
		fmt.Println(value.(*handler).Method.Func)
		return true
	})
}

func (dc *digimonCli) GetHandler(router string) (*handler, error) {
	h, ok := dc.handlerRegister.Load(router)
	if !ok {
		return nil, fmt.Errorf("handler not found")
	}
	return h.(*handler), nil
}

func (dc *digimonCli) SetConn(conn peer.Connection) {
	dc.conn = conn
}

func GetDigimonCli() *digimonCli {
	return dc
}

func isProperMethod(m reflect.Method) bool {
	if m.Type.NumIn() != 2 {
		return false
	}
	if m.Type.NumOut() != 1 {
		return false
	}
	if m.Type.In(0).Kind() != reflect.Ptr {
		return false
	}
	return true
}
