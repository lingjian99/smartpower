package utils

import (
	"context"
	"encoding/json"
)

func CtxInt64(ctx context.Context, key string) (int64, bool) {
	val := ctx.Value(key)
	switch vv := val.(type) {
	case int64:
		return vv, true
	case json.Number:
		vi, err := vv.Int64()
		if err != nil {
			return 0, false
		}
		return vi, true
	}
	return 0, false
}
