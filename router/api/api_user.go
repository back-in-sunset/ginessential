package api

import (
	"gin-essential/bll"
	"gin-essential/ginx"
	"gin-essential/schema"

	"gin-essential/pkg/errors"
	jwtauth "gin-essential/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户
type User struct {
	UserBll bll.User
}

// Register 注册
func (a *User) Register(c *gin.Context) {
	ctx := c.Request.Context()
	// 获取参数
	var params schema.User
	ginx.ParseJSON(c, &params)
	err := params.Validate()
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	// 判断手机号
	if a.UserBll.IsTelePhoneExist(ctx, params.Telephone) {
		ginx.ResError(c, errors.ErrPhoneRegistered)
		return
	}

	dkpassword, err := jwtauth.Scrypt(params.Password, jwtauth.Salt)
	if err != nil {
		ginx.ResError(c, errors.New500Response("注册失败"))
		return
	}
	// 创建用户
	params.Password = dkpassword
	err = a.UserBll.Register(ctx, params)
	if err != nil {
		panic(err)
	}
	ginx.ResOK(c)

}

// NatsMessage ..
func (a *User) NatsMessage(c *gin.Context) {
	ginx.ResError(c, errors.New400Response("注册失败"))
}

// // QueryStatistics ..
// func (a *User) QueryStatistics(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	err := a.UserBll.QueryStatistics(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	c.JSON(200, gin.H{
// 		"msg": "query clickhouse",
// 	})
// }

// QueryPage 查询分页
func (a *User) QueryPage(c *gin.Context) {
	var params schema.UserQueryParams
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	userResult, err := a.UserBll.QueryPage(c.Request.Context(), params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, userResult.Data, userResult.PageResult)
}
