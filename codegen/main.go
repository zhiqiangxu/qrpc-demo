package main

import (
	"github.com/zhiqiangxu/qrpc-demo/codegen/service"
	"github.com/zhiqiangxu/qrpc/codegen"
)

//请在generated目录运行该main函数，
//该main函数会将Client和Server代码生成到generated目录
//在哪个目录下运行，生成的package便是什么名字
//然后便可运行run_test中的TestCGPerformance
func main() {

	// 声明目标服务
	var s service.Service
	g := codegen.New("demo")

	// 自动注册s的方法
	g.Register(&s)

	// 生成代码，缓存到内存
	g.Generate()

	// 将生成的代码输出到Cwd
	g.Output()
}
