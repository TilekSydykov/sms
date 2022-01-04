package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solar-faza/entity"
	"solar-faza/i18n"
	"solar-faza/repository"
	"strconv"
)

var t *i18n.I18n

func init() {
	t = i18n.NewI18n()
}

type BasicController struct {
	data map[string]interface{}
	user *entity.User
}

func (b *BasicController) Init() {
	b.InitData()
}

func (b *BasicController) InitData() {
	b.data = make(map[string]interface{})
	b.user = nil
	b.data["T"] = t.Langs["ru"].Values
}

func (b *BasicController) ParseIntParams(c *gin.Context, paramNames ...string) ([]uint64, error) {
	var params = make([]uint64, len(paramNames))
	for i, paramName := range paramNames {
		id, err := strconv.ParseUint(c.Param(paramName), 10, 64)
		if err != nil {
			return nil, err
		}
		params[i] = id
	}
	return params, nil
}

func (b *BasicController) FetchUser(c *gin.Context) {
	b.InitData()
	id, ok := c.Get("userId")
	if ok && id.(uint64) > 0 {
		rep := repository.NewUserRepository()
		b.user = rep.GetUserById(id.(uint64))
		if b.user != nil {
			b.data["User"] = b.user
		}
	} else {
		delete(b.data, "User")
	}
}

func (b *BasicController) IsLoggedIn() bool {
	return b.user != nil && b.user.Id > 0
}

func (b *BasicController) HasLoggedInData(c *gin.Context) bool {
	id, ok := c.Get("userId")
	return ok && id.(uint64) > 0
}

func (b *BasicController) ShouldBeLoggedIn(c *gin.Context) {
	if !b.HasLoggedInData(c) {
		c.Redirect(302, "/login")
	}
}

func (b *BasicController) Render(c *gin.Context, templateName string) {
	c.HTML(http.StatusOK, templateName, b.data)
}

func (b *BasicController) AddData(key string, data interface{}) {
	b.data[key] = data
}

func (b *BasicController) SetTitle(data string) {
	b.data["Title"] = data
}

func (b *BasicController) HandleError(err error, c *gin.Context) {
	if err != nil {
		b.Render(c, "500.gohtml")
		return
	}
}

func (b *BasicController) Error(c *gin.Context) {
	b.Render(c, "500.gohtml")
	return
}

func (b *BasicController) Error404(c *gin.Context) {
	b.Render(c, "404.gohtml")
	return
}
