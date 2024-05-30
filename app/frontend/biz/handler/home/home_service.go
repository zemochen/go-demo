package home

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/zemochen/go-demo/gomall/app/frontend/biz/service"
	"github.com/zemochen/go-demo/gomall/app/frontend/biz/utils"
	home "github.com/zemochen/go-demo/gomall/app/frontend/hertz_gen/frontend/home"
)

// Home .
// @router /hello [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "home.tmpl", resp)
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
