package main

import (
	_ "protoactor_demo/demo/single"
	"protoactor_demo/demo/node1"
	"protoactor_demo/demo/node2"
)

func main() {
	//single.RunSingle()
	go node2.Run()
	node1.Run()
}