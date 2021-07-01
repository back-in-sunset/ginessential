package api

import (
	"gin-essential/bll"
	"gin-essential/ginx"
	"gin-essential/model/entity"
	"gin-essential/pkg/util"
	"gin-essential/schema"
	"net/http"

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
	// 获取参数
	ctx := c.Request.Context()
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "must be 11 numbers"})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "must longer 6 numbers"})
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	// 判断手机号
	if a.UserBll.IsTelePhoneExist(ctx, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "user has existed"})
		return
	}

	dkpassword, err := jwtauth.Scrypt(password, jwtauth.Salt)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "register failed place try it again"})
		return
	}
	// 创建用户
	newUser := schema.User{
		UserEntity: entity.UserEntity{
			Name:      name,
			Password:  dkpassword,
			Telephone: telephone,
		},
	}

	err = a.UserBll.Register(ctx, newUser)
	if err != nil {
		panic(err)
	}
	ginx.ResOK(c)

}

// NatsMessage ..
func (a *User) NatsMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}

// QueryStatistics ..
func (a *User) QueryStatistics(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.UserBll.QueryStatistics(ctx)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"msg": "query clickhouse",
	})
}

func (a *User) QueryPage(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.UserQueryParams
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}

	params.Pagination = true
	a.UserBll.QueryPage(ctx, params)
}
