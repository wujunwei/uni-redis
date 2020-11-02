package v2

import (
	"bytes"
	"strconv"
)

type RespEncoder struct {
}

func (e RespEncoder) Encode(param []interface{}) (bytes.Buffer, error) {
	buf := bytes.Buffer{}
	err := e.encodeArray(param, buf)
	return buf, err
}

func (e RespEncoder) encodeArray(param []interface{}, buf bytes.Buffer) (err error) {
	buf.WriteByte(Array)
	buf.WriteString(strconv.Itoa(len(param)))
	buf.WriteByte(CR)
	buf.WriteByte(LF)
	for _, p := range param {
		switch p.(type) {
		case int:
			e.encodeIntegers(p.(int), buf)
		case string:
			e.encodeString(p.(string), buf)
		case []int:
			arr := p.([]int)
			buf.WriteByte(Array)
			buf.WriteString(strconv.Itoa(len(arr)))
			buf.WriteByte(CR)
			buf.WriteByte(LF)
			for _, i := range arr {
				e.encodeIntegers(i, buf)
			}
		case []string:
			arr := p.([]string)
			buf.WriteByte(Array)
			buf.WriteString(strconv.Itoa(len(arr)))
			buf.WriteByte(CR)
			buf.WriteByte(LF)
			for _, s := range arr {
				e.encodeString(s, buf)
			}
		default:
			err = UnKnownType
		}
	}
	return
}

func (e RespEncoder) encodeString(s string, buf bytes.Buffer) {
	buf.WriteByte(BulkString)
	buf.WriteString(strconv.Itoa(len(s)))
	buf.WriteByte(CR)
	buf.WriteByte(LF)
	buf.WriteString(s)
	buf.WriteByte(CR)
	buf.WriteByte(LF)
}

func (e RespEncoder) encodeIntegers(i int, buf bytes.Buffer) {
	buf.WriteByte(Integer)
	buf.WriteString(strconv.Itoa(i))
	buf.WriteByte(CR)
	buf.WriteByte(LF)
}
