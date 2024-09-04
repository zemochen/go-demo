package utils

import (
	"context"
	"log"

	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
)

func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	session := sessions.Default(c)

	// content["user_id"] = ctx.Value(frontendutils.UserIdKey)
	content["user_id"] = session.Get("user_id")
	log.Println("WarpResponse: user_id = ", content["user_id"])
	return content
}
