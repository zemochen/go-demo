package service

import (
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
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
	session := sessions.Default(h.RequestContext)

	session.Set("user_id", 1)
	session.Set("name", "test")
	session.Set("user_email", "test@gomall.com")
	err = session.Save()
	if err != nil {
		log.Println("SignUpService:save session error")
		return nil, err
	} else {
		log.Println("SignUpService:save session")
	}
	return
}
