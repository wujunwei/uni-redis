package v2

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/wujunwei/uni-redis/pkg/redis/protocol"
	"strconv"
)

//Reply : a safe redis response transfer,if the value doesn't support the kind which you choose ,it will offer a default-value
type Reply struct {
	isNIl bool
	value interface{}
	types byte
}

func (r Reply) ParseInt() int {
	if r.IsInteger() {
		return r.value.(int)
	}
	if r.IsArray() {
		return len(r.value.([]*Reply))
	}
	return 0
}

func (r Reply) ParseString() string {
	if r.IsString() || r.IsError() {
		return r.value.(string)
	}
	if r.IsBytes() {
		return string(r.value.([]byte))
	}
	if r.IsInteger() {
		return strconv.Itoa(r.value.(int))
	}
	if r.IsArray() {
		return "array"
	}
	return "nil"
}

func (r Reply) ParseBytes() []byte {
	if r.IsString() || r.IsError() {
		return []byte(r.value.(string))
	}
	if r.IsBytes() {
		return r.value.([]byte)
	}
	if r.IsInteger() {
		return []byte(strconv.Itoa(r.value.(int)))
	}
	return nil
}

func (r Reply) ParseArray() []*Reply {
	if !r.IsArray() {
		return nil
	}
	return r.value.([]*Reply)
}
func (r Reply) Value() interface{} {
	return r.value
}
func (r Reply) IsInteger() bool {
	return r.types == Integer
}
func (r Reply) IsString() bool {
	return r.types == SimpleString
}
func (r Reply) IsBytes() bool {
	return r.types == BulkString
}
func (r Reply) IsArray() bool {
	return r.types == Array
}
func (r Reply) IsError() bool {
	return r.types == Error
}
func (r Reply) IsNil() bool {
	return r.isNIl
}
func (r Reply) GetType() byte {
	return r.types
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
	fmt.Println()
	return
}
func (RespDecoder) readLine(reader *bufio.Reader) (string, error) {
	temp, err := reader.ReadBytes(CR)
	if err != nil {
		return "", err
	}
	_, _ = reader.Discard(1)
	return string(temp[:len(temp)-1]), nil
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
	reply.value = val
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
	if integer == Null {
		reply.isNIl = true
		reply.value = nil
		return
	}
	i, _ := strconv.Atoi(integer)
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
