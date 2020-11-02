package v2

import (
	"bufio"
	"errors"
	"github.com/wujunwei/uni-redis/pkg/redis/protocol"
	"strconv"
)

type Reply struct {
	value interface{}
	types byte
}

func (r Reply) parseInt() int {
	return r.value.(int)
}
func (r Reply) Value() interface{} {
	return r.value
}
func (r Reply) IsError() bool {
	return r.types == Error
}

var (
	UnKnownType = errors.New("unknown type")
)

type RespDecoder struct {
}

func (r *RespDecoder) Decode(reader *bufio.Reader) (reply protocol.Reply, err error) {
	return r.decode(reader)
}

func (r *RespDecoder) decode(reader *bufio.Reader) (reply *Reply, err error) {
	respType, err := reader.ReadByte()
	if err != nil {
		return
	}
	switch respType {
	case Array:
		reply, err = r.decodeArray(reader)
	case Integer:
		reply, err = r.decodeInt(reader)
	case SimpleString:
		reply, err = r.decodeSimpleString(reader)
	case BulkString:
		reply, err = r.decodeBulkString(reader)
	case Error:
		reply, err = r.decodeError(reader)
	default:
		err = UnKnownType
	}
	return
}
func (RespDecoder) readLine(reader *bufio.Reader) (string, error) {
	temp, err := reader.ReadString(CR)
	if err != nil {
		return "", err
	}
	_, _ = reader.Discard(1)
	return temp, nil
}
func (r *RespDecoder) decodeArray(reader *bufio.Reader) (reply *Reply, err error) {
	reply = &Reply{types: Array}
	temp, err := r.readLine(reader)
	arrLen, _ := strconv.Atoi(temp)
	var val = make([]*Reply, arrLen)
	for i := 0; i < arrLen; i++ {
		val[i], err = r.decode(reader)
		if err != nil {
			return
		}
	}
	return
}
func (r *RespDecoder) decodeInt(reader *bufio.Reader) (reply *Reply, err error) {
	reply = &Reply{types: Integer}
	integer, err := r.readLine(reader)
	i, _ := strconv.Atoi(integer)
	reply.value = i
	return
}
func (r *RespDecoder) decodeBulkString(reader *bufio.Reader) (reply *Reply, err error) {
	reply = &Reply{
		types: BulkString,
	}
	integer, err := r.readLine(reader)
	i, _ := strconv.Atoi(integer)
	if i == -1 {
		reply.value = nil
		return
	}
	val := make([]byte, i)
	_, err = reader.Read(val)
	if err != nil {
		return
	}
	reply.value = val
	_, _ = reader.Discard(2)
	return
}
func (r *RespDecoder) decodeError(reader *bufio.Reader) (reply *Reply, err error) {
	reply = &Reply{
		types: Error,
	}
	e, err := r.readLine(reader)
	if err != nil {
		return
	}
	reply.value = e
	return
}
func (r *RespDecoder) decodeSimpleString(reader *bufio.Reader) (reply *Reply, err error) {
	reply = &Reply{
		types: SimpleString,
	}
	e, err := r.readLine(reader)
	if err != nil {
		return
	}
	reply.value = e
	return
}
