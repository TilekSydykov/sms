package routes

import (
	"github.com/gin-gonic/gin"
	"solar-faza/controllers"
	"solar-faza/routes/front"
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
	front.ParticipantRouter(r.Group("participants"))
}
