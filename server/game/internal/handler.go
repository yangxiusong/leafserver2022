package internal

import (
	"fmt"
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	hander(&msg.Hello{}, handlerHello)
	hander(&msg.Add{}, handlerAdd)
}

func hander(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("[%s] Hello %v", a.RemoteAddr().String(), m.Name)
	a.WriteMsg(&msg.Hello{
		Name: "servertoClient==>handlerHello",
	})
}

func handlerAdd(args []interface{}) {
	m := args[0].(*msg.Add)
	a := args[1].(gate.Agent)

	log.Debug("[%s] Add %d + %d = %d", a.RemoteAddr(), m.A, m.B, m.A+m.B)
	a.WriteMsg(&msg.Hello{
		Name: fmt.Sprintf("handlerAdd %d + %d = %d", m.A, m.B, m.A+m.B),
	})
}
