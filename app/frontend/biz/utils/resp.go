package utils

import (
	"context"

	frontendutils "github.com/zemochen/go-demo/gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	content["user_id"] = ctx.Value(frontendutils.UserIdKey)
	return content
}
