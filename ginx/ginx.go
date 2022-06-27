package ginx

import (
	"fmt"
	contextx "gin-essential/ctx"
	"gin-essential/logger"
	"gin-essential/pkg/errors"
	"gin-essential/pkg/jsonx"
	"gin-essential/schema"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/validator.v2"
)

const (
	// ReqBodyKey 请求body
	ReqBodyKey = "/req-body"
	// ResBodyKey 响应body
	ResBodyKey = "/res-body"

	// header
	authorization = "Authorization"
	jwtPrefix     = "Bearer "
	contentType   = "application/json; charset=utf-8"
)

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	if err := validator.Validate(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("参数校验不通过 - %s", err.Error()))
	}

	return nil
}

// ParseQuery 解析Query参数
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ParseForm 解析Form请求
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	// ResSuccess(c, schema.StatusResult{Status: schema.OKStatus})
	resSuccess(c, schema.SuccessResult{
		Status: schema.OKStatus,
		Data:   schema.StatusResult{Status: schema.OKStatus},
	})

}

// ResList 响应列表数据
func ResList(c *gin.Context, v interface{}) {
	resSuccess(c, schema.SuccessResult{
		Status: schema.OKStatus,
		Data: schema.ListResult{
			List: v,
		},
	})
}

// ResPage 响应分页数据
func ResPage(c *gin.Context, v interface{}, pr *schema.PaginationResult) {
	resSuccess(c, schema.SuccessResult{
		Status: schema.OKStatus,
		Data: schema.ListResult{
			List:       v,
			Pagination: pr,
		},
	})
}

// ResItem 响应单条数据
func ResItem(c *gin.Context, v interface{}) {
	resSuccess(c, schema.SuccessResult{
		Status: schema.OKStatus,
		Data:   v,
	})
}

// ResSuccess 响应成功
func resSuccess(c *gin.Context, v interface{}) {
	resJSON(c, http.StatusOK, v)
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	var res *errors.ResponseError

	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}
	}

	if code := res.Code; code >= 400 && code < 500 {
		logger.Warn(fmt.Sprintf("%+v", err))
	} else if code >= 500 {
		traceID, ok := contextx.FromTraceID(c.Request.Context())
		if ok {
			logger.Error(fmt.Sprintf("[%s] %+v", traceID, err))
		} else {
			logger.Error(fmt.Sprintf("%+v", err))
		}
	}

	eitem := schema.ErrorItem{
		Status:  schema.ErrorStatus,
		Code:    res.Code,
		Message: res.Message,
	}

	resJSON(c, res.StatusCode, schema.ErrorResult{ErrorItem: eitem})
}

// resJSON 响应JSON数据
func resJSON(c *gin.Context, status int, v interface{}) {
	buf, err := jsonx.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, contentType, buf)
	c.Abort()
}

// GetBody Get request body
func GetBody(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}

// GetToken 获取token
func GetToken(c *gin.Context) string {
	var token string
	token = c.GetHeader(authorization)
	if strings.HasPrefix(token, jwtPrefix) {
		token = token[len(jwtPrefix):]
	}
	return token
}
