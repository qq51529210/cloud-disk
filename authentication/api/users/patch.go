package users

// type patchReq struct {
// 	// 密码，SHA1 格式
// 	Password *string `json:"password" binding:"required"`
// }

// // @Summary  用户管理
// // @Tags     修改
// // @Description 修改数据
// // @Param    id path string true "User.ID"
// // @Param    data body patchReq true "修改的字段"
// // @Security ApiKeyAuth
// // @Success  201
// // @Failure  400 {object} internal.Error
// // @Failure  401
// // @Failure  403
// // @Failure  500 {object} internal.Error
// // @Router   /users/{id} [patch]
// func patch(ctx *gin.Context) {
// 	// 参数
// 	var req patchReq
// 	err := ctx.ShouldBindJSON(&req)
// 	if err != nil {
// 		internal.Submit400(ctx, err.Error())
// 		return
// 	}
// 	if util.IsNilOrEmpty(&req) {
// 		internal.SubmitEmpty400(ctx)
// 		return
// 	}
// 	// 数据库
// 	var model db.User
// 	util.CopyStructAll(&model, &req)
// 	model.ID = ctx.Params[0].Value
// 	_, err = db.UpdateUser(&model)
// 	if err != nil {
// 		internal.DB500(ctx, err)
// 		return
// 	}
// 	//
// 	ctx.Status(http.StatusCreated)
// }
