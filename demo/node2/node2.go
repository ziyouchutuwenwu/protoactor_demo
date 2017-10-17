package node2

import (
	"runtime"

	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"protoactor_demo/demo/messages"
	"github.com/AsynkronIT/protoactor-go/remote"
)

type helloActor struct{}

func (*helloActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *messages.HelloRequest:
		ctx.Respond(&messages.HelloResponse{
			Message: "Hello from remote node",
		})
	}
}

func newHelloActor() actor.Actor {
	return &helloActor{}
}

func Run() {
	remote.Register("hello", actor.FromProducer(newHelloActor))

	runtime.GOMAXPROCS(runtime.NumCPU())

	remote.Start("127.0.0.1:8080")

	console.ReadLine()
}