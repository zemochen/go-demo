package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	home "github.com/zemochen/go-demo/gomall/app/frontend/hertz_gen/frontend/home"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp := make(map[string]any)

	items := []map[string]any{
		{"Name": "T-shirt", "Price": 18, "img": "/static/images/item_1.jpg"},
		{"Name": "T-shirt2", "Price": 181, "img": "/static/images/item_2.jpg"},
		{"Name": "T-shirt3", "Price": 182, "img": "/static/images/item_3.jpg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
