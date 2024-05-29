package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/conf"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echo"
	"github.com/zemochen/go-demo/gomall/demo/demo_proto/middleware"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	c, err := echo.NewClient("demo_proto",
		client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		panic(err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")

	res, error := c.Echo(ctx, &pbapi.Request{Message: "hello world"})
	var bizErr *kerrors.GRPCBizStatusError
	if error != nil {
		ok := errors.As(error, &bizErr)
		if ok {
			fmt.Printf("biz error: %#v\n", bizErr)
		}
		klog.Fatal(error)
	}
	fmt.Printf("%v", res)
}
