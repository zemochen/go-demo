package service

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/zemochen/go-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/zemochen/go-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type SignInService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSignInService(Context context.Context, RequestContext *app.RequestContext) *SignInService {
	return &SignInService{RequestContext: RequestContext, Context: Context}
}

func (h *SignInService) Run(req *auth.SignInReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	session := sessions.Default(h.RequestContext)

	session.Set("user_id", 1)
	session.Set("name", "test")
	session.Set("user_email", "test@gomall.com")
	session.Save()
	log.Println("SignIn:save session")
	// resp = &common.Empty{}
	return
}
