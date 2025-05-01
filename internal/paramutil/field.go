package paramutil

import (
	"github.com/stainless-sdks/perigon-sdk-go/packages/param"
	"github.com/stainless-sdks/perigon-sdk-go/packages/resp"
)

func AddrIfPresent[T comparable](v param.Opt[T]) *T {
	if v.IsPresent() {
		return &v.Value
	}
	return nil
}

func ToOpt[T comparable](v T, meta resp.Field) param.Opt[T] {
	if meta.IsPresent() {
		return param.NewOpt(v)
	} else if meta.IsExplicitNull() {
		return param.NullOpt[T]()
	}
	return param.Opt[T]{}
}
