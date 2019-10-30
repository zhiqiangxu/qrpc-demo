package service

import (
	"context"
)

type Service struct {
}

// Result for response
type Result struct {
	Err string
	N   int
}

// SetError for output
func (r *Result) SetError(err error) {
	if err == nil {
		return
	}

	r.Err = err.Error()
}

// Hello method
func (s *Service) Hello(ctx context.Context, v int) Result {
	return Result{N: v}
}
