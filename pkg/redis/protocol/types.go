package protocol

import (
	"bufio"
	"bytes"
)

type Reply interface {
	Value() interface{}
}

type Encoder interface {
	Encode([]interface{}) (*bytes.Buffer, error)
}
type Decode interface {
	Decode(reader *bufio.Reader) (Reply, error)
}
