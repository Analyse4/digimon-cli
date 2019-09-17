package codec

type Codec interface {
	Marshal(string, interface{}) ([]byte, error)
	UnMarshal([]byte) (string, interface{}, error)
}
