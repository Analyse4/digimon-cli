package codec

import (
	"digimon-cli/handler"
	"digimon-cli/pbprotocol"
	"github.com/Analyse4/digimon/logger"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"reflect"
)

type protobufCodec struct{}

var (
	log *logrus.Entry
)

func init() {
	log = logger.GetLogger().WithField("pkg", "codec")
}

func NewProtobufCodec() *protobufCodec {
	return new(protobufCodec)
}

func (pbc *protobufCodec) Marshal(router string, data interface{}) ([]byte, error) {
	d, ok := data.(proto.Message)
	if !ok {
		log.WithFields(logrus.Fields{
			"type": reflect.TypeOf(data),
		}).Warn("invalid type: type need implement interface proto.Message")
	}
	md, err := proto.Marshal(d)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	bm := new(pbprotocol.MsgPack)
	bm.Router = router
	bm.Data = md
	msg, err := proto.Marshal(bm)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return msg, nil
}

func (pbc *protobufCodec) UnMarshal(msg []byte) (string, interface{}, error) {
	bm := new(pbprotocol.MsgPack)
	err := proto.Unmarshal(msg, bm)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}
	router := bm.Router
	mdata := bm.Data
	h, err := handler.GetDigimonCli().GetHandler(router)
	if err != nil {
		return router, nil, err
	}
	data := reflect.New(h.AckType).Interface()
	proto.Unmarshal(mdata, data.(proto.Message))
	return router, data, nil
}
