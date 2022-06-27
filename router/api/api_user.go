package api

import (
	"fmt"
	"gin-essential/ginx"
	"gin-essential/schema"

	"gin-essential/pkg/errors"
	"gin-essential/pkg/utils"

	usersrv "gin-essential/srv/user"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"))

// User 用户
type User struct {
	UserSrv usersrv.User
}

// Register 注册
// @Tags Users 用户
// @Summary 注册
// @Description 注册
// @Accept json
// @Produce json
// @Param body body schema.User true "用户"
// @Success 200 {object} schema.StatusResult "{staus:"OK", data:响应数据}"
// @Failure 400 {object} schema.ErrorItem "{code:400, status:"OK", message:"请求参数错误"}"
// @Failure 404 {object} schema.ErrorItem "{code:404, status:"OK", message:"资源不存在"}"
// @Router /api/users [post]
func (a *User) Register(c *gin.Context) {
	ctx := c.Request.Context()
	// 获取参数
	var user schema.User
	err := ginx.ParseJSON(c, &user)
	if err != nil {
		ginx.ResError(c, err)
		return
	}

	// 判断手机号
	if a.UserSrv.IsTelePhoneExist(ctx, user.Telephone) {
		ginx.ResError(c, errors.ErrPhoneRegistered)
		return
	}

	dkpassword, err := utils.Scrypt(user.Password, utils.Salt)
	if err != nil {
		ginx.ResError(c, errors.New400Response("注册失败"))
		return
	}
	// 创建用户
	user.Password = dkpassword
	err = a.UserSrv.Register(ctx, user)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResOK(c)
}

// NatsMessage ..
func (a *User) NatsMessage(c *gin.Context) {
	ginx.ResError(c, errors.New400Response("注册失败"))
}

// DocInt .
type DocInt int

// Query 查询多条数据
// @Tags Users 用户
// @Summary 查询多条数据
// @Description 查询多条数据
// @Accept json
// @Produce json
// @Param current query int false "分页索引" default(1)
// @Param page_size query int false "分页大小" default(10)
// @Param pagination query bool false "是否分页" default(true)
// @Param user_name query string false "用户名称"
// @Success 200 {object} schema.SuccessResult{data=schema.UserQueryResult} "{status:"OK", data:响应数据}"
// @Failure 400 {object} schema.ErrorItem "{code:400, status:"OK", message:"请求参数错误"}"
// @Failure 404 {object} schema.ErrorItem "{code:404, status:"OK", message:"资源不存在"}"
// @Router /api/users [get]
func (a *User) Query(c *gin.Context) {
	var params schema.UserQueryParams
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	userResult, err := a.UserSrv.QueryPage(c.Request.Context(), params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResPage(c, userResult.List, userResult.Pagination)
}

// Get 查询单条数据
// @Tags Users 用户
// @Summary 查询数据
// @Description 查询数据
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Param authorization header string false "jwt"
// @Success 200 {object} schema.User "{staus:"OK", data:响应数据}"
// @Failure 400 {object} schema.ErrorItem "{code:400, status:"OK", message:"请求参数错误"}"
// @Failure 404 {object} schema.ErrorItem "{code:404, status:"OK", message:"资源不存在"}"
// @Router /api/users/{id} [get]
func (a *User) Get(c *gin.Context) {
	fmt.Println("-------------------")
	user, err := a.UserSrv.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	} else if user == nil {
		ginx.ResError(c, errors.ErrNotFound, 200)
		return
	}

	ginx.ResItem(c, user)
}

// MockGet 查询单条数据
// @Tags Users 用户
// @Summary 查询数据
// @Description 查询数据
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} schema.User "{staus:"OK", data:响应数据}"
// @Failure 400 {object} schema.ErrorItem "{code:400, status:"OK", message:"请求参数错误"}"
// @Failure 404 {object} schema.ErrorItem "{code:404, status:"OK", message:"资源不存在"}"
// @Router /api/users/:id/detail [get]
func MockGet() {

}

// Start 查询单条数据
func (a *User) Start(c *gin.Context) {
	user, err := a.UserSrv.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	} else if user == nil {
		ginx.ResError(c, errors.ErrNotFound)
		return
	}

	ginx.ResItem(c, user)
}
