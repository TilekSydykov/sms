package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"solar-faza/entity"
	"solar-faza/repository"
	"strconv"
	"strings"
)

type ParticipantController struct {
	BasicController
}

func (m *ParticipantController) MainPage(c *gin.Context) {
	m.FetchUser(c)
	m.AddData("Participants", repository.NewParticipantRepository().All())
	m.SetTitle("Участники")
	m.Render(c, "participants.gohtml")
}

func (m *ParticipantController) CreateBook(c *gin.Context) {
	m.FetchUser(c)
	m.SetTitle("Создание книги")
	m.Render(c, "createParticipant.gohtml")
}

//func (m *ParticipantController) AddBookVolume(c *gin.Context) {
//	m.FetchUser(c)
//	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		m.Error(c)
//		return
//	}
//	participant := repository.NewParticipantRepository().GetById(id)
//	if participants.Id < 1 || participants.UserId != m.user.Id {
//		m.Error(c)
//		return
//	}
//	m.SetTitle(participants.Title)
//	volume := repository.NewVolumeRepository().AddByBook(participants)
//	if volume.Num > 0 {
//		c.Redirect(302, fmt.Sprintf("/participants/%d/%d", participants.Id, volume.Num))
//	} else {
//		m.Error(c)
//	}
//}
//
//func (m *ParticipantController) GetVolume(c *gin.Context) {
//	m.FetchUser(c)
//	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		m.Error(c)
//		return
//	}
//	vid, err := strconv.ParseUint(c.Param("vid"), 10, 64)
//	if err != nil {
//		m.Error(c)
//		return
//	}
//	participants := repository.NewBookRepository().GetById(id)
//	if participants.Id < 1 {
//		m.Error(c)
//		return
//	}
//	volume := repository.NewVolumeRepository().GetByBookAndNum(id, vid)
//	volume.Chapters = repository.NewChapterRepository().AllFromVolume(volume.Id)
//	if volume.Id < 1 || volume.BookId != participants.Id {
//		m.Error404(c)
//		return
//	}
//	m.AddData("Book", participants)
//	m.AddData("Volume", volume)
//	m.SetTitle(participants.Title + " Том " + fmt.Sprintf("%d", volume.Num))
//	m.Render(c, "volume.gohtml")
//}

func (m *ParticipantController) EditParticipant(c *gin.Context) {
	m.FetchUser(c)
	m.ShouldBeLoggedIn(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		m.Error(c)
		return
	}
	participant := repository.NewParticipantRepository().GetById(id)
	submit, title, desc, active :=
		c.PostForm("submit"),
		c.PostForm("title"),
		c.PostForm("desc"),
		c.PostForm("active")

	if submit != "" || title != "" {
		var newText = strings.Replace(desc, "\n", "<br>", -1)
		participant.Title = title
		participant.Desc = newText
		participant.IsActive = active == "on"
		repository.NewParticipantRepository().Update(participant)
		c.Redirect(302, fmt.Sprintf("/participants/%d/edit", participant.Id))
		return
	}
	var newText = strings.Replace(participant.Desc, "<br>", "\n", -1)
	participant.Desc = newText
	m.AddData("Participant", participant)
	m.SetTitle(participant.Title + ": Редактирование")
	m.Render(c, "editParticipant.gohtml")
}

func (m *ParticipantController) GetParticipant(c *gin.Context) {
	m.FetchUser(c)
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		m.Error(c)
		return
	}
	participants := repository.NewParticipantRepository().GetById(id)
	if participants.Id < 1 {
		m.Error(c)
		return
	}
	m.SetTitle(participants.Title)
	participants.Views++
	repository.NewParticipantRepository().Update(participants)

	m.AddData("Participant", participants)
	m.Render(c, "participant.gohtml")
}

func (m *ParticipantController) CreateParticipantPost(c *gin.Context) {
	m.FetchUser(c)
	m.SetTitle("Создание")
	m.ShouldBeLoggedIn(c)
	participants := new(entity.Participant)
	participants.Title = c.PostForm("title")
	participants.Desc = c.PostForm("desc")
	if participants.Title == "" || participants.Desc == "" {
		m.AddData("Message", "Пустые поля")
		m.Render(c, "createBook.gohtml")
		return
	}
	participants.IsActive = c.PostForm("active") == "on"
	repository.NewParticipantRepository().Create(participants)
	c.Redirect(302, "/participants")
}
