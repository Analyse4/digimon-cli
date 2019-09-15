package codec

import (
	"digimon-cli/pbprotocol"
	"github.com/Analyse4/digimon/logger"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"reflect"
)

type ProtobufCodec struct{}

var (
	log *logrus.Entry
)

func init() {
	log = logger.GetLogger().WithField("pkg", "codec")
}

func NewProtobufCodec() *ProtobufCodec {
	return new(ProtobufCodec)
}

func (pbc *ProtobufCodec) Marshal(router string, data interface{}) []byte {
	d, ok := data.(proto.Message)
	if !ok {
		log.WithFields(logrus.Fields{
			"type": reflect.TypeOf(data),
		}).Warn("invalid type: type need implement interface proto.Message")
	}
	md, err := proto.Marshal(d)
	if err != nil {
		log.Warn(err)
	}
	bm := new(pbprotocol.MsgPack)
	bm.Router = router
	bm.Data = md
	msg, err := proto.Marshal(bm)
	if err != nil {
		log.Warn(err)
	}
	return msg
}

func (pbc *ProtobufCodec) UnMarshal([]byte) interface{} {
	return nil
}
