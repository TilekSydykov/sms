package routes

import (
	"github.com/gin-gonic/gin"
	"solar-faza/controllers"
	"solar-faza/routes/midlewares"
)

func MainRouter(r *gin.RouterGroup) {
	main := new(controllers.MainController)
	main.Init()
	r.Use(midlewares.UserAuthMidleware())
	r.GET("", main.MainPage)
	r.GET("login", main.LoginPage)
	r.POST("login", main.LoginPost)
	r.GET("register", main.RegisterPage)
	r.GET("logout", main.LogOut)
	r.POST("register", main.RegisterPost)
	r.GET("profile", main.ProfilePage)

	r.GET("websocket", main.WebSocketHandler)
	r.GET("dashboard", main.Dashboard)
	r.GET("tokens", main.Tokens)
	r.GET("history", main.History)

	r.GET("send", main.MainForm)
	r.POST("send", main.MainForm)
}
