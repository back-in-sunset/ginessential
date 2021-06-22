package api

import (
	"gin-essential/bll"
	"gin-essential/common/util"
	"gin-essential/model/entity"
	"gin-essential/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 用户
type User struct {
	BllUser bll.User
}

// Register 注册
func (a *User) Register(c *gin.Context) {
	// 获取参数
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
	if a.BllUser.IsTelePhoneExist(telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "user has existed"})
		return
	}

	dkpassword, err := util.Scrypt(password)
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

	err = a.BllUser.Register(newUser)
	if err != nil {
		panic(err)
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

// NatsMessage ..
func (a *User) NatsMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "注册成功",
	})

}
