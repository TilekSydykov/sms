package controllers

import (
	"github.com/gin-gonic/gin"
	"solar-faza/entity"
	"solar-faza/repository"
	"solar-faza/utils"
	"time"
)

type MainController struct {
	BasicController
}

func (m *MainController) MainPage(c *gin.Context) {
	m.FetchUser(c)
	m.SetTitle("Гонкиии!!!!")
	m.Render(c, "index.gohtml")
}

func (m *MainController) LoginPage(c *gin.Context) {
	m.InitData()
	m.SetTitle("Логин")
	m.Render(c, "login.gohtml")
}

func (m *MainController) LoginPost(c *gin.Context) {
	m.InitData()
	m.SetTitle("Логин")
	user := entity.User{}
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	if user.Email == "" || user.Password == "" {
		m.AddData("Message", "Пустые поля")
		m.Render(c, "login.gohtml")
		return
	}
	in, err := utils.SingIn(user.Email, user.Password)
	if err != nil {
		m.AddData("Message", "Неправильные данные для входа")
		m.Render(c, "login.gohtml")
		return
	}
	user.LastLogin = time.Now()
	repository.NewUserRepository().Update(&user)
	utils.GenerateCookies(c, in.AccessToken)
	c.Redirect(302, "/profile")
}

func (m *MainController) RegisterPage(c *gin.Context) {
	m.InitData()
	m.SetTitle("Регистрация")
	m.Render(c, "register.gohtml")
}

func (m *MainController) ProfilePage(c *gin.Context) {
	m.FetchUser(c)
	m.ShouldBeLoggedIn(c)
	m.AddData("Books", repository.NewParticipantRepository().GetByUserId(m.user.Id))
	m.SetTitle("Профиль")
	m.Render(c, "profile.gohtml")
}

func (m *MainController) RegisterPost(c *gin.Context) {
	m.InitData()
	user := entity.User{}
	user.Email = c.PostForm("email_reg")
	user.Password = c.PostForm("password_reg")
	if user.Email == "" || user.Password == "" {
		m.AddData("Message", "Пустые поля")
		m.Render(c, "register.gohtml")
		return
	}
	err := utils.RegisterUser(&user)
	if err != nil {
		m.AddData("Message", "Такой пользователь уже существует")
		m.Render(c, "register.gohtml")
		return
	}
	m.AddData("Message", "Зарегистрирован")
	m.Render(c, "login.gohtml")
}

func (m *MainController) LogOut(c *gin.Context) {
	m.InitData()
	utils.GenerateCookies(c, "")
	c.Redirect(302, "/")
}
