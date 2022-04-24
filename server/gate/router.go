package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Add{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)
}
