package front

import (
	"github.com/gin-gonic/gin"
	"solar-faza/controllers"
)

func ParticipantRouter(participant *gin.RouterGroup) {
	main := new(controllers.ParticipantController)
	main.Init()
	participant.GET("", main.MainPage)
	participant.GET(":id", main.GetParticipant)
	participant.GET(":id/edit", main.EditParticipant)
	participant.POST(":id/edit", main.EditParticipant)
	participant.GET("create", main.CreateBook)
	participant.POST("create", main.CreateParticipantPost)

	//participant.GET(":id/add", main.AddBookVolume)
	//participant.GET(":id/:vid", main.GetVolume)
	//participant.GET(":id/:vid/add", main.CreateChapter)
	//participant.GET(":id/:vid/:cid/edit", main.EditChapter)
	//participant.POST(":id/:vid/:cid/edit", main.EditChapter)
	//participant.GET(":id/:vid/:cid", main.ReadChapter)
}
