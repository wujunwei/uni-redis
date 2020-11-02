package redis

import (
	"github.com/wujunwei/uni-redis/pkg/redis/protocol"
	v2 "github.com/wujunwei/uni-redis/pkg/redis/protocol/v2"
)

var defaultEncoder = &v2.RespEncoder{}
var defaultDecoder = &v2.RespDecoder{}

type Analyzer struct {
	En protocol.Encoder
	De protocol.Decode
}

func NewAnalyzer(en protocol.Encoder, de protocol.Decode) *Analyzer {
	if en == nil || de == nil {
		return &Analyzer{En: defaultEncoder, De: defaultDecoder}
	}
	return &Analyzer{En: en, De: de}
}

func init() {

}
