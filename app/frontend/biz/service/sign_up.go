package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/zemochen/go-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/zemochen/go-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type SignUpService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignUpService(Context context.Context, RequestContext *app.RequestContext) *SignUpService {
	return &SignUpService{RequestContext: RequestContext, Context: Context}
}

func (h *SignUpService) Run(req *auth.SignUpReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
