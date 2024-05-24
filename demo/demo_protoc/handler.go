package main

import (
	"context"
	pdapi "github.com/zemochen/go-demo/gomall/demo/demo_proto/kitex_gen/pdapi"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/biz/service"
)

// EchoServiceImpl implements the last service interface defined in the IDL.
type EchoServiceImpl struct{}

// Echo implements the EchoServiceImpl interface.
func (s *EchoServiceImpl) Echo(ctx context.Context, req *pdapi.Request) (resp *pdapi.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
