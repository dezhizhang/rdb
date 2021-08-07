package model

import (
	"context"
	"rdb/driver"
)

type StringOption struct {
	ctx context.Context
}

type StringResult struct {
	Result string
	Err error
}

func NewStringResult(result string,err error) *StringResult  {
	return &StringResult{Result: result,Err: err}
}

func NewStringOption() *StringOption  {
	return &StringOption{ctx:context.Background()}
}

func (s *StringOption) Get(key string) *StringResult {
	return NewStringResult(driver.Redis().Get(s.ctx,key).Result())
}

func (s *StringResult) Unwrap() string {
	if s.Err != nil {
		panic(s.Err)

	}
	return s.Result
}

func (s *StringResult) UnwrapOr(str string) string {
	if s.Err != nil {
		return str
	}else {
		return s.Result
	}
}