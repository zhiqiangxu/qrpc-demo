package service

import (
	"context"

	"github.com/zhiqiangxu/qchat/pkg/core"
)

type Service struct {
}

// Result for response
type Result struct {
	core.BaseResp
	N int
}

// SetError for output
func (r *Result) SetError(err error) {
	if err == nil {
		return
	}

	r.SetBase(core.ErrAPI, err.Error())
}

// Hello method
func (s *Service) Hello(ctx context.Context, v int) Result {
	return Result{N: v}
}
