package errorx

import (
	"errors"
	"fmt"
)

type Errorx struct {
	Code int64
	Msg  string
	Meta map[string]any
}

type Option func(*Errorx)

func WithMeta(meta map[string]interface{}) Option {
	return func(o *Errorx) {
		o.Meta = meta
	}
}

func New(code int64, msg string, opts ...Option) *Errorx {
	options := &Errorx{}
	for _, opt := range opts {
		opt(options)
	}
	return &Errorx{
		Code: code,
		Msg:  msg,
		Meta: options.Meta,
	}
}

func (e *Errorx) Error() string {
	if e.Meta != nil {
		return fmt.Sprintf("%s,%v", e.Msg, e.Meta)
	} else {
		return e.Msg
	}
}

func Is(err error, code int64) bool {
	if err == nil {
		return false
	}
	var e *Errorx
	if errors.As(err, &e) {
		return e.Code == code
	}
	return false
}
