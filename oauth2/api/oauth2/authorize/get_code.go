package authorize

import "github.com/gin-gonic/gin"

// getCode response_type=code
func getCode(ctx *gin.Context, req *getReq) {
	getHTML(ctx, req)
}
