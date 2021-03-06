package node1

import (
	"fmt"
	"time"
	console "github.com/AsynkronIT/goconsole"
	"protoactor_demo/demo/messages"
	"github.com/AsynkronIT/protoactor-go/remote"
)

func Run() {
	timeout := 5 * time.Second
	remote.Start("127.0.0.1:8081")
	pidResp, _ := remote.SpawnNamed("127.0.0.1:8080", "remote", "hello", timeout)
	pid := pidResp.Pid
	res, _ := pid.RequestFuture(&messages.HelloRequest{}, timeout).Result()
	response := res.(*messages.HelloResponse)
	fmt.Printf("Response from remote %v", response.Message)

	console.ReadLine()
}