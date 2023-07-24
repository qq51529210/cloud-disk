package authorize

import "github.com/gin-gonic/gin"

// getToken response_type=token
func getToken(ctx *gin.Context, req *getReq) {
	getHTML(ctx, req)
}
