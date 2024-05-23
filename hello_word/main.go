package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
		// ctx.Data(consts.StatusOK,consts.MIMETextPlain, []byte("Hello World!"))
		ctx.JSON(consts.StatusOK, utils.H{"message": "Hello World!"})
	})

	h.Spin()
}
