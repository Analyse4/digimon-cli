package codec

type Codec interface {
	Marshal(string, interface{}) []byte
	UnMarshal([]byte) interface{}
}
