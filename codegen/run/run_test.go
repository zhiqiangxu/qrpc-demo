package run

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/zhiqiangxu/qrpc"
	"github.com/zhiqiangxu/qrpc-demo/codegen/generated"
	"github.com/zhiqiangxu/qrpc-demo/codegen/service"
)

const (
	addr = "0.0.0.0:8001"
	n    = 100000
)

func startServer() {
	sm := generated.NewDemoServiceMux()
	var svc service.Service
	sm.Register(&svc)
	bindings := []qrpc.ServerBinding{
		qrpc.ServerBinding{Addr: addr, Handler: sm.Mux(), ReadFrameChSize: 10000}}
	server := qrpc.NewServer(bindings)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TestCGPerformance for test cg perf
func TestCGPerformance(t *testing.T) {
	// 启动服务端
	go startServer()
	time.Sleep(time.Second)

	// 生成客户端
	client := generated.NewDemoClient([]string{addr}, qrpc.ConnectionConfig{})

	i := 0
	var wg sync.WaitGroup
	startTime := time.Now()
	for {

		qrpc.GoFunc(&wg, func() {
			// 通过被替换的函数指针调用服务
			result := client.Hello(context.Background(), 123)
			if !(result.Err == "" && result.N == 123) {
				panic(fmt.Sprintf("fail:%v", result))
			}

		})
		i++
		if i > n {
			break
		}
	}

	wg.Wait()
	endTime := time.Now()
	t.Log(n, "request took", endTime.Sub(startTime))
}
