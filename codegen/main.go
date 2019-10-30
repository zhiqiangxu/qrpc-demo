package main

import (
	"github.com/zhiqiangxu/qrpc-demo/codegen/service"
	"github.com/zhiqiangxu/qrpc/codegen"
)

//请在generated目录运行该main函数
func main() {

	var s service.Service
	g := codegen.New("demo")

	g.Register(&s)

	g.Generate()

	g.Output()
}
