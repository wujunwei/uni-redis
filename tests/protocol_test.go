package tests

import (
	"bufio"
	"fmt"
	"github.com/wujunwei/uni-redis/pkg/redis"
	v2 "github.com/wujunwei/uni-redis/pkg/redis/protocol/v2"
	"strings"
	"testing"
)

func Test_encoder(t *testing.T) {
	a := redis.NewAnalyzer(nil, nil)
	bf, err := a.En.Encode([]interface{}{"123", 123, []int{1, 2, 3}})
	if err != nil {
		t.Fatal(err)
	}
	if bf.String() != "*3\r\n$3\r\n123\r\n:123\r\n*3\r\n:1\r\n:2\r\n:3\r\n" {
		t.Fatal("is not same")
	}

}

func Test_decoder(t *testing.T) {
	b := redis.NewAnalyzer(nil, nil)
	rep, err := b.De.Decode(bufio.NewReader(strings.NewReader("*3\r\n$3\r\n123\r\n:123\r\n*3\r\n:1\r\n:2\r\n:3\r\n")))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rep.Value().([]*v2.Reply)[0])
}
