package authorize

import (
	"net/http"
	"net/url"
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

// postToken 处理用户确认授权后的 token 流程
func postToken(ctx *gin.Context, req *postReq) {
	// 会话
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session[*db.User])
	// 令牌
	token := new(db.AccessToken)
	token.Type = *req.form.Client.TokenType
	token.Scope = parsePostScope(ctx, req.form.Client)
	token.ClientID = req.ClientID
	token.UserID = sess.Data.ID
	err := db.PutAccessToken(token)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 重定向
	if req.RedirectURI != "" {
		// 重定向地址
		_u, err := url.Parse(req.RedirectURI)
		if err != nil {
			internal.Submit400(ctx, err.Error())
			return
		}
		_u.RawQuery = util.HTTPQuery(token, _u.Query()).Encode()
		// 跳转
		ctx.Redirect(http.StatusSeeOther, _u.String())
		return
	}
	// 没有重定向，返回 JSON
	ctx.JSON(http.StatusOK, token)
}
