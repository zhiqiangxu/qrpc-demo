package generated

import (
	context "context"
	json "encoding/json"
	fmt1 "fmt"
	qrpc "github.com/zhiqiangxu/qrpc"
	service "github.com/zhiqiangxu/qrpc-demo/codegen/service"
	codegen "github.com/zhiqiangxu/qrpc/codegen"
)

type DemoClient interface {
	Hello(context.Context, int) service.Result
}

type demoClient struct {
	*codegen.Client
}

func NewDemoClient(addrs []string, conf qrpc.ConnectionConfig) DemoClient {
	return demoClient{codegen.NewClient(qrpc.Cmd(10000), qrpc.Cmd(10001), addrs, conf)}
}

func (c demoClient) Hello(ctx context.Context, input int) (output service.Result) {
	cc := c.Client
	inBytes, err := json.Marshal(input)
	if err != nil {
		output.SetError(err)
		return
	}
	outBytes, err := cc.Request(ctx, "", "Hello", inBytes)
	if err != nil {
		output.SetError(err)
		return
	}
	err = json.Unmarshal(outBytes, &output)
	if err != nil {
		output.SetError(err)
	}
	return
}

type DemoServiceMux interface {
	Register(DemoService)
	RegisterSub(string, interface{})
	Mux() *qrpc.ServeMux
}

type DemoService interface {
	Hello(context.Context, int) service.Result
}

type demoServiceMux struct {
	callback map[string]codegen.MethodCall
	mux      *qrpc.ServeMux
}

func NewDemoServiceMux() DemoServiceMux {
	return &demoServiceMux{callback: make(map[string]codegen.MethodCall)}
}

func (m *demoServiceMux) Register(s DemoService) {
	m.callback[codegen.FQMethod("", "Hello")] = func(ctx context.Context, inBytes []byte) (outBytes []byte, err error) {
		var input int
		err = json.Unmarshal(inBytes, &input)
		if err != nil {
			return
		}
		output := s.Hello(ctx, input)
		outBytes, err = json.Marshal(output)
		return
	}
}

func (m *demoServiceMux) RegisterSub(ns string, ss interface{}) {
	switch ns {
	default:
		panic(fmt1.Sprintf("unknown ns:%v", ns))
	}
}

func (m *demoServiceMux) Mux() *qrpc.ServeMux {
	if m.mux != nil {
		return m.mux
	}
	mux := qrpc.NewServeMux()
	mux.Handle(qrpc.Cmd(10000), codegen.NewServiceHandler(qrpc.Cmd(10001), m.callback))
	m.mux = mux
	return mux
}
