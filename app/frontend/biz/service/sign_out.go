package service

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	common "github.com/zemochen/go-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type SignOutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignOutService(Context context.Context, RequestContext *app.RequestContext) *SignOutService {
	return &SignOutService{RequestContext: RequestContext, Context: Context}
}

func (h *SignOutService) Run(req *common.Empty) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	session := sessions.Default(h.RequestContext)

	session.Clear()
	err = session.Save()
	if err != nil {
		log.Println("SignOutService:save session error")
		return nil, err
	} else {
		log.Println("SignOutService:save session")
	}
	return
}
