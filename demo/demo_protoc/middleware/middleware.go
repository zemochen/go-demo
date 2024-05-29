package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func Middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// do something before invoking the endpoint
		begin := time.Now()
		err = next(ctx, req, resp)
		// do something after invoking the endpoint
		fmt.Println("time cost:", time.Since(begin))

		return err
	}
}
