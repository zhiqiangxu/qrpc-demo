package main

import (
	"github.com/zhiqiangxu/qrpc-demo/codegen/service"
	"github.com/zhiqiangxu/qrpc/codegen"
)

func main() {
	var s service.Service
	g := codegen.New("demo")

	g.Register(&s)

	g.Generate()

	g.Output()
}
