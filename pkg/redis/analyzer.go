package redis

import (
	"github.com/wujunwei/uni-redis/pkg/redis/protocol"
	v2 "github.com/wujunwei/uni-redis/pkg/redis/protocol/v2"
)

var defaultEncoder = &v2.RespEncoder{}
var defaultDecoder = &v2.RespDecoder{}

type Analyzer struct {
	en protocol.Encoder
	de protocol.Decode
}

func NewAnalyzer(en protocol.Encoder, de protocol.Decode) *Analyzer {
	if en == nil || de == nil {
		return &Analyzer{en: defaultEncoder, de: defaultDecoder}
	}
	return &Analyzer{en: en, de: de}
}

func init() {

}
