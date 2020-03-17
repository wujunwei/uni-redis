package protocol

import (
	"bufio"
	"errors"
	"strconv"
)

const (
	Array        = '*'
	Integer      = ':'
	BulkString   = '$'
	SimpleString = '+'
	Errors       = '-'
	CR           = '\r'
	LF           = '\n'
)

var (
	UnKnownType = errors.New("unknown type")
)

type Reply struct {
	value interface{}
	types byte
}

func (reply Reply) parseInt() int {
	return reply.value.(int)
}
func (reply Reply) Value() interface{} {
	return reply.value
}
func (reply Reply) SetValue(v interface{}) {
	reply.value = v
}
func (reply Reply) SetType(t byte) {
	reply.types = t
}

type RespAnalyzer struct {
}

func (analyzer *RespAnalyzer) Decode(reader *bufio.Reader) (reply *Reply, err error) {
	respType, err := reader.ReadByte()
	if err != nil {
		return

	}
	switch respType {
	case Array:
		{
			return analyzer.decodeArray(reader)
		}

	case Integer:
		{
			return analyzer.decodeInt(reader)
		}

	case SimpleString:
		{
			return analyzer.decodeSimpleString(reader)
		}

	case BulkString:
		{
			return analyzer.decodeBulkString(reader)
		}
	case Errors:
		{
			return analyzer.decodeError(reader)
		}
	default:
		{
			err = UnKnownType
			return
		}

	}
}

func (analyzer *RespAnalyzer) decodeArray(reader *bufio.Reader) (reply *Reply, err error) {
	reply = &Reply{}
	temp, err := reader.ReadString(CR)
	arrLen, _ := strconv.Atoi(temp)
	for i := 0; i < arrLen; i++ {

	}
	return
}
func (analyzer *RespAnalyzer) decodeInt(reader *bufio.Reader) (reply *Reply, err error) {
	return
}
func (analyzer *RespAnalyzer) decodeBulkString(reader *bufio.Reader) (reply *Reply, err error) {
	return
}
func (analyzer *RespAnalyzer) decodeError(reader *bufio.Reader) (reply *Reply, err error) {
	return
}
func (analyzer *RespAnalyzer) decodeSimpleString(reader *bufio.Reader) (reply *Reply, err error) {
	return
}

//func (analyzer *RespAnalyzer) encode(interface{}) bytes.Buffer {
//	return
//}

func init() {

}
