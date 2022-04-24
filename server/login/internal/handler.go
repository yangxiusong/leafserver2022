package internal

import (
	"reflect"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

	"server/msg"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	hander(&msg.Login{}, handlerLogin)
}

func hander(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerLogin(args []interface{}) {
	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)

	log.Debug("[%s] Login Name:%s Pwd:%s", a.RemoteAddr().String(), m.Name, m.Pwd)
	a.WriteMsg(&msg.Hello{
		Name: "servertoClient:handlerLogin",
	})
}
